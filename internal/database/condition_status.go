package database

import (
	"database/sql"
	"fmt"
	"github.com/DataStarETL/Orchestrator/internal/models"
)

func SelectConditionStatus(db *sql.DB) ([]models.ConditionStatus, error) {
	var conditionStatus []models.ConditionStatus

	rows, err := db.Query("SELECT * FROM t_condition_status")
	if err != nil {
		return nil, fmt.Errorf("SelectConditionStatus: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var conditionStatusObj models.ConditionStatus
		if err := rows.Scan(&conditionStatusObj.ID, &conditionStatusObj.Status, &conditionStatusObj.LastChangedBy, &conditionStatusObj.LastChangedAt); err != nil {
			return nil, fmt.Errorf("SelectConditionStatus %v", err)
		}
		conditionStatus = append(conditionStatus, conditionStatusObj)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SelectConditionStatus: %v", err)
	}
	return conditionStatus, nil
}

func SelectConditionStatusSpecific(db *sql.DB, id string) (*models.ConditionStatus, error) {
	var conditionStatus models.ConditionStatus
	row := db.QueryRow("SELECT * FROM t_condition_status WHERE id = $1", id)
	switch err := row.Scan(&conditionStatus.ID, &conditionStatus.Status, conditionStatus.LastChangedBy, conditionStatus.LastChangedAt); err {
	case nil:
		return &conditionStatus, nil
	default:
		return nil, err
	}
}

func InsertConditionStatusSpecific(db *sql.DB, conditionStatus *models.ConditionStatus) error {
	_, err := db.Exec("INSERT INTO  t_condition_status (id, status, LAST_CHANGED_BY, LAST_CHANGED_AT) VALUES ($1, $2, $3, $4)", conditionStatus.ID, conditionStatus.Status, conditionStatus.LastChangedBy, conditionStatus.LastChangedAt)
	if err != nil {
		return fmt.Errorf("InsertConditionStatusSpecific: %v", err)
	}
	return nil
}

func DeleteConditionStatusSpecific(db *sql.DB, id string) error {
	result, err := db.Exec("DELETE FROM t_condition_status WHERE id =$1", id)
	if err != nil {
		return fmt.Errorf("DeleteConditionStatusSpecific: %v", err)
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteConditionStatusSpecific: %v", err)
	}
	fmt.Println(ra)
	return nil
}

func UpdateConditionStatusSpecific(db *sql.DB, conditionStatus *models.ConditionStatus) error {
	result, err := db.Exec("UPDATE t_condition_status SET STATUS=$2, LAST_CHANGED_BY=$3, LAST_CHANGED_AT=$4 WHERE id =$1", conditionStatus.ID, conditionStatus.Status, conditionStatus.LastChangedBy, conditionStatus.LastChangedAt)
	if err != nil {
		return fmt.Errorf("UpdateConditionStatusSpecific: %v", err)
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdateConditionStatusSpecific: %v", err)
	}
	fmt.Println(ra)
	return nil
}
