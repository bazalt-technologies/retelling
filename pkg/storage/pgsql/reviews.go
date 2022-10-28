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
	RETURNING id
	`,
		data.ContentID,
		data.UserID,
		data.Review,
		data.Date).Scan(&id)
	_, err = s.pool.Exec(context.Background(), `
	UPDATE users SET review_count = review_count+1 WHERE id = $1`,
		data.UserID)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) GetReviews(req models.Request) ([]models.Review, error) {
	var data []models.Review
	if len(req.ObjectIDs) == 0 && req.ObjectID != 0 {
		req.ObjectIDs = append(req.ObjectIDs, req.ObjectID)
	}
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id,
		user_id,
		content_id,
		review,
		date
	FROM reviews
		WHERE (user_id=$1 OR $1 = 0) AND (content_id=ANY($2) OR array_length($2,1) is NULL)
	`,
		req.UserID,
		intToInt32Array(req.ObjectIDs))

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.Review
		err = rows.Scan(
			&item.ID,
			&item.UserID,
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
	_, err := s.pool.Exec(context.Background(), `
		UPDATE reviews
		SET content_id = $2
			user_id = $3
			review = $4
			date = $5
		WHERE id = $1
	`, data.ID,
		data.ContentID,
		data.UserID,
		data.Review,
		data.Date)
	return err
}

func (s *Storage) DeleteReview(data models.Review) error {
	_, err := s.pool.Exec(context.Background(), `
		DELETE FROM reviews
		WHERE id = $1;
	`,
		data.ID)
	if err != nil {
		return err
	}
	_, err = s.pool.Exec(context.Background(), `
	UPDATE users SET review_count = review_count-1 WHERE id = $1`,
		data.UserID)
	return err
}
