package validator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"gorm.io/gorm"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidationRule is a struct represent validation rule.
type ValidationRule struct {
	Rule    validation.Rule
	RuleKey string
	RuleOpt []RuleOpt
}

// ValidationRules is a struct to store multiple validation rule.
type ValidationRules struct {
	Rules []ValidationRule
}

// AddRule is a constructor will initialize ValidationRules.
func (v *Validator) AddRule() *ValidationRules {
	return &ValidationRules{}
}

// Apply is a function uses to apply validation.Rule were has been stored in ValidationRules.
func (vr *ValidationRules) Apply() []ValidationRule {
	return vr.Rules
}

// When is a function to set the rule that current field when meet predefined condition.
func (vr *ValidationRules) When(condition bool, rules *ValidationRules) *ValidationRules {
	for _, rule := range rules.Rules {
		vr.Rules = append(vr.Rules, ValidationRule{
			Rule:    validation.When(condition, rule.Rule),
			RuleOpt: nil,
		})
	}

	return vr
}

// Required is a function to set the rule that current field is required.
func (vr *ValidationRules) Required() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    validation.Required.Error("api.msg.error.validation.is_required"),
		RuleOpt: nil,
	})
	return vr
}

// Empty is a function to set the rule that current field value is empty.
func (vr *ValidationRules) Empty() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    validation.Empty,
		RuleOpt: nil,
	})
	return vr
}

// Nil is a function to set the rule that current field value is nil.
func (vr *ValidationRules) Nil() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    validation.Nil,
		RuleOpt: nil,
	})
	return vr
}

// NotNil is a function to set the rule that current field value is not nil.
func (vr *ValidationRules) NotNil() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    validation.NotNil,
		RuleOpt: nil,
	})
	return vr
}

// NilOrEmpty is a function to set the rule that current field value is nil or empty.
func (vr *ValidationRules) NilOrEmpty() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    validation.NilOrNotEmpty,
		RuleOpt: nil,
	})
	return vr
}

// In is a function to set the rule that current field value must on of slices.
func (vr *ValidationRules) In(slice ...interface{}) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.In(slice...).Error("api.msg.error.validation.must_be_in"),
		RuleOpt: []RuleOpt{
			{
				Key:   "Options",
				Value: joinSlice(slice),
			},
		},
	})
	return vr
}

// NotIn is a function to set the rule that current field value is not on of slices.
func (vr *ValidationRules) NotIn(slice ...interface{}) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.NotIn(slice).Error("api.msg.error.validation.must_be_not_in"),
		RuleOpt: []RuleOpt{
			{
				Key:   "Options",
				Value: joinSlice(slice),
			},
		},
	})
	return vr
}

// MinValue is a function to set the rule that current field value must be no less than the length.
func (vr *ValidationRules) MinValue(length interface{}) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Min(length).Error("api.msg.error.validation.must_be_no_less_than_value"),
		RuleOpt: []RuleOpt{
			{
				Key:   "Length",
				Value: length,
			},
		},
	})
	return vr
}

// MinValue is a function to set the rule that current field value must be no less than the length.
func (vr *ValidationRules) MinLength(length interface{}) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Min(length).Error("api.msg.error.validation.must_be_no_less_than_length"),
		RuleOpt: []RuleOpt{
			{
				Key:   "Length",
				Value: length,
			},
		},
	})
	return vr
}

// MaxValue is a function to set the rule that current field value must be no more than the length.
func (vr *ValidationRules) MaxValue(length interface{}) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Max(length).Error("api.msg.error.validation.must_be_no_more_than_value"),
		RuleOpt: []RuleOpt{
			{
				Key:   "Length",
				Value: length,
			},
		},
	})
	return vr
}

// MaxValue is a function to set the rule that current field value must be no more than the length.
func (vr *ValidationRules) MaxLength(length interface{}) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Max(length).Error("api.msg.error.validation.must_be_no_more_than_length"),
		RuleOpt: []RuleOpt{
			{
				Key:   "Length",
				Value: length,
			},
		},
	})
	return vr
}

// Length is a function to set the rule that current field value must be between min and max.
func (vr *ValidationRules) Length(min int, max int) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    validation.Length(min, max).Error("api.msg.error.validation.must_be_length_between"),
		RuleKey: "length",
		RuleOpt: []RuleOpt{
			{
				Key:   "Min",
				Value: min,
			},
			{
				Key:   "Max",
				Value: max,
			},
		},
	})
	return vr
}

// EqualTo is a function to set the rule that current field value must be equal to target field's value.
func (vr *ValidationRules) EqualTo(targetField string, targetValue string) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.By(EqualValue(targetValue)),
		RuleOpt: []RuleOpt{
			{
				Key:   "Target",
				Value: "attributes." + targetField,
			},
		},
	})
	return vr
}

// IsAlpha is a function to set the rule that current field value must be letters only.
func (vr *ValidationRules) IsAlpha() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    is.Alpha.Error("api.msg.error.validation.must_be_alpha"),
		RuleOpt: nil,
	})
	return vr
}

// IsAlphaSpace is a function to set the rule that current field value must be letters and space character only.
func (vr *ValidationRules) IsAlphaSpace() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Match(regexp.MustCompile("^[a-zA-Z\\s]*$")).
			Error("api.msg.error.validation.must_be_alpha_space"),
		RuleOpt: nil,
	})
	return vr
}

