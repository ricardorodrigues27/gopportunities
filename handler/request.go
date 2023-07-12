package handler

import (
	"fmt"
)

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

// UpdateOpening

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
	Currency string `json:"currency"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 || r.Currency != "" {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}

func (r UpdateOpeningRequest) Map() map[string]interface{} {
	out := make(map[string]interface{})

	if r.Role != "" {
		out["Role"] = r.Role
	}
	if r.Company != "" {
		out["Company"] = r.Company
	}
	if r.Location != "" {
		out["Location"] = r.Location
	}
	if r.Link != "" {
		out["Link"] = r.Link
	}
	if r.Remote != nil {
		out["Remote"] = *r.Remote
	}
	if r.Salary > 0 {
		out["Salary"] = r.Salary
	}
	if r.Currency != "" {
		out["Currency"] = r.Currency
	}

	return out
}
