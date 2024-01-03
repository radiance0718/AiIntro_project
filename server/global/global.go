package global

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	GLO_DB   *gorm.DB
	GLO_SALT string
)

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
