package validation

import (
	"Skillture_Form/internal/domain/entities"
	"Skillture_Form/internal/domain/enums"
	"errors"
)

// ValidateResponseDomain checks the domain rules for Response entity
func ValidateResponseDomain(response *entities.Response) error {
	return response.IsValid()
}

// ValidateResponseBusiness checks business rules for Response
func ValidateResponseBusiness(response *entities.Response, form *entities.Form) error {
	if form == nil {
		return errors.New("form not found")
	}
	if form.Status != enums.FormStatusPublished {
		return errors.New("form is not accepting responses")
	}
	return nil
}
