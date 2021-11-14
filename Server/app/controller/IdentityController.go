package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OnlyM1ss/transport-manager/v2/app/model"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	SECRET = "42isTheAnswer"
)
type JWTData struct {
	// Standard claims are the standard jwt claims from the IETF standard
	// https://tools.ietf.org/html/rfc7519
	jwt.StandardClaims
	CustomClaims map[string]string `json:"custom,omitempty"`
}
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Login failed!", http.StatusUnauthorized)
	}

	var userData map[string]string
	var account model.User
	json.Unmarshal(body, &userData)

	filter := bson.D{{"Email", userData["email"]}, {"Password", userData["password"]}}
	client.Database("admin").Collection("User").FindOne(context.TODO(), filter).Decode(&account)
	if account.Email != "" {
		claims := JWTData{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour).Unix(),
			},

			CustomClaims: map[string]string{
				"userid": "u1",
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(SECRET))
		if err != nil {
			log.Println(err)
			http.Error(w, "Login failed!", http.StatusUnauthorized)
		}

		json, err := json.Marshal(struct {
			Token string `json:"token"`
		}{
			tokenString,
		})

		if err != nil {
			log.Println(err)
			http.Error(w, "Login failed!", http.StatusUnauthorized)
		}

		w.Write(json)
	} else {
		http.Error(w, "Login failed!", http.StatusUnauthorized)
	}
}
func Account(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	authArr := strings.Split(authToken, " ")

	if len(authArr) != 2 {
		log.Println("Authentication header is invalid: " + authToken)
		http.Error(w, "Request failed!", http.StatusUnauthorized)
	}

	jwtToken := authArr[1]

	claims, err := jwt.ParseWithClaims(jwtToken, &JWTData{}, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, errors.New("Invalid signing algorithm")
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "Request failed!", http.StatusUnauthorized)
	}

	data := claims.Claims.(*JWTData)

	userID := data.CustomClaims["userid"]

	// fetch some data based on the userID and then send that data back to the user in JSON format
	jsonData, err := GetAccountDataById(userID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Request failed!", http.StatusUnauthorized)
	}

	w.Write(jsonData)
}

// GetAccountDataById Get account from db by id
func GetAccountDataById(userID string) ([]byte, error) {
	// TODO убрать захардкоженные значения
	output := model.User{"admin@gmail.com","admin123","admin","Никита"}
	json, err := json.Marshal(output)
	if err != nil {
		return nil, err
	}

	return json, nil
}