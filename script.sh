#Global install
docker pull mysql/mysql-server
docker run -d --name mysql --privileged=true -e MYSQL_ROOT_PASSWORD="root123$%^" -e MYSQL_USER="real_estate" -e MYSQL_PASSWORD="root123$%^" -e MYSQL_DATABASE="real_estate" -p 3306:3306 mysql/mysql-server
#docker exec -it mysql mysql -uroot -p

#Go install 
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm