// package controller

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/rahullodha85/gin-user-app/entity"
// 	"github.com/rahullodha85/gin-user-app/service"
// )

// type VideoController interface {
// 	FindAll() []entity.Video
// 	Save(ctx *gin.Context) entity.Video
// }

// type videoController struct {
// 	service service.VideoService
// }

// func New(service service.VideoService) VideoController {
// 	return videoController {
// 		service: service,
// 	}
// }