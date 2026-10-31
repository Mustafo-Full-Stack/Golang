package v1

import (
	"crud-api/internal/handler/response"
	"crud-api/internal/models"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func writeError(c *gin.Context, op string, err error) {
	switch {
	case errors.Is(err, models.ErrInvalidInput):
		response.Error(c, http.StatusBadRequest, models.ErrInvalidInput.Error())
	case errors.Is(err, models.ErrPasswordTooShort):
		response.Error(c, http.StatusBadRequest, models.ErrPasswordTooShort.Error())
	case errors.Is(err, models.ErrPasswordTooLong):
		response.Error(c, http.StatusBadRequest, models.ErrPasswordTooLong.Error())
	case errors.Is(err, models.ErrUserNotFound):
		response.Error(c, http.StatusNotFound, models.ErrUserNotFound.Error())
	case errors.Is(err, models.ErrEmailAlreadyExists):
		response.Error(c, http.StatusConflict, models.ErrEmailAlreadyExists.Error())
	case errors.Is(err, models.ErrUserAlreadyExists):
		response.Error(c, http.StatusConflict, models.ErrUserAlreadyExists.Error())
	case errors.Is(err, models.ErrInvalidCredentials):
		response.Error(c, http.StatusUnauthorized, models.ErrInvalidCredentials.Error())
	case errors.Is(err, models.ErrInvalidToken):
		response.Error(c, http.StatusUnauthorized, models.ErrInvalidToken.Error())
	default:
		log.Printf("handler %s: %v", op, err)
		response.Error(c, http.StatusInternalServerError, "internal server error")
	}
}
