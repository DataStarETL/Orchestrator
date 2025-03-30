package rest

import (
	"fmt"
	"github.com/DataStarETL/Orchestrator/internal/database"
	"github.com/DataStarETL/Orchestrator/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (r *RestService) GetConditionStatus(c *gin.Context) {
	res, err := database.SelectConditionStatus(r.DB)
	if err != nil {
		fmt.Println(err)
		createErrorResponse(c, http.StatusInternalServerError, "Search for Condition Status failed.")
	}
	createDataResponse(c, http.StatusOK, res)
}

func (r *RestService) getConditionStatusByID(id string) (*models.ConditionStatus, error) {
	return database.SelectConditionStatusSpecific(r.DB, id)
}

func (r *RestService) GetConditionStatusSpecific(c *gin.Context) {
	id := c.Param("id")
	condition, err := r.getConditionStatusByID(id)
	if err != nil {
		fmt.Println(err)
		createErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Search for Condition Status with ID %s failed.", id))
		return
	}
	createDataResponse(c, http.StatusOK, condition)
}

func (r *RestService) UpdateConditionStatus(c *gin.Context) {
	var newConditionStatus models.ConditionStatus
	if err := c.BindJSON(&newConditionStatus); err != nil {
		createErrorResponse(c, http.StatusBadRequest, "Invalid json data.")
		fmt.Println(err)
		return
	}
	newConditionStatus.LastChangedBy = "SYSTEM"
	newConditionStatus.LastChangedAt = time.Now()
	err := database.UpdateConditionStatusSpecific(r.DB, &newConditionStatus)
	if err != nil {
		createErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Update for Condition Status with ID %s failed.", newConditionStatus.ID))
		fmt.Println(err)
		return
	}
	createDataResponse(c, http.StatusOK, newConditionStatus)
}
