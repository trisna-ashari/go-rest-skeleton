package validator_test

import (
	"go-rest-skeleton/pkg/validator"
	"testing"

	ozzoValidation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	validation := validator.New()

	assert.Equal(t, validation, &validator.Validator{})
}

func TestSet(t *testing.T) {
	validation := validator.New()
	validation.
		Set("test", "value", validation.
			AddRule().
			Required().
			IsEmail().
			Apply())

	for _, v := range validation.Rules {
		assert.EqualValues(t, v.Field, "test")

		assert.IsType(t, ozzoValidation.Required, v.Rules[0].Rule)
		assert.IsType(t, is.EmailFormat, v.Rules[1].Rule)
	}
}

func TestValidate_Success(t *testing.T) {
	validation := validator.New()
	validation.
		Set("test", "test@gmail.com", validation.
			AddRule().
			Required().
			IsEmail().
			Apply())

	result := validation.Validate()
	assert.EqualValues(t, 0, len(result))
}

func TestValidate_Error(t *testing.T) {
	validation := validator.New()
	validation.
		Set("test", "value", validation.
			AddRule().
			Required().
			IsEmail().
			Apply())

	result := validation.Validate()
	assert.EqualValues(t, 1, len(result))
}
