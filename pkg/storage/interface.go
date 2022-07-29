package storage

import "retelling/pkg/models"

type DB interface {
	GetUsers(req models.Request) ([]models.User, error)
	NewUser(data models.User) (int, error)
	AuthUser(req models.Request) (int, error)
	UpdateUser(item models.User) (int, error)
	DeleteUser(id int) (int, error)
	
	NewReview(data models.Review) (int, error)
	GetReviews(req models.Request) ([]models.Review, error)
	UpdateReview(data models.Review) (int, error)
	DeleteReview(int id) (int, error)
	
	NewGenre(data models.Genre) (int, error)
	GetGenres(req models.Request) ([]models.Genre, error)
	UpdateGenre(data models.Genre) (int, error)
	DeleteGenre(id int) (int, error)
	
	NewType(data models.Type) (int, error)
	GetTypes(req models.Request) ([]models.Type, error)
	UpdateType(data models.Type) (int, error)
	DeleteType(id int) (int, error)
}
