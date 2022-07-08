package pgsql

import (
	"context"
	"retelling/pkg/models"
)

// TODO - поменять методы БД  структуры и методы апи, согласно вынесению жанров и типов в отдельные таблицы.

func (s *Storage) NewReview(data models.Review) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	INSERT INTO reviews (
		user_id,
		type,
		genre1,
		genre2,
		genre3,
		title,
		rating,
		review
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8) 
	RETURNING Id
	`,
		data.UserID,
		data.Type,
		data.Genre1,
		data.Genre2,
		data.Genre3,
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
		genre1,
		genre2,
		genre3,
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
			&item.Genre1,
			&item.Genre2,
			&item.Genre3,
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

func (s *Storage) PatchReview(data models.Review) error {
	err := s.pool.QueryRow(context.Background(), `
		UPDATE reviews
		SET type = $2
			genre1 = $3
			genre2 = $4
			genre3 = $5
			title = $6
			rating = $7
			review = $8
			likes = $9
		WHERE id = $1
	`, data.UserID,
		data.Type,
		data.Genre1,
		data.Genre2,
		data.Genre3,
		data.Title,
		data.Rating,
		data.Review).Scan()
	return err
}

func (s *Storage) DeleteReview(data models.Review) error {
	err := s.pool.QueryRow(context.Background(), `
		DELETE FROM reviews
		WHERE id = $1
	`,
		data.UserID).Scan()
	return err
}
