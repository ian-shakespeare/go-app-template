package models

import (
	"database/sql"
	"encoding/json"
)

var Examples exampleRepo

type Example struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewExample struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type EditExample struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type exampleRepo struct {
	db *sql.DB
}

func (e *exampleRepo) Create(example NewExample) (Example, error) {
	stmt := `
	INSERT INTO examples (name, description)
	VALUES (?,?)
	`

	res, err := e.db.Exec(stmt, example.Name, example.Description)
	if err != nil {
		return Example{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Example{}, err
	}

	return e.Get(int(id))
}

func (e *exampleRepo) Get(id int) (Example, error) {
	query := `
	SELECT json_object(
		'id', e.id
		, 'name', e.name
		, 'description', e.description
	)
	FROM examples e
	WHERE id = ?
	`

	row := e.db.QueryRow(query, id)

	var data []byte
	if err := row.Scan(&data); err != nil {
		return Example{}, err
	}

	var ex Example
	err := json.Unmarshal(data, &ex)
	return ex, err
}

func (e *exampleRepo) All() ([]Example, error) {
	query := `
	SELECT json_group_array(
		json_object(
			'id', e.id
			, 'name', e.name
			, 'description', e.description
		)
	)
	FROM examples e;
	`

	row := e.db.QueryRow(query)

	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, err
	}

	var ex []Example
	err := json.Unmarshal(data, &ex)
	return ex, err
}

func (e *exampleRepo) Update(id int, example EditExample) (Example, error) {
	stmt := `
	UPDATE examples
	SET
		name = IFNULL(?, name)
		, description = IFNULL(?, description)
	WHERE id = ?
	`

	if _, err := e.db.Exec(stmt, example.Name, example.Description, id); err != nil {
		return Example{}, err
	}

	return e.Get(id)
}

func (e *exampleRepo) Delete(id int) error {
	stmt := `
	DELETE FROM examples
	WHERE id = ?
	`

	_, err := e.db.Exec(stmt, id)
	return err
}
