package shortner

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

type ShortnerService interface {
	Find(code string) (*Shortner, error)
	Store(shortner *Shortner) error
}

type shortnerService struct {
	repo ShortnerRepository
}

func NewShortnerService(repo ShortnerRepository) ShortnerService {

	NewShortnerThread(repo).StartCleanupJob()

	return &shortnerService{
		repo: repo,
	}
}

var (
	shortnerCache = cache.New(12*time.Hour, 2*time.Hour)
)

func (s *shortnerService) Find(code string) (*Shortner, error) {
	if shortner, found := shortnerCache.Get(code); found {
		log.Printf("Cache hit for code: %s", code)

		go s.repo.HitIncrement(shortner.(*Shortner))
		return shortner.(*Shortner), nil
	}

	short, err := s.repo.Find(code)
	if err != nil {
		return nil, err
	}

	if short != nil {
		shortnerCache.Set(code, short, cache.DefaultExpiration)
	}

	go s.repo.HitIncrement(short)

	return short, nil
}

func (s *shortnerService) Store(shortner *Shortner) error {

	exist, _ := s.repo.FindByUrl(shortner.Url)
	if exist != nil {
		shortnerCache.Set(exist.Code, exist, cache.DefaultExpiration)

		shortner.Code = exist.Code
		shortner.RedirectCount = exist.RedirectCount
		shortner.CreatedAt = exist.CreatedAt
		shortner.UpdatedAt = exist.UpdatedAt
		shortner.ID = exist.ID
		shortner.Redirect = exist.Redirect

		return nil
	}

	shortner.Code = generateCode(shortner.Url)
	shortner.RedirectCount = 0

	return s.repo.Store(shortner)
}

func generateCode(url string) string {
	currentDate := time.Now().Format("20060102")

	input := url + currentDate
	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)

	encoded := base64.URLEncoding.EncodeToString(hashBytes)

	if len(encoded) > 10 {
		encoded = encoded[:10]
	}

	return encoded
}
