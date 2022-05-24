package pgsql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"retelling/pkg/models"
)

type Storage struct {
	pool *pgxpool.Pool
}

func (s *Storage) GetUsers(req models.Request) ([]models.User, error) {
	var data []models.User
	if len(req.UserIDs) == 0 {
		req.UserIDs = append(req.UserIDs, req.UserID)
	}
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id,
		name,
		login,
		password,
		age
	FROM users
		WHERE (id=ANY($1) OR array_length($1) is NULL)
	`, intToInt32Array(req.UserIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.User
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Login,
			&item.Password,
			&item.Age,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (s *Storage) GetReviews(req models.Request) ([]models.Review, error) {
	var data []models.Review
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id, 
		type,
		genre,
		title,
		rating,
		review,
		likes
	FROM reviews
		WHERE (user_id=$1 OR $1 = 0)
	`,
		req.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.Review
		err = rows.Scan(
			&item.ID,
			&item.Type,
			&item.Genre,
			&item.Rating,
			&item.Review,
			&item.Likes,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (s *Storage) NewUser(data models.User) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	INSERT INTO users (
		name, 
		login,
		password,
		age	
	)
	VALUES ($1,$2,$3,$4) RETURNING id
	`,
		data.Name,
		data.Login,
		data.Password,
		data.Age).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) NewReview(data models.Review) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	INSERT INTO reviews (
		user_id,
		type,
		genre,
		title,
		rating,
		review
	)
	VALUES ($1,$2,$3,$4,$5,$6) 
	RETURNING Id
	`,
		data.UserID,
		data.Type,
		data.Genre,
		data.Title,
		data.Rating,
		data.Review).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func intToInt32Array(in []int) []int32 {
	var out []int32
	for _, val := range in {
		out = append(out, int32(val))
	}
	return out
}
