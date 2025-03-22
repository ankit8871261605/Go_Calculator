package main

import (
	"Go_Calculator/internal/apps"
	controller "Go_Calculator/internal/controllers"
	"fmt"
)

var User *apps.User

func  init() {
	User = apps.GetUserDetails()
}

func main(){
	fmt.Printf("Operand 1: ")
	var Opr1 int
	fmt.Scan(&Opr1)
	fmt.Printf("Operand 2: ")
	var Opr2 int
	fmt.Scan(&Opr2)
	fmt.Printf("Add Operation ( +,-,*,/): ")
	var Opr string
	fmt.Scan(&Opr)
	con:= controller.GetController(User)
	result,err:=con.Operation(Opr1,Opr2,Opr)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%v %v %v = %v",Opr1,Opr,Opr2,result)
	}
}