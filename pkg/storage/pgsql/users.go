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
	if len(req.ObjectIDs) == 0 {
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
		profession
	FROM users
		WHERE (id=ANY($1) OR array_length($1) is NULL)
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
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}

func (s *Storage) NewUser(item models.User) (int, error) {
	var id int
	var login string
	err := s.pool.QueryRow(context.Background(),
		`SELECT id FROM users WHERE login = $1`, item.Login).Scan(&login)
	if err != pgx.ErrNoRows {
		return -1, errors.New("Already exists")
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(item.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	user := models.User{
		Login:    item.Login,
		Password: string(pwd),
		Data: &models.UserData{
			Age:        item.Data.Age,
			Name:       item.Data.Name,
			Profession: item.Data.Profession,
		},
	}

	err = s.pool.QueryRow(context.Background(), `
	INSERT INTO users (
		name, 
		login,
		password,
		age,
		profession
	)
	VALUES ($1,$2,$3,$4) RETURNING id
	`,
		user.Data.Name,
		user.Login,
		string(user.Password),
		user.Data.Age,
		user.Data.Profession,
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
		return -1, errors.New("Wrong password")
	}
	return id, nil
}

func (s *Storage) UpdateUser(item models.User) (int, error) {
	var id int
	err := s.pool.QueryRow(context.Background(), `
	UPDATE users
		SET
    		name = $2,
    		login = $3,
    		password = $4,
    		age = $5,
    		review_count = $6,
    		rating = $7,
    		profession = $8
	WHERE id = $1
	RETURNING id
	`,
		item.ID,
		item.Data.Name
		item.Login,
		item.Password,
		item.Data.Age,
		item.Data.ReviewCount,
		item.Data.Rating,
		item.Data.Profession,
		item.Date).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil	
}

func (s *Storage) DeleteUser(id int) (int, error) {
	err := s.pool.QueryRow(context.Background(), `
	DELETE FROM users WHERE id = $1
	DELETE FROM reviews WHERE user_id = $2
	`,
		id, id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
