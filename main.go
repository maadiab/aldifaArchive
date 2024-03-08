package main

import (
	"context"
	"log"
	"net/http"

	Database "github.com/maadiab/aldifaArchive/database"
	"github.com/maadiab/aldifaArchive/routes"
)

func main() {

	ctx := context.Background()

	_, err := Database.ConnectDB(ctx)

	if err != nil {
		log.Println("Error Connecting Database !!!", err)
	}

	err = Database.CreateUsers()
	if err != nil {
		log.Println("main: Error Creating users Table !!!", err)
	}

	err = Database.CreatePhotographerTable()
	if err != nil {
		log.Println("main: Error Creating photographers Table !!!", err)
	}

	r := routes.Router()

	log.Println("server is running at port: 8080 ...")
	http.ListenAndServe(":8080", r)

	defer Database.DB.Close()

}
