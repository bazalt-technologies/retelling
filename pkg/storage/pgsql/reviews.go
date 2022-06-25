package pgsql

import (
	"context"
	"retelling/pkg/models"
)

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
