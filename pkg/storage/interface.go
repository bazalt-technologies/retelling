package storage

import "retelling/pkg/models"

type DB interface {
	GetUsers(req models.Request) ([]models.User, error)
	GetReviews(req models.Request) ([]models.Review, error)
	NewUser(data models.User) (int, error)
	NewReview(data models.Review) (int, error)
	AuthUser(req models.Request) (int, error)
}
