package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

var myKey = []byte("secretKey")

func GenerateJwt(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Nandan"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(myKey)

	checkErr(err)

	fmt.Fprintf(w, tokenString)
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(homePage)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	http.HandleFunc("/jwt", GenerateJwt)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Client")
	handleRequests()
}
