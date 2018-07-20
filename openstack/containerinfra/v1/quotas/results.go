package quotas

import (
	"fmt"
	"time"

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

func (r CreateResult) Extract() (quotaID string, err error) {
	var s *UUID
	err = r.ExtractInto(&s)
	if err == nil {
		quotaID = s.UUID
	}
	return quotaID, err

}

func (r UpdateResult) Extract() (quotaID string, err error) {
	var s *UUID
	err = r.ExtractInto(&s)
	if err == nil {
		quotaID = s.UUID
	}
	return quotaID, err

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

// DeleteResult is the result from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

func (r DeleteResult) Extract() (string, error) {
	requestID := ""
	idKey := "X-Openstack-Request-Id"
	if len(r.Header[idKey]) > 0 {
		requestID = r.Header[idKey][0]
		if requestID == "" {
			return "", fmt.Errorf("No value for header %s", idKey)
		}
	} else {
		return "", fmt.Errorf("Missing %s for header", idKey)
	}
	return requestID, r.ExtractErr()
}
