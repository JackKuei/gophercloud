package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/quotas"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreateQuota(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleCreateQuotaSuccessfully(t)

	opts := quotas.CreateOpts{
		ProjectID: "aa5436ab58144c768ca4e9d2e9f5c3b2",
		Resource:  "Cluster",
		HardLimit: 10,
	}

	sc := fake.ServiceClient()
	sc.Endpoint = sc.Endpoint + "v1/"

	res := quotas.Create(sc, opts)
	th.AssertNoErr(t, res.Err)

	requestID := res.Header.Get("X-OpenStack-Request-Id")
	th.AssertEquals(t, requestUUID, requestID)

	quota, err := res.Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, projectID, quota.ProjectID)
}

func TestDeleteQuota(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleDeleteQuotaSuccessfully(t)

	sc := fake.ServiceClient()
	sc.Endpoint = sc.Endpoint + "v1/"

	err := quotas.Delete(sc, projectID, resourceType).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestListQuotas(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleListQuotasSuccessfully(t)

	count := 0

	sc := fake.ServiceClient()
	sc.Endpoint = sc.Endpoint + "v1/"

	quotas.List(sc, quotas.ListOpts{Limit: 2}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		quotas, err := quotas.ExtractQuotas(page)
		th.AssertNoErr(t, err)
		for idx := range quotas {
			quotas[idx].CreatedAt = quotas[idx].CreatedAt.UTC()
			quotas[idx].UpdatedAt = quotas[idx].UpdatedAt.UTC()
		}
		th.AssertDeepEquals(t, ExpectedQuotas, quotas)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestUpdateQuota(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleUpdateQuotaSuccessfully(t)

	updateOpts := quotas.UpdateOpts{
		ProjectID: "aa5436ab58144c768ca4e9d2e9f5c3b2",
		Resource:  "Cluster",
		HardLimit: 20,
	}

	sc := fake.ServiceClient()
	sc.Endpoint = sc.Endpoint + "v1/"

	res := quotas.Update(sc, projectID, resourceType, updateOpts)
	th.AssertNoErr(t, res.Err)

	requestID := res.Header.Get("X-OpenStack-Request-Id")
	th.AssertEquals(t, requestUUID, requestID)

	quota, err := res.Extract()
	th.AssertNoErr(t, err)

	quota.CreatedAt = quota.CreatedAt.UTC()
	quota.UpdatedAt = quota.UpdatedAt.UTC()

	th.AssertDeepEquals(t, &ExpectedQuota, quota)
}
