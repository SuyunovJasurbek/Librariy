package postgres

import (
	"fmt"
	"library/model"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type bookRepoImpl struct {
	db *sqlx.DB
}

var (
	booksTable = "books"
)

// GetAllSearchBooks implements storage.BookRepoI
func (h bookRepoImpl) GetAllSearchBooks(offset string, limit string, search string) (*model.GetAllBook, error) {
	var (
		resp   model.GetAllBook
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
				name,
				owner,
				cost,
				id,
				createdat,
				updatedat
			FROM books
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = limit
	params["offset"] = offset

	rows, err := h.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book model.Books

		err = rows.Scan(
			&book.Name,
			&book.Owner,
			&book.Cost,
			&book.ID,
			&book.CreatedAt,
			&book.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Book = append(resp.Book, book)
	}

	return &resp, nil
}

func (h bookRepoImpl) CreateBook(entity model.Books, id string) error {
	query := fmt.Sprintf(`INSERT INTO %s (name, owner, cost, id,createdat,updatedat) VALUES ($1, $2, $3, $4, $5, $6)`, booksTable)
	_, err := h.db.Exec(query, entity.Name, entity.Owner, entity.Cost, id, entity.CreatedAt, entity.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (h bookRepoImpl) GetBookName(id string) (string, error) {
	var book model.Books
	query := `SELECT * FROM books WHERE id=$1;`
	row := h.db.QueryRow(query, id)
	err := row.Scan(
		&book.Name,
		&book.Owner,
		&book.Cost,
		&book.ID,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	fmt.Println(book.Name)
	if err != nil {
		return "", err
	}
	return book.Name, nil
}

func (h bookRepoImpl) UpdateBook(entity model.UbdateBookRequest, id string) (*model.Books, error) {

	var (
		params       map[string]interface{}
		updated_book model.Books
	)
	params = map[string]interface{}{}
	//  fmt.Println("HEllo from repo")
	query := "UPDATE books  SET "

	if len(entity.Name) > 0 {
		params["name"] = entity.Name
		//query += "firstname = :firstname, "
	}

	if len(entity.Owner) > 0 {
		params["owner"] = entity.Owner
		//query += "lastname = :lastname, "
	}

	if len(entity.Cost) > 0 {
		params["cost"] = entity.Cost
		//query += "username = :username, "
	}
	params["id"] = id
	params["updatedat"] = time.Now()
	k := len(params)
	for i := range params {
		//  fmt.Printf(" %s = :%s ", i, i)
		query += fmt.Sprintf(" %s = :%s ", i, i)
		k--
		if k != 0 {
			query += ","
		}
	}
	params["id"] = id
	params["updatedat"] = time.Now()
	//s := "UPDATE author SET firstname=:firstname, lastname=:lastname"
	// query += " WHERE id = :id RETURNING name, owner, cost,updatedat ;" // ubdet qilganimizdan sung , uni biror nimasini qaytarmoqchi bulsak, shundan foydalanish kerak
	query += " WHERE id = :id RETURNING * ;" // xamma update bulgan datani qayatrishda ishlatishimiz mumkin
	fmt.Println(query)
	rows, err := h.db.NamedQuery(query, params)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// fmt.Println("res suc")
	for rows.Next() {
		err := rows.Scan(
			&updated_book.Name,
			&updated_book.Owner,
			&updated_book.Cost,
			&updated_book.ID,
			&updated_book.CreatedAt,
			&updated_book.UpdatedAt,
		)
		if err != nil {
			//    fmt.Println("scan err")
			return nil, err
		}
	}
	//  fmt.Println("scan suc")
	return &updated_book, nil

}

func (h bookRepoImpl) DeleteBook(id string) error {
	var book model.Books
	quer := `SELECT * FROM books WHERE id=$1;`
	row := h.db.QueryRow(quer, id)
	err := row.Scan(
		&book.Name,
		&book.Owner,
		&book.Cost,
		&book.ID,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		return err
	}

	query := `DELETE  FROM books WHERE id=$1;`
	fmt.Println("------------------", id)
	_, err2 := h.db.Query(query, id)
	if err2 != nil {
		return err
	}
	return nil
}
