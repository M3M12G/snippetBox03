package postgresql

import (
	"database/sql"
	"errors"
	"se03.com/pkg/models"
	"strconv"
	"time"
)

type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := "INSERT INTO snippets (title, content, created, expires) VALUES ($1, $2, $3, $4) RETURNING id"

	/*//as this driver does not support lastInsertedId() method,
	then we should use sql.DB.QueryRow().Scan() methods to receive that last id
	below the code from the book
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil{
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil{
		return 0, err
	}

	return int(id), nil
	*/

	//может быть это костыли, но она работает
	created := time.Now()
	day, _ := strconv.Atoi(expires)
	expiresAt := created.AddDate(0, 0, day)

	id := 0
	err := m.DB.QueryRow(stmt, title, content, created, expiresAt).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id, nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	s := &models.Snippet{}
	err := m.DB.QueryRow("SELECT id, title, content, created, expires FROM snippets WHERE expires > NOW() AND id = $1", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
