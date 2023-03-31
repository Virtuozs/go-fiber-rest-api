package main

import (
	"fmt"
	"golang-fiber-crud/config"
	"golang-fiber-crud/controller"
	"golang-fiber-crud/model"
	"golang-fiber-crud/repository"
	"golang-fiber-crud/router"
	"golang-fiber-crud/service"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Print("Running Service....")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variable", err)
	}

	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	noteRepository := repository.NewNoteRepositoryImpl(db)

	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	noteController := controller.NewNoteController(noteService)

	routes := router.NewRouter(noteController)

	app := fiber.New()

	app.Mount("/api", routes)

	log.Fatal(app.Listen(":8000"))
}
