package pgsql

import (
	"context"
	"retelling/pkg/models"
)

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
		age,
		review_count,
		rating,
		profession
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
			&item.Data.Name,
			&item.Login,
			&item.Password,
			&item.Data.Age,
			&item.Data.ReviewCount,
			&item.Data.Rating,
			&item.Data.Profession,
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
		age,
		profession
	)
	VALUES ($1,$2,$3,$4) RETURNING id
	`,
		data.Data.Name,
		data.Login,
		data.Password,
		data.Data.Age,
		data.Data.Profession,
	).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}
