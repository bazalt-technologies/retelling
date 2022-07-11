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
	RETURNING Id
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
		req.UserID)
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
		type_id = $2,
		genre1_id = $3,
		genre2_id = $4,
		genre3_id = $5,
		title = $6,
		rating = $7,
		date = $8,
		review = $9,
		image_link = $10,
		likes = $11
	WHERE user_id = $1
	RETURNING Id
	`,
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
	// Возможно стоит учитывать то, что будут обновляться только некоторые поля
	// Для этого можно сделать по методу на каждое поле, либо как-то вводить data с пустыми полями,
	// Чтобы можно было определить, что их трогать не надо
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) DeleteReview(id int) (int, error) {
	err := s.pool.QueryRow(context.Background(), `
	DELETE FROM reviews WHERE user_id = $1
	`,
		id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
