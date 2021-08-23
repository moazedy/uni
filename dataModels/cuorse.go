package dataModels

import "errors"

// Cuorse is a struct to demonstration of cuorse concept in system
type Cuorse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

// Validate validates target cuorse data
func (c *Cuorse) Validate() error {
	err := c.NameValidation()
	if err != nil {
		return err
	}

	err = c.GradeValidation()
	if err != nil {
		return err
	}

	return nil
}

// NameValidation validates target cuorse name
func (c *Cuorse) NameValidation() error {
	if c.Name == "" {
		return errors.New("cuorese name can not be empty!")
	}

	if len(c.Name) > 15 {
		return errors.New("the name is too long, it should be less than 15 characters")
	}

	if len(c.Name) < 4 {
		return errors.New("the name is too short, it should be more than 4 characters")
	}

	return nil
}

// GradeValidation validates target cuorse's grade
func (c *Cuorse) GradeValidation() error {
	if c.Grade < 0 || c.Grade > 20 {
		return errors.New("grade is out of range, it should be between 0 to 20")
	}

	return nil
}

type CreateCuorseRequest struct {
	Name string `json:"name"`
}
