package validator_test

import (
	"go-rest-skeleton/pkg/validator"
	"regexp"
	"testing"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	ozzoValidation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddRule(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule()

	assert.Equal(t, rules, &validator.ValidationRules{})
}

func TestApply(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().Apply()

	assert.IsType(t, rules, []validator.ValidationRule{})
	assert.Equal(t, rules, []validator.ValidationRule(nil))
}

func TestValidationRules_Required(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().Required().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.Required)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_Empty(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().Empty().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.Empty)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_Nil(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().Nil().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.Nil)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_NotNil(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().NotNil().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.NotNil)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_NilOrEmpty(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().NilOrEmpty().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.NilOrNotEmpty)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_In(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().In("a", "b").Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.InRule{})
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt{{
			Key:   "Options",
			Value: "a/b",
		}})
	}
}

func TestValidationRules_NotIn(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().NotIn("a", "b").Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.NotInRule{})
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt{{
			Key:   "Options",
			Value: "a/b",
		}})
	}
}

func TestValidationRules_Min(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().MinValue(10).Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.ThresholdRule{})
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt{{
			Key:   "Length",
			Value: 10,
		}})
	}
}

func TestValidationRules_Max(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().MaxValue(255).Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.ThresholdRule{})
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt{{
			Key:   "Length",
			Value: 255,
		}})
	}
}

func TestValidationRules_Length(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().Length(10, 255).Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.LengthRule{})
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt{{
			Key:   "Min",
			Value: 10,
		}, {
			Key:   "Max",
			Value: 255,
		}})
	}
}

func TestValidationRules_EqualTo(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().EqualTo("field_name", "value").Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.By(validator.EqualValue("field_name")))
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt{{
			Key:   "Target",
			Value: "attributes.field_name",
		}})
	}
}

func TestValidationRules_IsAlpha(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsAlpha().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.Alpha)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsAlphaSpace(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsAlphaSpace().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.Match(regexp.MustCompile("^[a-zA-Z\\s]*$")))
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsAlphaNumeric(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsAlphaNumeric().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.Alphanumeric)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsAlphaNumericSpace(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsAlphaSpace().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, ozzoValidation.Match(regexp.MustCompile("^[a-zA-Z0-9\\s]*$")))
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsDigit(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsDigit().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.Digit)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsUUID(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsUUID().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.UUID)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsInt(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsInt().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.Int)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsFloat(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsFloat().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.Float)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsEmail(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsEmail().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.EmailFormat)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsPhone(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsPhone().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.Digit)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsURL(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsURL().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.URL)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}

func TestValidationRules_IsJSON(t *testing.T) {
	validation := validator.New()
	rules := validation.AddRule().IsJSON().Apply()

	for _, r := range rules {
		assert.IsType(t, r.Rule, is.JSON)
		assert.Equal(t, r.RuleOpt, []validator.RuleOpt(nil))
	}
}
