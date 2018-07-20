package quotas

import (
	"net/http"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToQuotaCreateMap() (map[string]interface{}, error)
}

// CreateOpts params
type CreateOpts struct {
	ProjectID string `json:"project_id"`
	Resource  string `json:"resource"`
	HardLimit int    `json:"hard_limit"`
}

// ToQuotaCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToQuotaCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "")
}

// Create requests the creation of a new quota.
func Create(client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToQuotaCreateMap()
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

// Get retrieves a specific quota
func Get(client *gophercloud.ServiceClient, id string, resource string) (r GetResult) {
	var result *http.Response
	result, r.Err = client.Get(getURL(client, id, resource), &r.Body, &gophercloud.RequestOpts{OkCodes: []int{200}})
	if r.Err == nil {
		r.Header = result.Header
	}
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToQuotasListQuery() (string, error)
}

type ListOpts struct {
	Marker  string `q:"marker"`
	Limit   int    `q:"limit"`
	SortKey string `q:"sort_key"`
	SortDir string `q:"sort_dir"`
}

// ToQuotasListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToQuotasListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// quotas. It accepts a ListOptsBuilder, which allows you to sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToQuotasListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return QuotasPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type UpdateOpts struct {
	ProjectID string `json:"project_id"`
	Resource  string `json:"resource"`
	HardLimit int    `json:"hard_limit"`
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToQuotasUpdateMap() (map[string]interface{}, error)
}

// ToQuotasUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateOpts) ToQuotasUpdateMap() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Update implements quota updated request.
func Update(client *gophercloud.ServiceClient, id string, resource string, opt UpdateOptsBuilder) (r UpdateResult) {
	b, err := opt.ToQuotasUpdateMap()
	if err != nil {
		r.Err = err
		return r
	}
	var result *http.Response
	result, r.Err = client.Patch(updateURL(client, id, resource), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 202},
	})

	if r.Err == nil {
		r.Header = result.Header
	}
	return
}

// Delete deletes the specified quota.
func Delete(client *gophercloud.ServiceClient, id string, resource string) (r DeleteResult) {
	var result *http.Response
	result, r.Err = client.Delete(deleteURL(client, id, resource), nil)
	r.Header = result.Header
	return
}
