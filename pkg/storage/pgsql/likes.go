package pgsql

import (
	"context"
	"retelling/pkg/models"
)

func (s *Storage) GetLikes(req models.Request) ([]models.Content, error) {
	var data []models.Content
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id,
		type_id,
		genre1_id,
		genre2_id,
		genre3_id,
		title,
		description,
		users_liked
	FROM content
		WHERE $1 = ANY(users_liked)
	`, req.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.Content
		err = rows.Scan(
			&item.ID,
			&item.TypeID,
			&item.GenreID1,
			&item.GenreID2,
			&item.GenreID3,
			&item.Title,
			&item.Description,
			&item.UsersLiked,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (s *Storage) GetUsersLiked(req models.Request) ([]models.User, error) {
	var data []models.User
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id,
		name,
		login,
		password,
		age,
		review_count,
		rating,
		profession,
		likes
	FROM users
		WHERE $1 = ANY(likes)
	`, req.ObjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.User
		err = rows.Scan(
			&item.ID,
			&item.Data.Name,
			&item.Login,
			&item.Password,
			&item.Data.Age,
			&item.Data.ReviewCount,
			&item.Data.Rating,
			&item.Data.Profession,
			&item.Data.Likes,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (s *Storage) NewLike(req models.Request) error {
	_ = s.pool.QueryRow(context.Background(), `
		UPDATE users
		SET likes = array_append(likes, $2)
		WHERE id = $1
	`, req.UserID, req.ObjectID).Scan()
	_ = s.pool.QueryRow(context.Background(), `
		UPDATE content
		SET users_liked = array_append(users_liked, $1)
		WHERE id = $2
	`, req.UserID, req.ObjectID).Scan()
	return nil
}

func (s *Storage) Unlike(req models.Request) error {
	_ = s.pool.QueryRow(context.Background(), `
		UPDATE users
		SET likes = array_remove(likes, $2)
		WHERE id = $1
	`, req.UserID, req.ObjectID).Scan()
	_ = s.pool.QueryRow(context.Background(), `
		UPDATE content
		SET users_liked = array_remove(users_liked, $1)
		WHERE id = $2
	`, req.UserID, req.ObjectID).Scan()
	return nil
}
