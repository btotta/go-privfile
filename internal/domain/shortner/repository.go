package shortner

import (
	"time"

	"gorm.io/gorm"
)

type ShortnerRepository interface {
	Find(code string) (*Shortner, error)
	Store(shortner *Shortner) error
	HitIncrement(shortner *Shortner) error
	FindToDelete() ([]Shortner, error)
	Delete(shortner *Shortner) error
	FindByUrl(url string) (*Shortner, error)
}

var (
	daysToKeepLasUpdated = 7
)

type shortnerRepository struct {
	db *gorm.DB
}

func NewShortnerRepository(db *gorm.DB) ShortnerRepository {
	return &shortnerRepository{
		db: db,
	}
}

func (r *shortnerRepository) Find(code string) (*Shortner, error) {
	var shortner Shortner
	if err := r.db.Where("code = ?", code).First(&shortner).Error; err != nil {
		return nil, err
	}

	return &shortner, nil
}

func (r *shortnerRepository) Store(shortner *Shortner) error {
	if err := r.db.Create(shortner).Error; err != nil {
		return err
	}

	return nil
}

func (r *shortnerRepository) HitIncrement(shortner *Shortner) error {
	if err := r.db.Model(shortner).Update("redirect_count", shortner.RedirectCount+1).Error; err != nil {
		return err
	}

	return nil
}

func (r *shortnerRepository) FindToDelete() ([]Shortner, error) {
	var shortners []Shortner

	query := r.db.Model(&Shortner{})
	query.Where("updated_at < ?", time.Now().AddDate(0, 0, -daysToKeepLasUpdated))

	if err := query.Find(&shortners).Error; err != nil {
		return nil, err
	}

	return shortners, nil
}

func (r *shortnerRepository) Delete(shortner *Shortner) error {

	if err := r.db.Exec("DELETE FROM shortners WHERE id = ?", shortner.ID).Error; err != nil {
		return err
	}
	return nil
}


func (r *shortnerRepository) FindByUrl(url string) (*Shortner, error) {
	var shortner Shortner
	if err := r.db.Where("url = ?", url).First(&shortner).Error; err != nil {
		return nil, err
	}

	return &shortner, nil
}
