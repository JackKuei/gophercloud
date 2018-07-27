package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clusters"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreateCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleCreateClusterSuccessfully(t)

	opts := clusters.CreateOpts{
		Name:              "k8s",
		DiscoveryURL:      "",
		MasterCount:       1,
		ClusterTemplateID: "0562d357-8641-4759-8fed-8173f02c9633",
		NodeCount:         2,
		CreateTimeout:     60,
		KeyPair:           "my_keypair",
		MasterFlavorID:    "",
		Labels:            map[string]string{},
		FlavorID:          "",
	}

	sc := fake.ServiceClient()
	sc.Endpoint = sc.Endpoint + "v1/"
	res := clusters.Create(sc, opts)
	th.AssertNoErr(t, res.Err)

	requestID := res.Header.Get("X-OpenStack-Request-Id")
	th.AssertEquals(t, requestUUID, requestID)

	actual, err := res.Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, clusterUUID, actual)
}

func TestDeleteCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleDeleteClusterSuccessfully(t)

	sc := fake.ServiceClient()
	sc.Endpoint = sc.Endpoint + "v1/"
	uuid, err := clusters.Delete(sc, clusterUUID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, requestUUID, uuid)
}

func TestListClusters(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleListClusterSuccessfully(t)

	count := 0
	sc := fake.ServiceClient()
	sc.Endpoint = sc.Endpoint + "v1/"
	clusters.List(sc, clusters.ListOpts{Limit: 2}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := clusters.ExtractClusters(page)
		th.AssertNoErr(t, err)
		for idx := range actual {
			actual[idx].CreatedAt = actual[idx].CreatedAt.UTC()
			actual[idx].UpdatedAt = actual[idx].UpdatedAt.UTC()
		}
		th.AssertDeepEquals(t, ExpectedClusters, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestUpdateCluster(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleUpdateClusterSuccessfully(t)

	updateOpts := []clusters.UpdateOptsBuilder{
		clusters.UpdateOpts{
			Path:  "/master_lb_enabled",
			Value: "True",
			Op:    "replace",
		},
		clusters.UpdateOpts{
			Path:  "/registry_enabled",
			Value: "True",
			Op:    "replace",
		},
	}

	sc := fake.ServiceClient()
	sc.Endpoint = sc.Endpoint + "v1/"
	res := clusters.Update(sc, clusterUUID, updateOpts)
	th.AssertNoErr(t, res.Err)

	requestID := res.Header.Get("X-OpenStack-Request-Id")
	th.AssertEquals(t, requestUUID, requestID)

	actual, err := res.Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, clusterUUID, actual)
}
