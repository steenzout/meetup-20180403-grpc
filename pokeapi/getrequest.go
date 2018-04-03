package pokeapi

import (
	"fmt"
)

// ValidateGetRequest returns wether GetRequest is valid.
// If valid, ValidateGetRequest returns nil.
// Otherwise, it returns an error that describes the problem.
func ValidateGetRequest(req *GetRequest) error {
	if req.Id < 1 {
		return fmt.Errorf("ID: %d is invalid", req.Id)
	}

	return nil
}

// NewGetRequest creates a GetRequest.
func NewGetRequest(id uint64) (*GetRequest, error) {
	req := &GetRequest{
		Id: id,
	}

	return req, ValidateGetRequest(req)
}
