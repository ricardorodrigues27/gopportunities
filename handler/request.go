package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type:  %s) is required", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
	Currency string `json:"currency"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "boolean")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "number")
	}
	if r.Currency == "" {
		return errParamIsRequired("currency", "string")
	}
	return nil
}
