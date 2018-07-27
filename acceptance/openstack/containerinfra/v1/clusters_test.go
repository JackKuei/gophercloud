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

// TODO: Delete below when above works
/*
func TestClusters(t *testing.T) {
	testName = tools.RandomString("TESTACC-", 8)
	clusterCreate(t)
	defer clusterDelete(t)
	clusterGet(t)
	clusterList(t)
	clusterUpdate(t)
}

func clusterCreate(t *testing.T) {
	clusterName := testName
	clusterTemplateID := testName

	// TODO: Hard coded for testing
	clusterTemplateID = "TESTACC-HHlNrAEl"

	client, err := clients.NewContainerInfraV1Client()
	if err != nil {
		t.Fatalf("Unable to create clustering client: %v", err)
	}
	client.Microversion = "1.5"

	createOpts := clusters.CreateOpts{
		Name:              clusterName,
		DiscoveryURL:      "",
		MasterCount:       1,
		ClusterTemplateID: clusterTemplateID,
		NodeCount:         1,
		CreateTimeout:     100,
		KeyPair:           "testKeyPair",
		MasterFlavorID:    "",
		Labels:            map[string]string{},
		FlavorID:          "m1.small",
	}

	createResult := clusters.Create(client, createOpts)
	th.AssertNoErr(t, createResult.Err)
	if len(createResult.Header["X-Openstack-Request-Id"]) > 0 {
		t.Logf("Cluster Create Request ID: %s", createResult.Header["X-Openstack-Request-Id"][0])
	}
	clusterID, err := createResult.Extract()
	if err != nil {
		t.Fatalf("Error creating cluster %s: %v", clusterName, err)
	} else {
		t.Logf("Cluster created: %+v", clusterID)
	}
	t.Logf("Cluster Create Completed: clusterID: %s", clusterID)
}

func clusterGet(t *testing.T) {
	client, err := clients.NewContainerInfraV1Client()
	if err != nil {
		t.Fatalf("Unable to create container-infra client: %v", err)
	}
	client.Microversion = "1.5"
	clusterName := testName

	cluster, err := clusters.Get(client, clusterName).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, clusterName, cluster.Name)

	tools.PrintResource(t, clusterName)
	t.Logf("Cluster Get Completed: %s", clusterName)
}

func clusterList(t *testing.T) {
	client, err := clients.NewContainerInfraV1Client()
	if err != nil {
		t.Fatalf("Unable to create container-infra client: %v", err)
	}
	client.Microversion = "1.5"

	testClusterFound := false
	clusters.List(client, clusters.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		allClusters, err := clusters.ExtractClusters(page)
		if err != nil {
			t.Fatalf("Error extracting page of cluster templates: %v", err)
		}

		for _, cluster := range allClusters {
			tools.PrintResource(t, cluster)
			if cluster.Name == testName {
				testClusterFound = true
			}
		}

		empty, err := page.IsEmpty()
		th.AssertNoErr(t, err)

		// Expect the page IS NOT empty
		th.AssertEquals(t, false, empty)

		return true, nil
	})

	th.AssertEquals(t, true, testClusterFound)
	t.Logf("Cluster List Completed")
}

func clusterUpdate(t *testing.T) {
	client, err := clients.NewContainerInfraV1Client()
	if err != nil {
		t.Fatalf("Unable to create container-infra client: %v", err)
	}
	client.Microversion = "1.5"

	clusterName := testName
	updateOpts := []clusters.UpdateOptsBuilder{
		clusters.UpdateOpts{
			Path:  "/node_count",
			Value: "2",
			Op:    "replace",
		},
	}

	updateResult := clusters.Update(client, clusterName, updateOpts)
	th.AssertNoErr(t, updateResult.Err)
	if len(updateResult.Header["X-Openstack-Request-Id"]) > 0 {
		t.Logf("Cluster Update Request ID: %s", updateResult.Header["X-Openstack-Request-Id"][0])
	}

	clusterID, err := updateResult.Extract()
	th.AssertNoErr(t, err)

	cluster, err := clusters.Get(client, clusterID).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, clusterName, cluster.Name)
	tools.PrintResource(t, cluster)

	// Revert back to original
	updateOpts = []clusters.UpdateOptsBuilder{
		clusters.UpdateOpts{
			Path:  "/node_count",
			Value: "1",
			Op:    "replace",
		},
	}

	updateResult = clusters.Update(client, clusterName, updateOpts)
	th.AssertNoErr(t, err)
	t.Logf("Headers: %+v", updateResult.Header)
	location := updateResult.Header.Get("Location")
	th.AssertEquals(t, true, location != "")
	t.Logf("Cluster Update Location: %s", location)

	clusterID, err = updateResult.Extract()
	cluster, err = clusters.Get(client, clusterID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, clusterName, cluster.Name)
	tools.PrintResource(t, cluster)

	t.Logf("Cluster Update Completed: %s", clusterName)
}

func clusterDelete(t *testing.T) {
	client, err := clients.NewContainerInfraV1Client()
	if err != nil {
		t.Fatalf("Unable to create container-infra client: %v", err)
	}

	clusterName := testName
	deleteRequestID, err := clusters.Delete(client, clusterName).Extract()
	t.Logf("Delete cluster. RequestID=%s", deleteRequestID)
	th.AssertNoErr(t, err)
	t.Logf("Cluster Delete Completed: %s", clusterName)
}
*/