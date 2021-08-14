package dataModels

import "errors"

type Student struct {
	Id           string
	Name         string
	Cuorses      []Cuorse
	AverageGrade int
}

// Validate validates target student data
func (s *Student) Validate() error {
	err := s.NameValidation()
	if err != nil {
		return err
	}

	err = s.AverageGradeValidation()
	if err != nil {
		return err
	}

	return nil
}

// NameValidation validates target cuorse name
func (s *Student) NameValidation() error {
	if s.Name == "" {
		return errors.New("student name can not be empty!")
	}

	if len(s.Name) > 15 {
		return errors.New("the name is too long, it should be less than 15 characters")
	}

	if len(s.Name) < 4 {
		return errors.New("the name is too short, it should be more than 4 characters")
	}

	return nil
}

// GradeValidation validates target cuorse's grade
func (s *Student) AverageGradeValidation() error {
	if s.AverageGrade < 0 || s.AverageGrade > 20 {
		return errors.New("average grade is out of range, it should be between 0 to 20")
	}

	return nil
}
