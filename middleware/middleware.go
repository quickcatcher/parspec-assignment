package middleware

import (
	"fmt"
	"os"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var dbRegistered sync.Once

func DBConnection(c *gin.Context) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dbRegistered.Do(func() {
		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

		// connectionString := "root:password@tcp(localhost:3306)/local?charset=utf8mb4&loc=Asia%2FShanghai"
		err = orm.RegisterDataBase("default", "mysql", connectionString)
		if err != nil {
			fmt.Println("Error while registering db", err)
			return
		}
	})

	err = MysqlTest(DbName)
	if err != nil {
		fmt.Println("Error while testing mysql connection ", err)
		return
	}
}

func MysqlTest(database string) error {
	o := orm.NewOrm()
	o.Using(database)
	_, err := o.Raw("SELECT 1").Exec()
	return err
}
