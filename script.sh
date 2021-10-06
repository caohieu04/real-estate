#Global install
docker pull mysql/mysql-server
docker run -d --name mysql --privileged=true -e MYSQL_ROOT_PASSWORD="root123$%^" -e MYSQL_USER="real_estate" -e MYSQL_PASSWORD="root123$%^" -e MYSQL_DATABASE="real_estate" -p 3306:3306 mysql/mysql-server
#to run mysql inside: "docker exec -it mysql mysql -uroot -p"
docker cp /content/real-estate/misc/Initialization.sql mysql:/
docker exec -it mysql bin/bash
#bash-4.4# mysql -u root -p
#mysql> USE real_estate;
#mysql> source Initialization.sql

#Go install 
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/aws/aws-sdk-go/