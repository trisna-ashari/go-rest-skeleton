package exception

import (
	"github.com/gin-gonic/gin"
	"go-rest-skeleton/pkg/translation"
)

// ErrorForm is a struct.
type ErrorForm struct {
	Field string
	Msg   string
	Data  map[string]interface{}
}

// Translate will translate ErrorForm.
func (ef ErrorForm) Translate(c *gin.Context) (string, string) {
	return translation.NewTranslation(c, "error", ef.Msg, ef.Data)
}

// TranslateErrorForm will translate ErrorForm.
func TranslateErrorForm(c *gin.Context, ef []ErrorForm) map[string]string {
	var errMessage = make(map[string]string)
	for _, efData := range ef {
		errMessage[efData.Field], _ = efData.Translate(c)
	}

	return errMessage
}
