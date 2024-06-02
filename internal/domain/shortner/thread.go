package shortner

import (
	"log"
	"time"
)

type ShortnerThread interface {
	StartCleanupJob()
}

type shortnerThread struct {
	repo ShortnerRepository
}

func NewShortnerThread(repo ShortnerRepository) ShortnerThread {
	return &shortnerThread{
		repo: repo,
	}
}

func (s *shortnerThread) StartCleanupJob() {
	ticker := time.NewTicker(1 * time.Hour)

	s.cleanupOldShortners()

	go func() {
		for range ticker.C {
			s.cleanupOldShortners()
		}
	}()
}

func (s *shortnerThread) cleanupOldShortners() {

	if s.repo == nil {
		log.Printf("Repository is nil")
		return
	}

	shortners, err := s.repo.FindToDelete()

	if err != nil {
		log.Printf("Error finding shortners to delete: %v", err)
		return
	}

	log.Printf("Found %d shortners to delete", len(shortners))

	for _, shortner := range shortners {
		if err := s.repo.Delete(&shortner); err != nil {
			log.Printf("Error deleting shortner with ID %d: %v", shortner.ID, err)
		} else {
			log.Printf("Deleted shortner with ID %d", shortner.ID)
		}
	}
}
