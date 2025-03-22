package service

import (
	"Go_Calculator/internal/apps"
	model "Go_Calculator/internal/models"
	"errors"
)

type History struct {
	User *apps.User
	Data model.History
}

func GetAll(app *apps.User) *History {
	return &History{
		User: app,
	}
}

func (h History) Operation(operand1 int, operand2 int, operator string) (int, error) {
    var result int
    switch operator {
    case "+":
        result = operand1 + operand2
    case "-":
        result = operand1 - operand2
    case "*":
        result = operand1 * operand2
    case "/":
        result = operand1 / operand2
    default:
        return 0, errors.New("invalid operation")
    }

	h.Data.Operand1 = operand1
	h.Data.Operand2 = operand2
	h.Data.Operator = operator
	h.Data.Result = result

    return result, nil
}

