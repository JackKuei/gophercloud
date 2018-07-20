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
		ProjectID:		   "aa5436ab58144c768ca4e9d2e9f5c3b2",
		Resource:          "Cluster",
		HardLimit:         10,
	}

	res := quotas.Create(fake.ServiceClient(), opts)
	th.AssertNoErr(t, res.Err)

	requestID := res.Header.Get("X-OpenStack-Request-Id")
	th.AssertEquals(t, requestUUID, requestID)

	actual, err := res.Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, projectID, actual)
}

func TestDeleteQuota(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleDeleteQuotaSuccessfully(t)

	uuid, err := quotas.Delete(fake.ServiceClient(), projectID, resourceID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, requestUUID, uuid)
}

func TestListQuotas(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleListQuotasSuccessfully(t)

	count := 0

	quotas.List(fake.ServiceClient(), quotas.ListOpts{Limit: 2}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := quotas.ExtractQuotas(page)
		th.AssertNoErr(t, err)
		for idx := range actual {
			actual[idx].CreatedAt = actual[idx].CreatedAt.UTC()
			actual[idx].UpdatedAt = actual[idx].UpdatedAt.UTC()
		}
		th.AssertDeepEquals(t, ExpectedQuotas, actual)

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
		ProjectID:		   "aa5436ab58144c768ca4e9d2e9f5c3b2",
		Resource:          "Cluster",
		HardLimit:         20,
	}

	res := quotas.Update(fake.ServiceClient(), projectID, resourceType, updateOpts)
	th.AssertNoErr(t, res.Err)

	requestID := res.Header.Get("X-OpenStack-Request-Id")
	th.AssertEquals(t, requestUUID, requestID)

	actual, err := res.Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, projectID, actual)
}
