package pgsql

import (
	"context"
	"retelling/pkg/models"
)

func (s *Storage) GetTypes(req models.Request) ([]models.Type, error) {
	var data []models.Type
	if len(req.ObjectIDs) == 0 && req.ObjectID != 0 {
		req.ObjectIDs = append(req.ObjectIDs, req.ObjectID)
	}
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id, 
		type
	FROM types
		WHERE (id=ANY($1) OR array_length($1,1) is NULL)
	`, intToInt32Array(req.ObjectIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.Type
		err = rows.Scan(
			&item.ID,
			&item.Type,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}
