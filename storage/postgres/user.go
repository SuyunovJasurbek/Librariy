package postgres

import (
	"fmt"
	"library/model"
	"time"

	"github.com/jmoiron/sqlx"
)

type userRepoImpl struct {
	db *sqlx.DB
}


func ( h userRepoImpl) GetAllSearchUser(offset string, limit string, search string) (*model.GetAllUser, error) {

	var (
		resp   model.GetAllUser
		err    error
		filter string
		params = make(map[string]interface{})
	)


	if search != "" {
		filter += " AND name ILIKE '%' || :search || '%' OR owner ILIKE '%' || :search || '%' "
		params["search"] = search
	}
	countQuery := `SELECT count(1) FROM books WHERE true ` + filter

	q, err := h.db.NamedQuery(countQuery, params)
	if err != nil {
		return nil, err
	}
	if q.Rows.Next() {
		q.Rows.Scan(&resp.Count)
	}

	query := `SELECT
				username,
				email,
				age,
				id,
				createdat,
				updatedat
			FROM users
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = limit
	params["offset"] = offset

	rows, err := h.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.Users

		err = rows.Scan(
			&user.Username,
			&user.Email,
			&user.Age,
			&user.ID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.User = append(resp.User, user)

	}

	return &resp, nil

}

func (h userRepoImpl) CreateUser(entity model.Users, id string) error {

	query := `INSERT INTO users (username, email, age, id, createdat, updatedat) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := h.db.Exec(query, entity.Username, entity.Email, entity.Age, id, entity.CreatedAt, entity.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (h userRepoImpl) GetUserName(id string) (string, error) {
	var user model.Users
	query := `SELECT * FROM users WHERE id=$1;`
	row := h.db.QueryRow(query, id)
	err := row.Scan(
		&user.Username,
		&user.Email,
		&user.Age,
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
func (h userRepoImpl) UpdateUser(entity model.UbdateUserRequest, id string) (*model.Users, error) {

	fmt.Println(entity)
	var (
		params       map[string]interface{}
		updated_user model.Users
	)
	params = map[string]interface{}{}
	query := "UPDATE users  SET "

	if len(entity.Username) > 0 {
		params["username"] = entity.Username
	}

	if len(entity.Email) > 0 {
		params["email"] = entity.Email
	}

	if len(entity.Age) > 0 {
		params["age"] = entity.Age
	}
	params["id"] = id
	params["updatedat"] = time.Now()
	k := len(params)
	for i := range params {
		query += fmt.Sprintf(" %s = :%s ", i, i)
		k--
		if k != 0 {
			query += ","
		}
	}
	// query += " WHERE id = :id RETURNING name, owner, cost,updatedat ;" // kerakligini qaytarsih
	query += " WHERE id = :id RETURNING * ;" // xamma update bulgan datani qaytarish
	fmt.Println(query)
	rows, err := h.db.NamedQuery(query, params)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// fmt.Println("res suc")
	for rows.Next() {
		err := rows.Scan(
			&updated_user.Username,
			&updated_user.Email,
			&updated_user.Age,
			&updated_user.ID,
			&updated_user.CreatedAt,
			&updated_user.UpdatedAt,
		)
		if err != nil {
			fmt.Println("scan err")
			return nil, err
		}
	}
	//  fmt.Println("scan suc")
	return &updated_user, nil

}
func (h userRepoImpl) DeleteUser(id string) error {
	var user model.Users
	quer := `SELECT * FROM users WHERE id=$1;`
	row := h.db.QueryRow(quer, id)
	err := row.Scan(
		&user.Username,
		&user.Email,
		&user.Age,
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	query := `DELETE  FROM users WHERE id=$1;`
	_, err2 := h.db.Query(query, id)
	if err2 != nil {
		return err
	}
	return nil
}
