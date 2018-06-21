package clustertemplates

import (
	"fmt"
	"net/http"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToClusterCreateMap() (map[string]interface{}, error)
}

// CreateOpts params
type CreateOpts struct {
	Labels              map[string]string `json:"labels,omitempty"`
	FixedSubnet         string            `json:"fixed_subnet,omitempty"`
	MasterFlavorID      string            `json:"master_flavor_id,omitempty"`
	NoProxy             string            `json:"no_proxy,omitempty"`
	HTTPSProxy          string            `json:"https_proxy,omitempty"`
	HTTPProxy           string            `json:"http_proxy,omitempty"`
	TLSDisabled         bool              `json:"tls_disabled"`
	KeyPairID           string            `json:"keypair_id"`
	Public              bool              `json:"public"`
	DockerVolumeSize    int               `json:"docker_volume_size"`
	ServerType          string            `json:"server_type,omitempty"`
	ExternalNetworkID   string            `json:"external_network_id"`
	ImageID             string            `json:"image_id"`
	VolumeDriver        string            `json:"volume_driver"`
	RegistryEnabled     bool              `json:"registry_enabled"`
	DockerStorageDriver string            `json:"docker_storage_driver"`
	Name                string            `json:"name" required:"true"`
	NetworkDriver       string            `json:"network_driver"`
	FixedNetwork        string            `json:"fixed_network,omitempty"`
	COE                 string            `json:"coe" required:"true"`
	FlavorID            string            `json:"flavor_id"`
	MasterLBEnabled     bool              `json:"master_lb_enabled"`
	DNSNameServer       string            `json:"dns_nameserver"`
}

// ToClusterCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToClusterCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create requests the creation of a new cluster.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToClusterCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	var result *http.Response
	result, r.Err = client.Post(createURL(client), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 201, 202},
	})

	if r.Err == nil {
		r.Header = result.Header
	}

	return
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToClusterTemplateListQuery() (string, error)
}

// ListOpts allows the sorting of paginated collections through
// the API. SortKey allows you to sort by a particular cluster templates attribute.
// SortDir sets the direction, and is either `asc' or `desc'.
// Marker and Limit are used for pagination.
type ListOpts struct {
	Marker  string `q:"marker"`
	Limit   int    `q:"limit"`
	SortKey string `q:"sort_key"`
	SortDir string `q:"sort_dir"`
}

// ToClusterTemplateListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToClusterTemplateListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// cluster-templates. It accepts a ListOptsBuilder, which allows you to sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToClusterTemplateListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ClusterTemplatePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type UpdateOpts struct {
	Path  string `json:"path" required:"true"`
	Value string `json:"value,omitempty"`
	Op    string `json:"op" required:"true"`
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToClusterTemplateUpdateMap() (map[string]interface{}, error)
}

// ToClusterUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateOpts) ToClusterTemplateUpdateMap() (map[string]interface{}, error) {
	if opts.Op != "remove" && opts.Value == "" {
		return nil, fmt.Errorf("Value field must be provied for Op=[%s]. Only when Op=remove can have empty Value field", opts.Op)
	}
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Update implements cluster updated request.
func Update(client *gophercloud.ServiceClient, id string, opts []UpdateOptsBuilder) (r UpdateResult) {
	var o []map[string]interface{}
	for _, opt := range opts {
		b, err := opt.ToClusterTemplateUpdateMap()
		if err != nil {
			r.Err = err
			return r
		}
		o = append(o, b)
	}
	var result *http.Response
	result, r.Err = client.Patch(updateURL(client, id), o, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})

	if r.Err == nil {
		r.Header = result.Header
	}
	return
}

// Delete deletes the specified cluster ID.
func Delete(client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	var result *http.Response
	result, r.Err = client.Delete(deleteURL(client, id), nil)
	r.Header = result.Header
	return
}
