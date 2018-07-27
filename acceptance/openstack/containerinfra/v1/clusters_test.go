// +build acceptance containerinfra

package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clusters"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestClustersCRUD(t *testing.T) {
	client, err := clients.NewContainerV1Client()
	th.AssertNoErr(t, err)

	clusterTemplate, err := CreateClusterTemplate(t, client)
	th.AssertNoErr(t, err)
	defer DeleteClusterTemplate(t, client, clusterTemplate.UUID)

	clusterID, err := CreateCluster(t, client, clusterTemplate.UUID)
	defer DeleteCluster(t, client, clusterID)
	th.AssertNoErr(t, err)

	allPages, err := clusters.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allClusters, err := clusters.ExtractClusters(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, v := range allClusters {
		if v.UUID == clusterID {
			found = true
		}
	}
	th.AssertEquals(t, found, true)

	updateOpts := []clusters.UpdateOptsBuilder{
		clusters.UpdateOpts{
			Path:  "/node_count",
			Value: "2",
			Op:    "replace",
		},
	}
	updateResult := clusters.Update(client, clusterID, updateOpts)
	th.AssertNoErr(t, updateResult.Err)

	if len(updateResult.Header["X-Openstack-Request-Id"]) > 0 {
		t.Logf("Cluster Update Request ID: %s", updateResult.Header["X-Openstack-Request-Id"][0])
	}

	clusterID, err = updateResult.Extract()
	th.AssertNoErr(t, err)

	err = WaitForCluster(client, clusterID, "SUCCESS")
	th.AssertNoErr(t, err)

	newCluster, err := clusters.Get(client, clusterID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, newCluster.UUID, clusterID)

	tools.PrintResource(t, newCluster)
}
