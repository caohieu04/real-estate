package main

import (
	"log"
	"net/http"
	"os"

	"github.com/caohieu04/real-estate/component"
	"github.com/caohieu04/real-estate/component/setenv"
	"github.com/caohieu04/real-estate/component/upcloud"
	"github.com/caohieu04/real-estate/middleware"
	sellergin "github.com/caohieu04/real-estate/module/seller/transport/gin"
	uploadgin "github.com/caohieu04/real-estate/module/upload/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	setenv.SetEnv()

	dsn := os.Getenv("DBConnectionStr")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := upcloud.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, uploadProvider upcloud.UploadProvider) error {
	appCtx := component.NewAppContext(db, uploadProvider)

	server := gin.Default()

	server.Use(middleware.Recover(appCtx))

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ping success!",
		})
	})

	server.POST("/upload", uploadgin.Upload(appCtx))

	sellers := server.Group("/sellers")
	{
		sellers.POST("", sellergin.CreateSeller(appCtx))
		sellers.GET("/:id", sellergin.GetSeller(appCtx))
		sellers.GET("", sellergin.ListSeller(appCtx))
		sellers.PATCH("/:id", sellergin.UpdateSeller(appCtx))
		sellers.DELETE("/:id", sellergin.DeleteSeller(appCtx))
	}
	return server.Run("103.154.100.225:8080")
	// server.POST("/upload)
}
