package main

import (
	"auth/db"
	"auth/internal/user"
	"auth/router"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	err = router.Start("0.0.0.0:8080")
	if err != nil {
		return
	}
}
