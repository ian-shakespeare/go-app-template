package examples

import (
	"encoding/json"

	"github.com/ian-shakespeare/go-app-template/internal/models"
)

type Example struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ExampleNew struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type ExampleEdit struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

func Create(e ExampleNew) (Example, error) {
	stmt := `
	INSERT INTO examples (name, description)
	VALUES (?,?)
	`

	res, err := models.Conn().Exec(stmt, e.Name, e.Description)
	if err != nil {
		return Example{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Example{}, err
	}

	return Get(int(id))
}

func Get(id int) (Example, error) {
	query := `
	SELECT json_object(
		'id', e.id
		, 'name', e.name
		, 'description', e.description
	)
	FROM examples e
	WHERE id = ?
	`

	row := models.Conn().QueryRow(query, id)

	var data []byte
	if err := row.Scan(&data); err != nil {
		return Example{}, err
	}

	var ex Example
	err := json.Unmarshal(data, &ex)
	return ex, err
}

func All() ([]Example, error) {
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

	row := models.Conn().QueryRow(query)

	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, err
	}

	var ex []Example
	err := json.Unmarshal(data, &ex)
	return ex, err
}

func Update(id int, e ExampleEdit) (Example, error) {
	stmt := `
	UPDATE examples
	SET
		name = IFNULL(?, name)
		, description = IFNULL(?, description)
	WHERE id = ?
	`

	if _, err := models.Conn().Exec(stmt, e.Name, e.Description, id); err != nil {
		return Example{}, err
	}

	return Get(id)
}

func Delete(id int) error {
	stmt := `
	DELETE FROM examples
	WHERE id = ?
	`

	_, err := models.Conn().Exec(stmt, id)
	return err
}
