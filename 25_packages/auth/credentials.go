package auth

import "fmt"


 func LoginWithCredentials(username string,password string){
	fmt.Println("login user using",username,password)
}

//use caps to start the name of function so that it can be exported.
