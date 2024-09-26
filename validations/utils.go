package validations

import "errors"

// Validate name field
func ValidateName(value string) error {
	if value == "" {
		return errors.New("the name field is required")
	}
	return nil
}

// Validate email field
func ValidateEmail(value string) error {
	if value == "" {
		return errors.New("the email field is required")
	}
	return nil
}

// Validate password field
func ValidatePassword(value string) error {
	if value == "" {
		return errors.New("the password field is required")
	}
	return nil
}

// Validate role field
func ValidateRole(value string) error {
	if value == "" {
		return errors.New("the role field is required")
	}
	return nil
}

// Validate status field
func ValidateStatus(value uint8) error {
	// Assuming status should be either 0 or 1 (or similar values)
	if value != 0 && value != 1 {
		return errors.New("the status field must be 0 or 1")
	}
	return nil
}

// Validate title field
func ValidateTitle(value string) error {
	if value == "" {
		return errors.New("the title field is required")
	}
	return nil
}

// Validate category field
func ValidateCategory(value string) error {
	if value == "" {
		return errors.New("the category field is required")
	}
	return nil
}

// Validate Content field
func ValidateContent(value string) error {
	if value == "" {
		return errors.New("the content field is required")
	}
	return nil
}

// Validate Date field
func ValidateDate(value string) error {
	if value == "" {
		return errors.New("the date field is required")
	}
	return nil
}