package validator

import (
	"go-rest-skeleton/pkg/response"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Validator struct {
	Rules []RuleGroup
}

type RuleGroup struct {
	Field string
	Data  interface{}
	Rules []Rule
}

type Rule struct {
	Rule    validation.Rule
	RuleOpt []RuleOpt
}

type RuleOpt struct {
	key   string
	value interface{}
}

func New() *Validator {
	return &Validator{}
}

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
			errMsgData = map[string]interface{}{ruleGroup.Field: ruleGroup.Data}

			for _, opt := range rulesOpt {
				errMsgData[opt.key] = opt.value
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
