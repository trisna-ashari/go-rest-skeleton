package validator

import (
	"go-rest-skeleton/pkg/response"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validator is a struct represent group of rules.
type Validator struct {
	Rules []RuleGroup
}

// RuleGroup is a struct represent definition of rule.
type RuleGroup struct {
	Field string
	Data  interface{}
	Rules []Rule
}

// Rule is a struct represent a rule and rule option.
type Rule struct {
	Rule    validation.Rule
	RuleOpt []RuleOpt
}

// RuleOpt is a struct represent it self.
type RuleOpt struct {
	Key   string
	Value interface{}
}

// New is a constructor will initialize Validator.
func New() *Validator {
	return &Validator{}
}

// Set is a function uses to set a Rule to the field.
func (v *Validator) Set(field string, data interface{}, rules []ValidationRule) *Validator {
	var fieldRule []Rule
	for _, rule := range rules {
		fieldRule = append(fieldRule, Rule{
			Rule:    rule.Rule,
			RuleOpt: rule.RuleOpt,
		})
	}

	rule := RuleGroup{
		Field: field,
		Data:  data,
		Rules: fieldRule,
	}
	v.Rules = append(v.Rules, rule)

	return v
}

// Validate is a function uses to validates the value of the field.
func (v *Validator) Validate() []response.ErrorForm {
	var errMsg []response.ErrorForm

	for _, ruleGroup := range v.Rules {
		var rules []validation.Rule
		var rulesOpt []RuleOpt

		for _, rule := range ruleGroup.Rules {
			rules = append(rules, rule.Rule)
			for _, ruleOpt := range rule.RuleOpt {
				rulesOpt = append(rulesOpt, ruleOpt)
			}
		}
		err := validation.Validate(ruleGroup.Data,
			rules...,
		)

		if err != nil {
			errMsgData := make(map[string]interface{})
			errMsgData[ruleGroup.Field] = ruleGroup.Data

			for _, opt := range rulesOpt {
				errMsgData[opt.Key] = opt.Value
			}

			errMsg = append(errMsg, response.ErrorForm{
				Field: ruleGroup.Field,
				Msg:   err.Error(),
				Data:  errMsgData,
			})

		}
	}

	return errMsg
}
