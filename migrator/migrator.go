package migrator

import (
	"fmt"
	"oauth2/database"
	"oauth2/model"
)

func Migrate() {
	migrator := database.DB.Migrator()
	if err := migrator.DropTable("clients"); err != nil {
		panic("drop table error" + err.Error())
		return
	}
	if err := migrator.CreateTable(&model.Client{}); err != nil {
		panic("create table error" + err.Error())
		return
	}
	fmt.Println("migrator running")
}
