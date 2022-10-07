package pgsql

import (
	"context"
	"log"
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
		description,
		users_liked
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7) 
	RETURNING id
	`,
		data.TypeID,
		data.GenreID1,
		data.GenreID2,
		data.GenreID3,
		data.Title,
		data.Description,
		data.UsersLiked).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) DeleteContent(req models.Request) error {
	err := s.pool.QueryRow(context.Background(), `
		DELETE FROM reviews WHERE content_id = $1;
	`,
		req.ObjectID).Scan()
	_ = s.pool.QueryRow(context.Background(), `
		DELETE FROM content WHERE id = $1;
	`,
		req.ObjectID).Scan()
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
			description = $7
			users_liked = $8
		WHERE id = $1
	`, data.ID,
		data.TypeID,
		data.GenreID1,
		data.GenreID2,
		data.GenreID3,
		data.Title,
		data.Description,
		intToInt32Array(data.UsersLiked)).Scan()
	return err
}

func (s *Storage) GetContent(req models.Request) ([]models.Content, error) {
	if len(req.ObjectIDs) == 0 && req.ObjectID != 0 {
		req.ObjectIDs = append(req.ObjectIDs, req.ObjectID)
	}
	log.Println(req)
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id,
		type_id,
		genre1_id,
		genre2_id,
		genre3_id,
		title,
		description,
		users_liked
	FROM content
	WHERE (content.id = ANY($1) OR array_length($1,1) is NULL)
	`,
		intToInt32Array(req.ObjectIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []models.Content
	for rows.Next() {
		var item models.Content
		err = rows.Scan(
			&item.ID,
			&item.TypeID,
			&item.GenreID1,
			&item.GenreID2,
			&item.GenreID3,
			&item.Title,
			&item.Description,
			&item.UsersLiked,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	log.Println(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
