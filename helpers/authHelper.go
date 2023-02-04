package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role bool) (err error) {
	seller := c.GetBool("seller")
	err = nil
	if seller != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return err
}

