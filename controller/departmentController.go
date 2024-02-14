package departmentcontroller

import (
	"golang_gin_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(ctx *gin.Context) {
	// var dept []models.Department
	// models.DB.Find(&dept)

	ctx.JSON(200, gin.H{
		"department": "Masuk Index",
	})
}

func GetOne(ctx *gin.Context) {
	var dept models.Department
	id := ctx.Param("id")

	if err := models.DB.First(&dept, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	ctx.JSON(200, gin.H{"department": dept})
}

func Insert(ctx *gin.Context) {
	var dept models.Department

	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	models.DB.Create(&dept)
	ctx.JSON(200, gin.H{"department": dept})
}

func Update(ctx *gin.Context) {
	var dept models.Department
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&dept).Where("department_id = ?", id).Updates(&dept).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Data berhasil di update"})
}

func Delete(ctx *gin.Context) {
	var dept models.Department
	input := map[string]string{"id": "0"}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&dept, id).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak ditemukan"})
		return
	}
	ctx.JSON(200, gin.H{"message": "Data berhasil di hapus"})
}
