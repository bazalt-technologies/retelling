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
	RETURNING id
	`,
		data.TypeID,
		data.GenreID1,
		data.GenreID2,
		data.GenreID3,
		data.Title,
		data.UsersLiked).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) DeleteContent(data models.Content) error {
	err := s.pool.QueryRow(context.Background(), `
		DELETE FROM content WHERE id = $1
		DELETE FROM review WHERE content_id = $1
	`,
		data.ID).Scan()
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
		WHERE id = $1
	`, data.ID,
		data.TypeID,
		data.GenreID1,
		data.GenreID2,
		data.GenreID3,
		data.Title,
		data.UsersLiked).Scan()
	return err
}

func (s *Storage) GetContent(req models.Request) (models.Content, error) {
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		type_id,
		genre1_id,
		genre2_id,
		genre3_id,
		title,
		likes
	FROM content
		WHERE id = $1
	`,
		req.ObjectID)
	if err != nil {
		return models.Content{}, err
	}
	defer rows.Close()
	var item models.Content
	err = rows.Scan(
		&item.ID,
		&item.TypeID,
		&item.GenreID1,
		&item.GenreID2,
		&item.GenreID3,
		&item.Title,
		&item.UsersLiked,
	)
	if err != nil {
		return models.Content{}, err
	}
	return item, nil
}
