package valid

import (
	"github.com/go-playground/validator/v10"
	"github.com/tdatIT/who-sent-api/pkgs/logger"

	"regexp"
)

type Validator struct {
	Valid *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	err := v.Valid.Struct(i)
	if err != nil {
		return err
	}
	return nil
}

func InitValidatorInstance() *validator.Validate {
	validate := validator.New()

	err := validate.RegisterValidation("lowerCaseNoSpace", validateLowerCaseNoSpace)
	if err != nil {
		return nil
	}

	return validate
}

func validateLowerCaseNoSpace(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	var lowercaseNoSpaceRegex = regexp.MustCompile(`^[a-z0-9_]+$`)
	if lowercaseNoSpaceRegex.MatchString(value) {
		return true
	}
	logger.Debug("invalid input: only lowercase letters, numbers, and underscores allowed without spaces or special characters")
	return false
}
