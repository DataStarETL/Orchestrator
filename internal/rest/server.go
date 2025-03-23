package rest

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type RestService struct {
	DB     *sql.DB
	router *gin.Engine
}

func CreateRestService(db *sql.DB) RestService {
	router := gin.Default()
	service := RestService{DB: db, router: router}
	service.registerV1Endpoints()
	return service
}

func (r *RestService) Start(port string) (err error) {
	return r.router.Run(port)
}

func (r *RestService) registerV1Endpoints() {
	r.router.GET("/v1/conditions", r.GetConditions)
	r.router.POST("/v1/conditions", r.PostCondition)
	r.router.GET("/v1/conditions/:id", r.GetConditionSpecific)
	r.router.PATCH("/v1/conditions/:id", r.UpdateCondition)
	r.router.DELETE("/v1/conditions/:id", r.DeleteConditionSpecific)
}
