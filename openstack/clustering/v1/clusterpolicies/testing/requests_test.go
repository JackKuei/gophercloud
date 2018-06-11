package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/clustering/v1/clusterpolicies"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestGetAction(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleGetSuccessfully(t, ExpectedGetClusterPolicy.ID)

	clusterID := "7d85f602-a948-4a30-afd4-e84f47471c15"
	policyID := "714fe676-a08f-4196-b7af-61d52eeded15"
	actual, err := clusterpolicies.Get(fake.ServiceClient(), clusterID, policyID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedGetClusterPolicy, *actual)
}
