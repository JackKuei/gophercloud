package clusters

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

func (r CreateResult) Extract() (clusterID string, err error) {
	var s *UUID
	err = r.ExtractInto(&s)
	if err == nil {
		clusterID = s.UUID
	}
	return clusterID, err

}

func (r UpdateResult) Extract() (clusterID string, err error) {
	var s *UUID
	err = r.ExtractInto(&s)
	if err == nil {
		clusterID = s.UUID
	}
	return clusterID, err

}

// Extract is a function that accepts a result and extracts a cluster resource.
func (r commonResult) Extract() (*Cluster, error) {
	var s *Cluster
	err := r.ExtractInto(&s)
	return s, err
}

type Cluster struct {
	Status            string             `json:"status"`
	UUID              string             `json:"uuid"`
	Links             []gophercloud.Link `json:"links"`
	StackID           string             `json:"stack_id"`
	CreatedAt         time.Time          `json:"created_at"`
	APIAddress        string             `json:"api_address"`
	DiscoveryURL      string             `json:"discovery_url"`
	UpdatedAt         time.Time          `json:"updated_at"`
	MasterCount       int                `json:"master_count"`
	COEVersion        string             `json:"coe_version"`
	KeyPair           string             `json:"keypair"`
	ClusterTemplateID string             `json:"cluster_template_id"`
	MasterAddresses   []string           `json:"master_addresses"`
	NodeCount         int                `json:"node_count"`
	NodeAddresses     []string           `json:"node_addresses"`
	StatusReason      string             `json:"status_reason"`
	CreateTimeout     int                `json:"create_timeout"`
	Name              string             `json:"name"`
}

// ClusterPage is the page returned by a pager when traversing over a
// collection of clusters.
type ClusterPage struct {
	pagination.LinkedPageBase
}

func (r ClusterPage) NextPageURL() (string, error) {
	var s struct {
		Next string `json:"next"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Next, nil
}

// IsEmpty checks whether a ClusterPage struct is empty.
func (r ClusterPage) IsEmpty() (bool, error) {
	is, err := ExtractClusters(r)
	return len(is) == 0, err
}

func ExtractClusters(r pagination.Page) ([]Cluster, error) {
	var s struct {
		Clusters []Cluster `json:"clusters"`
	}
	err := (r.(ClusterPage)).ExtractInto(&s)
	return s.Clusters, err
}

// DeleteResult is the result from a Delete operation. Call its Extract or ExtractErr
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
