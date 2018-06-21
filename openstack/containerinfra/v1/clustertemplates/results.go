package clustertemplates

import (
	"encoding/json"
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

// Extract is a function that accepts a result and extracts a cluster-template resource.
func (r commonResult) Extract() (*ClusterTemplate, error) {
	var s *ClusterTemplate
	err := r.ExtractInto(&s)
	return s, err
}

// Represents a template for a Cluster Template
type ClusterTemplate struct {
	InsecureRegistry    string             `json:"insecure_registry"`
	Links               []gophercloud.Link `json:"links"`
	HTTPProxy           string             `json:"http_proxy"`
	UpdatedAt           time.Time          `json:"-"`
	FloatingIPEnabled   bool               `json:"floating_ip_enabled"`
	FixedSubnet         string             `json:"fixed_subnet"`
	MasterFlavorID      string             `json:"master_flavor_id"`
	UUID                string             `json:"uuid"`
	NoProxy             string             `json:"no_proxy"`
	HTTPSProxy          string             `json:"https_proxy"`
	TLSDisabled         bool               `json:"tls_disabled"`
	KeyPairID           string             `json:"keypair_id"`
	Public              bool               `json:"public"`
	Labels              map[string]string  `json:"labels"`
	DockerVolumeSize    int                `json:"docker_volume_size"`
	ServerType          string             `json:"server_type"`
	ExternalNetworkID   string             `json:"external_network_id"`
	ClusterDistro       string             `json:"cluster_distro"`
	ImageID             string             `json:"image_id"`
	VolumeDriver        string             `json:"volume_driver"`
	RegistryEnabled     bool               `json:"registry_enabled"`
	DockerStorageDriver string             `json:"docker_storage_driver"`
	APIServerPort       string             `json:"apiserver_port"`
	Name                string             `json:"name"`
	CreatedAt           time.Time          `json:"-"`
	NetworkDriver       string             `json:"network_driver"`
	FixedNetwork        string             `json:"fixed_network"`
	COE                 string             `json:"coe"`
	FlavorID            string             `json:"flavor_id"`
	MasterLBEnabled     bool               `json:"master_lb_enabled"`
	DNSNameServer       string             `json:"dns_nameserver"`
}

func (r *ClusterTemplate) UnmarshalJSON(b []byte) error {
	type tmp ClusterTemplate
	var s struct {
		tmp
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = ClusterTemplate(s.tmp)

	if s.CreatedAt != "" {
		r.CreatedAt, err = time.Parse(time.RFC3339, s.CreatedAt)
		if err != nil {
			return err
		}
	}

	if s.UpdatedAt != "" {
		r.UpdatedAt, err = time.Parse(time.RFC3339, s.UpdatedAt)
		if err != nil {
			return err
		}
	}
	return nil
}

// ClusterTemplatePage is the page returned by a pager when traversing over a
// collection of cluster-templates.
type ClusterTemplatePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of cluster template has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r ClusterTemplatePage) NextPageURL() (string, error) {
	var s struct {
		Next string `json:"next"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Next, nil
}

// IsEmpty checks whether a ClusterTemplatePage struct is empty.
func (r ClusterTemplatePage) IsEmpty() (bool, error) {
	is, err := ExtractClusterTemplates(r)
	return len(is) == 0, err
}

// ExtractClusterTemplates accepts a Page struct, specifically a ClusterTemplatePage struct,
// and extracts the elements into a slice of cluster templates structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractClusterTemplates(r pagination.Page) ([]ClusterTemplate, error) {
	var s struct {
		ClusterTemplates []ClusterTemplate `json:"clustertemplates"`
	}
	err := (r.(ClusterTemplatePage)).ExtractInto(&s)
	return s.ClusterTemplates, err
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

func (r DeleteResult) Extract() (string, error) {
	uuid := ""
	idKey := "X-Openstack-Request-Id"
	if len(r.Header[idKey]) > 0 {
		uuid = r.Header[idKey][0]
		if uuid == "" {
			return "", fmt.Errorf("No value for header %s", idKey)
		}
	} else {
		return "", fmt.Errorf("Missing [%s] for header", idKey)
	}
	return uuid, r.ExtractErr()
}

// UpdateResult is the response of a Update operations.
type UpdateResult struct {
	commonResult
}
