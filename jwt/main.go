package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

var myKey = []byte("secretKey")

var tokenKey = []byte("secretKey")

// Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `User`
type User struct {
	Id          uint64 `gorm:"AUTO_INCREMENT"`
	Name        string
	PhoneNumber int64
	Address     string
}

func GenerateJwt(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Nandan"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(myKey)

	checkErr(err)

	dbHandler, err := gorm.Open("mysql", "root:@/login?charset=utf8&parseTime=True&loc=Local")
	checkErr(err)
	defer dbHandler.Close()

	var user = User{
		Name:        "Nandan",
		PhoneNumber: 8762171130,
		Address:     "Bangalore",
	}
	dbHandler.Create(&user)

	fmt.Fprintf(w, tokenString)
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (i interface{}, err error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return myKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, "Error in JWT Token")
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
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
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	dbHandler, err := gorm.Open("mysql", "root:@/login?charset=utf8&parseTime=True&loc=Local")
	dbHandler.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	defer dbHandler.Close()
	checkErr(err)
	print(dbHandler, err)
	fmt.Println("Client")
	handleRequests()
}
