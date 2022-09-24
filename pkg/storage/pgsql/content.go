package pgsql

import (
	"context"
	"retelling/pkg/models"
)

func (s *Storage) NewContent(data models.Content) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	INSERT INTO content (
		type_id,
		genre1_id,
		genre2_id,
		genre3_id,
		title,
		likes
	)
	VALUES ($1,$2,$3,$4,$5,$6) 
	RETURNING content_id
	`,
		data.TypeID,
		data.GenreID1,
		data.GenreID2,
		data.GenreID3,
		data.Title,
		data.Likes).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}