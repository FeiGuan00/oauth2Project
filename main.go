package main

import (
	"fmt"
	oauth2Project "oauthProject/jwtutil"
)

func main() {
	util := oauth2Project.JWTUtil{
		EncodingString: "something",
	}

	fmt.Println(util.GetToken("app1", "secret1"))
}
