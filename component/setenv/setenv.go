package setenv

import "os"

func SetEnv() {
	os.Setenv("DBConnectionStr", "real_estate:root123$%^@tcp(127.0.0.1:3306)/real_estate?charset=utf8mb4&parseTime=True&loc=Local")
	os.Setenv("S3BucketName", "real-estate-golang")
	os.Setenv("S3Region", "ap-southeast-1")
	os.Setenv("S3APIKey", "AKIATZPHKNS3T43EIQNP")
	os.Setenv("S3SecretKey", "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")
}
