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

func (s *Storage) DeleteContent(data models.Content) error {
	err := s.pool.QueryRow(context.Background(), `
		DELETE FROM content
		WHERE content_id = $1
	`,
		data.ContentID).Scan()
	return err
}

func (s *Storage) PatchContent(data models.Content) error {
	err := s.pool.QueryRow(context.Background(), `
		UPDATE content
		SET type_id = $2
			genre1_id = $3
			genre2_id = $4
			genre3_id = $5
			title = $6
			likes = $7
		WHERE content_id = $1
	`, data.ContentID,
		data.TypeID,
		data.GenreID1,
		data.GenreID2,
		data.GenreID3,
		data.Title,
		data.Likes).Scan()
	return err
}
