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
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (r *RestService) getConditionStatusByID(id string) (*models.ConditionStatus, error) {
	return database.SelectConditionStatusSpecific(r.DB, id)
}

func (r *RestService) GetConditionStatusSpecific(c *gin.Context) {
	id := c.Param("id")
	condition, err := r.getConditionStatusByID(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "condition not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, condition)
}

func (r *RestService) UpdateConditionStatus(c *gin.Context) {
	var newConditionStatus models.ConditionStatus
	if err := c.BindJSON(&newConditionStatus); err != nil {
		fmt.Println(err)
		return
	}
	newConditionStatus.LastChangedBy = "SYSTEM"
	newConditionStatus.LastChangedAt = time.Now()
	err := database.UpdateConditionStatusSpecific(r.DB, &newConditionStatus)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, newConditionStatus)
}
