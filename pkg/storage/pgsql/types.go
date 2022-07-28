package pgsql

import (
	"context"
	"retelling/pkg/models"
)

func (s *Storage) NewType(data models.Type) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	INSERT INTO types (
		type
	)
	VALUES ($1) 
	RETURNING id
	`,
		data.Type).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) GetTypes(req models.Request) ([]models.Type, error) {
	var data []models.Review
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id, 
		type
	FROM types
		WHERE (id=ANY($1) OR array_length($1) is NULL)
	`, intToInt32Array(req.UserIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.Type
		err = rows.Scan(
			&item.ID,
			&item.Type
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (s *Storage) UpdateType(data models.Type) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	UPDATE types
		SET
		type = $2
	WHERE id = $1
	RETURNING id
	`,
		data.ID,
		data.Type).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) DeleteType(id int) (int, error) {
	err := s.pool.QueryRow(context.Background(), `
	DELETE FROM types WHERE id = $1
  DELETE FROM reviews WHERE type_id = $2
	`,
		id, id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
