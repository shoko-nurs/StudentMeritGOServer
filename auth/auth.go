package auth

import (
	"StudentMerit/Structures"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

func Authenticate(r *http.Request) (uint64,error){
	var err = Structures.ERROR{Err:"Access Denied"}
	bearer := r.Header.Get("Authorization")

	// This list in form [Bearer 2132]
	list := strings.Split(bearer," ")

	if bearer==""|| len(list)!=2 || list[0]!="Bearer"{
		return 0,err
	}

	tokenStr := list[1]


	type MyClaims struct {
		Id uint64 `json:"id"`
		jwt.RegisteredClaims
	}

	token, _ := jwt.ParseWithClaims(tokenStr,
										&MyClaims{},
										func(token *jwt.Token)(interface{},error){
											return []byte(os.Getenv("SECRET_KEY")), nil
										})

	// Convert jwt.Claims into MyCustomStructure
	claims,ok1 := token.Claims.(*MyClaims)
	ok2 := token.Valid


	if ok1 && ok2 {
		return claims.Id,nil
	}

	return 0,err

}