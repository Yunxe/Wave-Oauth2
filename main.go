package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"oauth2/database"
	"oauth2/migrator"
	"oauth2/proto/rpc"
	"oauth2/router"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
		panic("Error loading .env file")
	}
	database.MysqlConnection(ctx)
	database.RedisConnection(ctx)
	rpc.InitGrpcClient()
	ifMigrate := flag.Bool("m", false, "recreate table")
	flag.Parse()
	if *ifMigrate {
		migrator.Migrate()
	}
	router.NewRouter()
}
