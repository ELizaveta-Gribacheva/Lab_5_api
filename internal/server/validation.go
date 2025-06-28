package server

import "errors"

func ValidateTaskInput(title, description string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	if description == "" {
		return errors.New("description cannot be empty")
	}
	return nil
}
