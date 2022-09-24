package pgsql

import (
	"context"
	"retelling/pkg/models"
)

func (s *Storage) NewReview(data models.Review) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	INSERT INTO reviews (
		content_id,
		user_id,
		review,
		date
	)
	VALUES ($1,$2,$3,$4) 
	RETURNING review_id
	`,
		data.ContentID,
		data.UserID,
		data.Review,
		data.Date).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) GetReviews(req models.Request) ([]models.Review, error) {
	var data []models.Review
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		review_id,
		content_id,
		review,
		date
	FROM reviews
		WHERE (user_id=$1 OR $1 = 0)
	`,
		req.ObjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.Review
		err = rows.Scan(
			&item.ReviewID,
			&item.ContentID,
			&item.Review,
			&item.Date,
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
		SET content_id = $2
			user_id = $3
			review = $4
			date = $5
		WHERE review_id = $1
	`, data.ReviewID,
		data.ContentID,
		data.UserID,
		data.Review,
		data.Date).Scan()
	return err
}

func (s *Storage) DeleteReview(data models.Review) error {
	err := s.pool.QueryRow(context.Background(), `
		DELETE FROM reviews
		WHERE id = $1
	`,
		data.ReviewID).Scan()
	return err
}
