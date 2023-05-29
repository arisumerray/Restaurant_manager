package main

import (
	"log"
	"order/db"
	"order/internal/dish"
	"order/internal/order"
	"order/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	dishRep := dish.NewRepository(dbConn.GetDB())
	dishSvc := dish.NewService(dishRep)
	dishHandler := dish.NewHandler(dishSvc)

	orderRep := order.NewRepository(dbConn.GetDB())
	orderSvc := order.NewService(orderRep)
	orderHandler := order.NewHandler(orderSvc)

	router.InitRouter(dishHandler, orderHandler)
	err = router.Start("0.0.0.0:8081")
	if err != nil {
		return
	}
}
