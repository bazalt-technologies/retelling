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

func (s *Storage) UpdateReview(data models.Review) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	UPDATE reviews
		SET
		type = $2,
		genre = $3,
		title = $4,
		rating = $5,
		review = $6
	WHERE user_id = $1
	RETURNING Id
	`,
		data.UserID,
		data.Type,
		data.Genre,
		data.Title,
		data.Rating,
		data.Review).Scan(&id)
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
