package pgsql

import (
	"context"
	"retelling/pkg/models"
)

func (s *Storage) NewGenre(data models.Genre) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	INSERT INTO genres (
		genre
	)
	VALUES ($1) 
	RETURNING id
	`,
		data.Genre).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) GetGenres(req models.Request) ([]models.Genre, error) {
	var data []models.Review
	rows, err := s.pool.Query(context.Background(), `
	SELECT 
		id, 
		genre
	FROM genres
		WHERE (id=ANY($1) OR array_length($1) is NULL)
	`, intToInt32Array(req.ReviewIDs)) // ReviewIDs - так как нет поля GenreIDs (возможно стоит переименовать ReviewIDs в RequestIDs/ItemIDs
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

func (s *Storage) UpdateGenre(data models.Genre) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	UPDATE genres
		SET
		genre = $2
	WHERE id = $1
	RETURNING id
	`,
		data.ID,
		data.Genre).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) DeleteGenre(id int) (int, error) {
	err := s.pool.QueryRow(context.Background(), `
	DELETE FROM genres WHERE id = $1
  DELETE FROM reviews WHERE genre_id = $2
	`,
		id, id) 
	if err != nil {
		return -1, err
	}
	return id, nil
}
