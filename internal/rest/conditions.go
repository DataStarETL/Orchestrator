package rest

import (
	"fmt"
	"github.com/DataStarETL/Orchestrator/internal/database"
	"github.com/DataStarETL/Orchestrator/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (r *RestService) GetConditions(c *gin.Context) {
	res, err := database.SelectConditions(r.DB)
	if err != nil {
		createErrorResponse(c, http.StatusInternalServerError, "Error on database connection.")
		return
	}
	createDataResponse(c, http.StatusOK, res)
}

func (r *RestService) PostCondition(c *gin.Context) {
	var newCondition models.Condition
	if err := c.BindJSON(&newCondition); err != nil {
		createErrorResponse(c, http.StatusBadRequest, "Invalid json data.")
		fmt.Println(err)
		return
	}
	newCondition.Creator = "SYSTEM"
	err := database.InsertConditionsSpecific(r.DB, &newCondition)
	if err != nil {
		createErrorResponse(c, http.StatusInternalServerError, "Failed to save new condition.")
		fmt.Println(err)
		return
	}
	newConditionStatus := models.ConditionStatus{ID: newCondition.ID, Status: false, LastChangedBy: "SYSTEM", LastChangedAt: time.Now()}
	err = database.InsertConditionStatusSpecific(r.DB, &newConditionStatus)
	if err != nil {
		createErrorResponse(c, http.StatusInternalServerError, "Failed to save new condition status.")
		fmt.Println(err)
		return
	}
	createDataResponse(c, http.StatusCreated, newCondition)
}

func (r *RestService) getConditionByID(id string) (*models.Condition, error) {
	return database.SelectConditionsSpecific(r.DB, id)
}

func (r *RestService) GetConditionSpecific(c *gin.Context) {
	id := c.Param("id")
	condition, err := r.getConditionByID(id)
	if err != nil {
		fmt.Println(err)
		createErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Search for Condition with ID %s failed.", id))
		return
	}
	if condition == nil {
		createErrorResponse(c, http.StatusNotFound, fmt.Sprintf("Condition with ID %s not found.", id))
		return
	}
	createDataResponse(c, http.StatusOK, condition)
}

func (r *RestService) DeleteConditionSpecific(c *gin.Context) {
	id := c.Param("id")
	condition, err := r.getConditionByID(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "condition not found"})
		return
	}
	if condition == nil {
		createErrorResponse(c, http.StatusNotFound, fmt.Sprintf("Condition with ID %s not found.", id))
		return
	}
	err = database.DeleteConditionsSpecific(r.DB, id)
	if err != nil {
		createErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Deletion of Condition with ID %s failed.", id))
		fmt.Println(err)
	}

	err = database.DeleteConditionStatusSpecific(r.DB, id)
	if err != nil {
		createErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Deletion of Condition Status with ID %s failed.", id))
		fmt.Println(err)
	}
	createDataResponse(c, http.StatusOK, condition)
}

func (r *RestService) UpdateCondition(c *gin.Context) {
	var newCondition models.Condition
	if err := c.BindJSON(&newCondition); err != nil {
		createErrorResponse(c, http.StatusBadRequest, "Invalid json data.")
		fmt.Println(err)
		return
	}
	err := database.UpdateConditionsSpecific(r.DB, &newCondition)
	if err != nil {
		createErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Update of Condition with ID %s failed.", newCondition.ID))
		fmt.Println(err)
		return
	}
	createDataResponse(c, http.StatusOK, newCondition)
}
