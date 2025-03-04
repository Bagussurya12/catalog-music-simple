package main

import (
	"log"

	"github.com/Bagussurya12/catalog-music-simple/pkg/internalsql"
	"github.com/Bagussurya12/catalog-music-simple/source/configs"
	"github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./source/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Failed Configuration:", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed connect to database, err: %+v", err)
	}

	db.AutoMigrate(&memberships.User{})

	r := gin.Default()
	r.Run(cfg.Service.Port)
}
