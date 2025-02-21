package bootstrap

import (
	"cms-server/constants"
	"regexp"

	"github.com/asaskevich/govalidator"
)

type Validator interface {
	customEmailOrPhoneValidator(str string, params ...string) bool
}

type validator struct{}

func RegisterValidator() {
	v := &validator{}
	
	govalidator.ParamTagRegexMap["email_phone"] = regexp.MustCompile(`^email_phone\((\w+)\)$`)
	govalidator.ParamTagMap["email_phone"] = govalidator.ParamValidator(v.customEmailOrPhoneValidator)
}

func (v *validator) customEmailOrPhoneValidator(str string, params ...string) bool {
	if len(params) == 0 {
		return false
	}

	if govalidator.IsEmail(str) {
		return true
	}
	return v.isPhoneNumber(str, params[0])
}

func (v *validator) isPhoneNumber(str, local string) bool {
	switch local {
	case constants.VI:
		regexp := regexp.MustCompile(`/(03|05|07|08|09|01[2|6|8|9])+([0-9]{8})\b/`)
		return regexp.MatchString(str)
	default:
		phoneRegex := regexp.MustCompile(`^\+?[0-9]{9,15}$`)
		return phoneRegex.MatchString(str)
	}
}