// IsLowerAlphaUnderscore is a function to set the rule that current field value must be lower letters and underscore character only.
func (vr *ValidationRules) IsLowerAlphaUnderscore() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Match(regexp.MustCompile("^[a-z_]*$")).
			Error("api.msg.error.validation.must_be_lower_alpha_underscore"),
		RuleOpt: nil,
	})
	return vr
}

// IsDigit is a function to set the rule that current field value must be digit only.
func (vr *ValidationRules) IsDigit() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Match(regexp.MustCompile("^[0-9]*$")).
			Error("api.msg.error.validation.must_be_digit"),
		RuleOpt: nil,
	})
	return vr
}

// IsAlphaNumeric is a function to set the rule that current field value must be letters and numbers only.
func (vr *ValidationRules) IsAlphaNumeric() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    is.Alphanumeric.Error("api.msg.error.validation.must_be_alphanumeric"),
		RuleOpt: nil,
	})
	return vr
}

// IsAlphaNumericSpace is a function to set the rule that current field value must be letters, numbers, and
// space characters only.
func (vr *ValidationRules) IsAlphaNumericSpace() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Match(regexp.MustCompile("^[a-zA-Z0-9\\s]*$")).
			Error("api.msg.error.validation.must_be_alphanumeric_space"),
		RuleOpt: nil,
	})
	return vr
}

// IsAlphaNumericSpaceAndSpecialCharacter is a function to set the rule that current field value must be letters, numbers,
// space characters, and allowed special characters only.
func (vr *ValidationRules) IsAlphaNumericSpaceAndSpecialCharacter() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Match(regexp.MustCompile("^[a-zA-Z0-9_+\\s]*$")).
			Error("api.msg.error.validation.must_be_alphanumeric_space_special_character"),
		RuleOpt: nil,
	})
	return vr
}

// IsDate is a function to set the rule that current field value is valid date format.
func (vr *ValidationRules) IsDate(layout string) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Date(layout).Error("api.msg.error.validation.must_be_date"),
		RuleOpt: []RuleOpt{
			{
				Key:   "Layout",
				Value: layout,
			},
		},
	})
	return vr
}

// IsTime is a function to set the rule that current field value is valid date format.
func (vr *ValidationRules) IsTime(layout string) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.Date(layout).Error("api.msg.error.validation.must_be_time"),
		RuleOpt: []RuleOpt{
			{
				Key:   "Layout",
				Value: layout,
			},
		},
	})
	return vr
}

// IsUUID is a function to set the rule that current field value must be valid UUID.
func (vr *ValidationRules) IsUUID() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    is.UUID.Error("api.msg.error.validation.must_be_uuid"),
		RuleOpt: nil,
	})
	return vr
}

// IsInt is a function to set the rule that current field value must be int data type.
func (vr *ValidationRules) IsInt() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    is.Int.Error("api.msg.error.validation.must_be_int"),
		RuleOpt: nil,
	})
	return vr
}

// IsFloat is a function to set the rule that current field value must be float data type.
func (vr *ValidationRules) IsFloat() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    is.Float.Error("api.msg.error.validation.must_be_float"),
		RuleOpt: nil,
	})
	return vr
}

// IsEmail is a function to set the rule that current field value must be valid email.
func (vr *ValidationRules) IsEmail() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    is.EmailFormat.Error("api.msg.error.validation.must_be_email"),
		RuleOpt: nil,
	})
	return vr
}

// IsPhone is a function to set the rule that current field value must be valid phone number with country code.
func (vr *ValidationRules) IsPhone() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    is.E164.Error("api.msg.error.validation.must_be_phone"),
		RuleOpt: nil,
	})
	return vr
}

// IsURL is a function to set the rule that current field value must be valid URL.
func (vr *ValidationRules) IsURL() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    is.URL.Error("api.msg.error.validation.must_be_url"),
		RuleOpt: nil,
	})
	return vr
}

// IsJSON is a function to set the rule that current field value must be valid JSON.
func (vr *ValidationRules) IsJSON() *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule:    is.JSON.Error("api.msg.error.validation.must_be_json"),
		RuleOpt: nil,
	})
	return vr
}

// IsExists
func (vr *ValidationRules) IsExists(tx *gorm.DB, tableName string, fieldName string) *ValidationRules {
	vr.Rules = append(vr.Rules, ValidationRule{
		Rule: validation.By(func(fieldValue interface{}) error {
			err := tx.Model(tableName).Where(fieldName+" = ?", fieldValue).Error
			if err != nil {
				return err
			}
			return nil
		}),
		RuleOpt: nil,
	})
	return vr
}

// joinSlice is a function uses to join slice of interfaces.
func joinSlice(sliceOfInterface []interface{}) string {
	var sliceOfString []string
	for _, slice := range sliceOfInterface {
		switch slice.(type) {
		case string:
			sliceOfString = append(sliceOfString, slice.(string))
		case int:
			sliceOfString = append(sliceOfString, strconv.Itoa(slice.(int)))
		}
	}

	return strings.Join(sliceOfString, "/")
}

// EqualValue is a closure uses by EqualTo method.
func EqualValue(targetValue string) validation.RuleFunc {
	return func(fieldValue interface{}) error {
		s, _ := fieldValue.(string)
		if s != targetValue {
			return errors.New("api.msg.error.validation.must_be_equal_to")
		}
		return nil
	}
}
