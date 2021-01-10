package postgresql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"se03.com/pkg/models"
	"strconv"
	"time"
)

type SnippetModel struct {
	DB *pgxpool.Pool
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := "INSERT INTO snippets (title, content, created, expires) VALUES ($1, $2, $3, $4) RETURNING id"

	/*as this driver does not support lastInsertedId() method,
	then we should use sql.DB.QueryRow().Scan() methods to receive that last id
	*/

	//может быть это костыли, но она работает
	created := time.Now()
	day, _ := strconv.Atoi(expires)
	expiresAt := created.AddDate(0, 0, day)

	id := 0
	err := m.DB.QueryRow(context.Background(), stmt, title, content, created, expiresAt).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id, nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	s := &models.Snippet{}
	err := m.DB.QueryRow(context.Background(), "SELECT id, title, content, created, expires FROM snippets "+
		"WHERE expires > NOW() AND id = $1", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := "SELECT id, title, content, created, expires FROM snippets " +
		"WHERE expires > NOW() ORDER BY created DESC LIMIT 10"

	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
