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
		type_id,
		genre1_id,
		genre2_id,
		genre3_id,
		title,
		rating,
		date,
		review,
		image_link,
		likes
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) 
	RETURNING id
	`,
		data.UserID,
		data.TypeID,
		data.Genre1ID,
		data.Genre2ID,
		data.Genre3ID,
		data.Title,
		data.Rating,
		data.Date,
		data.Review,
		data.ImageLink,
		data.Likes).Scan(&id)
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
		type_id,
		genre1_id,
		genre2_id,
		genre3_id,
		title,
		rating,
		date,
		review,
		image_link,
		likes
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
			&item.ID,
			&item.TypeID,
			&item.Genre1ID,
			&item.Genre2ID,
			&item.Genre3ID,
			&item.Title,
			&item.Rating,
			&item.Date
			&item.Review,
			&item.ImageLink,
			&item.Likes
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (s *Storage) UpdateReview(data models.Review) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	UPDATE reviews
		SET
		user_id = $2,
		type_id = $3,
		genre1_id = $4,
		genre2_id = $5,
		genre3_id = $6,
		title = $7,
		rating = $8,
		date = $9,
		review = $10,
		image_link = $11,
		likes = $12
	WHERE id = $1
	RETURNING id
	`,
		data.ID,
		data.UserID,
		data.TypeID,
		data.Genre1ID,
		data.Genre2ID,
		data.Genre3ID,
		data.Title,
		data.Rating,
		date.Date,
		data.Review,
		data.ImageLink,
		data.Likes).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) DeleteReview(id int) (int, error) {
	err := s.pool.QueryRow(context.Background(), `
	DELETE FROM reviews WHERE id = $1
	`,
		id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
