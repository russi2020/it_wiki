package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"it_wiki/app"
	"it_wiki/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")

	//Category routes
	router.HandleFunc("/api/it_wiki/categories", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/api/it_wiki/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/api/it_wiki/categories/{id}", controllers.GetCategory).Methods("GET")

	//Theme routes
	router.HandleFunc("/api/it_wiki/themes", controllers.CreateTheme).Methods("POST")
	router.HandleFunc("/api/it_wiki/themes", controllers.GetThemes).Methods("GET")
	router.HandleFunc("/api/it_wiki/themes/{id}", controllers.GetTheme).Methods("GET")

	//Question routes
	router.HandleFunc("/api/it_wiki/questions", controllers.CreateQuestion).Methods("POST")
	router.HandleFunc("/api/it_wiki/questions", controllers.GetQuestions).Methods("GET")
	router.HandleFunc("/api/it_wiki/questions/{id}", controllers.GetQuestion).Methods("GET")

	//Answer routes
	router.HandleFunc("/api/it_wiki/answers", controllers.CreateAnswer).Methods("POST")
	router.HandleFunc("/api/it_wiki/answers", controllers.GetAnswers).Methods("GET")
	router.HandleFunc("/api/it_wiki/answers/{id}", controllers.GetAnswer).Methods("GET")

	router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена

	http.Handle("/", router)
	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}
