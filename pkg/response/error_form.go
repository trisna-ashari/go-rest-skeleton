package response

import (
	"fmt"
	"go-rest-skeleton/pkg/translation"
	"strings"

	"github.com/gin-gonic/gin"
)

// ErrorForm is a struct.
type ErrorForm struct {
	Field string
	Msg   string
	Data  map[string]interface{}
}

// Translate will translate ErrorForm.
func (ef ErrorForm) Translate(c *gin.Context) (string, string) {
	ef.Data["Field"] = ef.TranslateField(c)
	ef.Data["Target"] = ef.TranslateTarget(c)

	message := ef.Msg
	messageData := ef.Data

	return translation.NewTranslation(c, "error", message, messageData)
}

// TranslateField will translate ErrorForm.Field.
func (ef ErrorForm) TranslateField(c *gin.Context) interface{} {
	var errMessage interface{}
	errMessage, _ = translation.NewTranslation(c, "error", fmt.Sprintf("attributes.%s", ef.Field), nil)

	return errMessage
}

// TranslateTarget will translate ErrorForm.Data when key is 'Target' & value has prefix 'attributes'.
// Prefix 'attributes' indicate target Field.
func (ef ErrorForm) TranslateTarget(c *gin.Context) interface{} {
	var errMessage interface{}
	for key, value := range ef.Data {
		if key == "Target" && strings.HasPrefix(value.(string), "attributes") {
			errMessage, _ = translation.NewTranslation(c, "error", value.(string), nil)
		}
	}

	return errMessage
}

// TranslateErrorForm will translate ErrorForm.
func TranslateErrorForm(c *gin.Context, ef []ErrorForm) map[string]string {
	var errMessage = make(map[string]string)
	for _, efMsg := range ef {
		errMessage[efMsg.Field], _ = efMsg.Translate(c)
	}

	return errMessage
}
