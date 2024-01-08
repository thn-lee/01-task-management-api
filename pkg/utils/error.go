package utils

import (
	"net/http"
	"strings"

	"github.com/thn-lee/01-task-management-api/pkg/models"
)

func NewError(code int, message ...string) (err *models.Error) {
	if len(message) == 0 {
		message = append(message, http.StatusText(code))
	}
	err = &models.Error{
		Code:    code,
		Source:  WhereAmI(2),
		Title:   http.StatusText(code),
		Message: strings.Join(message, " \n"),
	}
	return
}

func NewErrorSource(code int, source string, message ...string) (err *models.Error) {
	if len(message) == 0 {
		message = append(message, http.StatusText(code))
	}
	err = &models.Error{
		Code:    code,
		Source:  source,
		Title:   http.StatusText(code),
		Message: strings.Join(message, " \n"),
	}
	return
}
