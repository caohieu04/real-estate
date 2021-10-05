package main

import (
	"log"
	"os"

	"github.com/caohieu04/real-estate/component/setenv"
	"github.com/caohieu04/real-estate/component/upcloud"
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
	appCtx := component
}
