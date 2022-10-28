package pgsql

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"retelling/pkg/models"
)

func (s *Storage) GetUsers(req models.Request) ([]models.User, error) {
	var data []models.User
	if len(req.ObjectIDs) == 0 && req.ObjectID != 0 {
		req.ObjectIDs = append(req.ObjectIDs, req.ObjectID)
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
		profession,
		likes
	FROM users
		WHERE (id=ANY($1) OR array_length($1,1) is NULL)
	`, intToInt32Array(req.ObjectIDs))
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
			&item.Data.Likes,
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
	var login string
	err := s.pool.QueryRow(context.Background(),
		`SELECT id FROM users WHERE login = $1`, data.Login).Scan(&login)
	if err != pgx.ErrNoRows {
		return -1, errors.New("already exists")
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	user := models.User{
		Login:    data.Login,
		Password: string(pwd),
		Data: models.UserData{
			Age:        data.Data.Age,
			Name:       data.Data.Name,
			Profession: data.Data.Profession,
		},
	}

	err = s.pool.QueryRow(context.Background(), `
	INSERT INTO users (
		name, 
		login,
		password,
		age
	)
	VALUES ($1,$2,$3,$4) RETURNING id
	`,
		user.Data.Name,
		user.Login,
		user.Password,
		user.Data.Age,
	).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) AuthUser(req models.Request) (int, error) {
	var id int
	var pwd string
	err := s.pool.QueryRow(context.Background(), `
	SELECT
		id,
		password
	FROM
		users
	WHERE login=$1;
	`,
		req.Login).Scan(
		&id,
		&pwd,
	)
	if err != nil {
		return -1, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(req.Password))
	if err != nil {
		return -1, errors.New("wrong password")
	}
	return id, nil
}

func (s *Storage) UpdateUser(item models.User) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	UPDATE users
		SET
    		name = $2,
    		age = $3,
    		review_count = $4,
    		rating = $5,
    		profession = $6,
			likes = $7
	WHERE id = $1
	RETURNING id
	`,
		item.ID,
		item.Data.Name,
		item.Data.Age,
		item.Data.ReviewCount,
		item.Data.Rating,
		item.Data.Profession,
		item.Data.Likes).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Storage) DeleteUser(req models.Request) error {
	err := s.pool.QueryRow(context.Background(), `
	UPDATE content SET users_liked = array_remove(users_liked, $1)
	`,
		req.ObjectID).Scan()
	err = s.pool.QueryRow(context.Background(), `
	DELETE FROM reviews WHERE user_id = $1
	`,
		req.ObjectID).Scan()
	err = s.pool.QueryRow(context.Background(), `
	DELETE FROM users WHERE id = $1
	`,
		req.ObjectID).Scan()
	return err
}

func (s *Storage) UpdatePassword(req models.Request) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(pwd)
	_, err = s.pool.Exec(context.Background(), `
	UPDATE users SET password = $1 WHERE id = $2
	`,
		req.Password,
		req.UserID)
	return err
}
