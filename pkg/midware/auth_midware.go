package midware

import (
	"fmt"
	"github.com/mukeshpilaniya/auth/internal/token"
	"github.com/mukeshpilaniya/auth/internal/util"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if len(header)==0{
			w.Write([]byte("Authorization header is not present"))
			return
		}
		s := strings.Split(header," ")
		fmt.Println(len(s))
		if len(s)!=2 || s[0]!="Bearer"{
			w.Write([]byte("Authorization header is not valid"))
			return
		}
		tokenString := string(s[1])
		secretKey := viper.GetString("TOKEN_SECRET_KEY")
		jwtToken, err := token.NewJWTToken(secretKey)
		if err !=nil{
			util.WriteJSON(w,http.StatusNotAcceptable, util.Payload{Message: "internal server error", Error: true},nil)
			return
		}
		isValid, err :=jwtToken.VerifyAccessToken(tokenString)
		if err !=nil{
			util.WriteJSON(w,http.StatusUnauthorized,util.Payload{Message: err.Error(), Error: true},nil)
			return
		}
		if isValid==false{
			util.WriteJSON(w,http.StatusUnauthorized,util.Payload{Message: "unauthorized", Error: true},nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}
