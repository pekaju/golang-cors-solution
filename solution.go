package main

import (
    "log"
    "net/http"
    "github.com/rs/cors"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "myApp/src/controllers"
)

func main() {
    ac := new(controllers.AccountController)

    router := mux.NewRouter()
    router.HandleFunc("/signup", ac.SignUp).Methods("POST")
    router.HandleFunc("/signin", ac.SignIn).Methods("POST")

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8000"},
        AllowCredentials: true,
    })

    handler := c.Handler(router)
    log.Fatal(http.ListenAndServe(":3000", handler)
}
