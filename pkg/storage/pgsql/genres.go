package pgsql

import (
	"context"
	"retelling/pkg/models"
)


func (s *Storage) GetGenres(req models.Request) ([]models.Genre, error) {
	var data []models.Review
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id, 
		genre
	FROM genres
		WHERE (id=ANY($1) OR array_length($1) is NULL)
	`, intToInt32Array(req.ObjectIDs)) 
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.Genre
		err = rows.Scan(
			&item.ID,
			&item.Genre
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

