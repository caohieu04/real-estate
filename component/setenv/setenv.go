package setenv

import "os"

func SetEnv() {
	os.Setenv("DBConnectionStr", "real_estate:root123$%^@tcp(127.0.0.1:3306)/real_estate?charset=utf8mb4&parseTime=True&loc=Local")
	os.Setenv("")
}