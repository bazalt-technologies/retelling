package storage

import "retelling/pkg/models"

type DB interface {
	GetUsers(req models.Request) ([]models.User, error)
	NewUser(data models.User) (int, error)
	AuthUser(req models.Request) (int, error)
	UpdateUser(item models.User) (int, error)
	DeleteUser(req models.Request) error

	NewReview(data models.Review) (int, error)
	GetReviews(req models.Request) ([]models.Review, error)
	PatchReview(data models.Review) error
	DeleteReview(data models.Review) error

	NewContent(data models.Content) (int, error)
	DeleteContent(req models.Request) error
	PatchContent(data models.Content) error
	GetContent(req models.Request) ([]models.Content, error)

	GetLikes(req models.Request) ([]models.Content, error)
	GetUsersLiked(req models.Request) ([]models.User, error)

	GetGenres(req models.Request) ([]models.Genre, error)
	//TODO: NewGenre(data models.Genre) (int, error)
	//TODO: UpdateGenre(data models.Genre) (int, error)
	//TODO: DeleteGenre(id int) (int, error)

	GetTypes(req models.Request) ([]models.Type, error)
	//TODO: NewType(data models.Type) (int, error)
	//TODO: UpdateType(data models.Type) (int, error)
	//TODO: DeleteType(id int) (int, error)
}
