package main

import (
	"log"

	"github.com/Bagussurya12/catalog-music-simple/pkg/internalsql"
	"github.com/Bagussurya12/catalog-music-simple/source/configs"
	membershipHandler "github.com/Bagussurya12/catalog-music-simple/source/handlers/memberships"
	"github.com/Bagussurya12/catalog-music-simple/source/models/memberships"
	membershipsRepo "github.com/Bagussurya12/catalog-music-simple/source/repository/memberships"
	membershipSVC "github.com/Bagussurya12/catalog-music-simple/source/service/memberships"
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
	r := gin.Default()

	db.AutoMigrate(&memberships.User{})
	membershipsRepo := membershipsRepo.NewRepository(db)

	membershipSVC := membershipSVC.NewService(cfg, membershipsRepo)

	membershipHandler := membershipHandler.NewHandler(r, membershipSVC)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
