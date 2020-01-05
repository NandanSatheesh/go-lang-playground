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
	gorm.Model
	Name         string
	MobileNumber int64 `gorm:"unique;not null"`
	Address      string
}

type MobileNumberOtp struct {
	gorm.Model
	User         User
	UserId       int64 `gorm:"ForeignKey:id"`
	MobileNumber int64
	Otp          int16
}

func GenerateJwt(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Nandan"
	claims["exp"] = time.Now().Unix()

	tokenString, err := token.SignedString(myKey)

	checkErr(err)

	dbHandler, err := gorm.Open("mysql", "root:@/login?charset=utf8&parseTime=True&loc=Local")
	checkErr(err)
	defer dbHandler.Close()

	var user = User{
		Name:         "Alan",
		MobileNumber: 8762171140,
		Address:      "Bangalore",
	}

	var otp = MobileNumberOtp{
		MobileNumber: user.MobileNumber,
		Otp:          6969,
		User:         user,
	}

	u := User{}
	o := MobileNumberOtp{}

	dbHandler.Save(&otp)

	dbHandler.Model(&o).Association("user").Find(&o.User)

	fmt.Println(u)
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
	//dbHandler.Set("gorm:table_options", "ENGINE=InnoDB").DropTableIfExists(&User{}, &MobileNumberOtp{})
	//dbHandler.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}, &MobileNumberOtp{})
	dbHandler.AutoMigrate(&User{}, &MobileNumberOtp{})
	defer dbHandler.Close()
	checkErr(err)
	print(dbHandler, err)
	fmt.Println("Client")
	handleRequests()
}
