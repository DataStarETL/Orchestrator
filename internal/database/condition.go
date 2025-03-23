package database

import (
	"database/sql"
	"fmt"
	"github.com/DataStarETL/Orchestrator/internal/models"
)

func SelectConditions(db *sql.DB) ([]models.Condition, error) {
	var conditions []models.Condition

	rows, err := db.Query("SELECT * FROM t_conditions")
	if err != nil {
		return nil, fmt.Errorf("SelectConditions: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var con models.Condition
		if err := rows.Scan(&con.ID, &con.Name, &con.Creator, &con.CreatedAt); err != nil {
			return nil, fmt.Errorf("SelectConditions %v", err)
		}
		conditions = append(conditions, con)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SelectConditions: %v", err)
	}
	return conditions, nil
}

func SelectConditionsSpecific(db *sql.DB, id string) (*models.Condition, error) {
	var condition models.Condition
	row := db.QueryRow("SELECT * FROM t_conditions WHERE id = $1", id)
	switch err := row.Scan(&condition.ID, &condition.Name, &condition.Creator, &condition.CreatedAt); err {
	case nil:
		return &condition, nil
	default:
		return nil, err
	}
}

func InsertConditionsSpecific(db *sql.DB, condition *models.Condition) error {
	_, err := db.Exec("INSERT INTO  t_conditions (id, name, creator, created_at) VALUES ($1, $2, $3, now())", condition.ID, condition.Name, condition.Creator)
	if err != nil {
		return fmt.Errorf("InsertConditionsSpecific: %v", err)
	}
	return nil
}

func DeleteConditionsSpecific(db *sql.DB, id string) error {
	result, err := db.Exec("DELETE FROM t_conditions WHERE id =$1", id)
	if err != nil {
		return fmt.Errorf("DeleteConditionsSpecific: %v", err)
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteConditionsSpecific: %v", err)
	}
	fmt.Println(ra)
	return nil
}

func UpdateConditionsSpecific(db *sql.DB, condition *models.Condition) error {
	result, err := db.Exec("UPDATE t_conditions SET NAME=$2 WHERE id =$1", condition.ID, condition.Name)
	if err != nil {
		return fmt.Errorf("UpdateConditionsSpecific: %v", err)
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdateConditionsSpecific: %v", err)
	}
	fmt.Println(ra)
	return nil
}
