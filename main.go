package main

import (
	oauth2Project "oauthProject/jwtutil"
)

func main() {
	util := oauth2Project.JWTUtil{
		encodingString: "something",
	}
}
