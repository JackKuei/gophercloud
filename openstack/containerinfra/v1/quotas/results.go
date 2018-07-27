package quotas

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// CreateResult is the response of a Create operations.
type CreateResult struct {
	commonResult
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

// UpdateResult is the response of a Update operations.
type UpdateResult struct {
	commonResult
}

type UUID struct {
	UUID string `json:"uuid"`
}

// Extract is a function that accepts a result and extracts a quota resource.
func (r commonResult) Extract() (*Quotas, error) {
	var s *Quotas
	err := r.ExtractInto(&s)
	return s, err
}

type Quotas struct {
	Resource  string    `json:"resource"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	HardLimit int       `json:"hard_limit"`
	ProjectID string    `json:"project_id"`
	ID        string    `json:"id"`
}

func (r *Quotas) UnmarshalJSON(b []byte) error {
	type tmp Quotas
	var s struct {
		tmp
		ID interface{} `json:"id"`
	}

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Quotas(s.tmp)

	switch t := s.ID.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		r.ID = fmt.Sprint(t)
	case string:
		r.ID = t
	}

	return nil
}

// QuotasPage is the page returned by a pager when traversing over a
// collection of quotas.
type QuotasPage struct {
	pagination.LinkedPageBase
}

func (r QuotasPage) NextPageURL() (string, error) {
	var s struct {
		Next string `json:"next"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Next, nil
}

// IsEmpty checks whether a QuotasPage struct is empty.
func (r QuotasPage) IsEmpty() (bool, error) {
	is, err := ExtractQuotas(r)
	return len(is) == 0, err
}

func ExtractQuotas(r pagination.Page) ([]Quotas, error) {
	var s struct {
		Quotas []Quotas `json:"quotas"`
	}
	err := (r.(QuotasPage)).ExtractInto(&s)
	return s.Quotas, err
}
