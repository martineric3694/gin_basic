package main

import (
	"fmt"
	departmentcontroller "golang_gin_api/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("GIN Framework")

	r := gin.Default()
	// models.ConnectDB()
	r.GET("/", homePage)
	// r.GET("/data/:pesan", pathParam)
	// r.POST("/", postData)

	r.GET("/departments", departmentcontroller.Index)
	// r.GET("/department/:id", departmentcontroller.GetOne)
	// r.POST("/department", departmentcontroller.Insert)
	// r.PUT("/department/:id", departmentcontroller.Update)
	// r.DELETE("/department/", departmentcontroller.Delete)
	r.Run(":1234")
}

func homePage(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World GIN Framework",
	})
}

func postData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": ctx.Query("pesan"),
	})
}

func pathParam(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": ctx.Param("pesan"),
	})
}
