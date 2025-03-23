package rest

import (
	"errors"
	"fmt"
	"github.com/DataStarETL/Orchestrator/internal/database"
	"github.com/DataStarETL/Orchestrator/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RestService) GetConditions(c *gin.Context) {
	res, err := database.SelectConditions(r.DB)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (r *RestService) PostCondition(c *gin.Context) {
	var newCondition models.Condition
	if err := c.BindJSON(&newCondition); err != nil {
		fmt.Println(err)
		return
	}
	newCondition.Creator = "SYSTEM"
	err := database.InsertConditionsSpecific(r.DB, &newCondition)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newCondition)
}

func (r *RestService) getConditionByID(id string) (*models.Condition, error) {
	res, err := database.SelectConditionsSpecific(r.DB, id)
	if err != nil {
		panic(err)
	}
	if len(res) == 0 {
		return nil, errors.New("condition not found")
	}
	return &res[0], nil

}

func (r *RestService) GetConditionSpecific(c *gin.Context) {
	id := c.Param("id")
	condition, err := r.getConditionByID(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "condition not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, condition)
}

func (r *RestService) DeleteConditionSpecific(c *gin.Context) {
	id := c.Param("id")
	condition, err := r.getConditionByID(id)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "condition not found"})
		return
	}
	err = database.DeleteConditionsSpecific(r.DB, id)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, condition)
}

func (r *RestService) UpdateCondition(c *gin.Context) {
	var newCondition models.Condition
	if err := c.BindJSON(&newCondition); err != nil {
		fmt.Println(err)
		return
	}
	err := database.UpdateConditionsSpecific(r.DB, &newCondition)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, newCondition)
}
