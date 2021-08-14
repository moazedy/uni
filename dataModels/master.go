package dataModels

import "errors"

type Master struct {
	Id   string
	Name string
}

func (m *Master) Validate() error {
	err := m.NameValidation()
	if err != nil {
		return err
	}

	return nil
}

// NameValidation validates target master name
func (m *Master) NameValidation() error {
	if m.Name == "" {
		return errors.New("master name can not be empty!")
	}

	if len(m.Name) > 15 {
		return errors.New("the name is too long, it should be less than 15 characters")
	}

	if len(m.Name) < 4 {
		return errors.New("the name is too short, it should be more than 4 characters")
	}

	return nil
}
