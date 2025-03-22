package controller

import (
	"Go_Calculator/internal/apps"
	service "Go_Calculator/internal/services"
)

type controller struct {
	User *apps.User
	History service.History
}

func GetController(app *apps.User) *controller{
	return  &controller{
		User: app,
		History: *service.GetAll(app),
	}
}

func  (c *controller) Operation(operand1 int, operand2 int, operator string) (int, error) {
	return c.History.Operation(operand1, operand2 , operator)
}