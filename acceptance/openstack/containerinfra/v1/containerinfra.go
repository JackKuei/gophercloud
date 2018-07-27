package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clustertemplates"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/quotas"
)

// CreateClusterTemplate will create a random cluster tempalte. An error will be returned if the
// cluster-template could not be created.
func CreateClusterTemplate(t *testing.T, client *gophercloud.ServiceClient) (*clustertemplates.ClusterTemplate, error) {
	choices, err := clients.AcceptanceTestChoicesFromEnv()
	if err != nil {
		return nil, err
	}

	name := tools.RandomString("TESTACC-", 8)
	t.Logf("Attempting to create cluster template: %s", name)

	boolFalse := false
	createOpts := clustertemplates.CreateOpts{
		Name:                name,
		MasterFlavorID:      "m1.small",
		Public:              &boolFalse,
		ServerType:          "vm",
		ExternalNetworkID:   choices.ExternalNetworkID,
		ImageID:             choices.MagnumImageID,
		RegistryEnabled:     &boolFalse,
		DockerStorageDriver: "devicemapper",
		COE:                 "swarm",
		FlavorID:            choices.FlavorID,
		MasterLBEnabled:     &boolFalse,
		DNSNameServer:       "8.8.8.8",
		FloatingIPEnabled:   &boolFalse,
	}

	res := clustertemplates.Create(client, createOpts)
	if res.Err != nil {
		return nil, res.Err
	}

	requestID := res.Header.Get("X-OpenStack-Request-Id")
	th.AssertEquals(t, true, requestID != "")

	t.Logf("Cluster Template %s request ID: %s", name, requestID)

	clusterTemplate, err := res.Extract()
	if err != nil {
		return nil, err
	}

	t.Logf("Successfully created cluster template: %s", clusterTemplate.Name)

	tools.PrintResource(t, clusterTemplate)
	tools.PrintResource(t, clusterTemplate.CreatedAt)

	th.AssertEquals(t, name, clusterTemplate.Name)
	th.AssertEquals(t, choices.ExternalNetworkID, clusterTemplate.ExternalNetworkID)
	th.AssertEquals(t, choices.MagnumImageID, clusterTemplate.ImageID)

	return clusterTemplate, nil
}

/*
// DeleteClusterTemplate will delete a given cluster-template. A fatal error will occur if the
// cluster-template could not be deleted. This works best as a deferred function.
func DeleteClusterTemplate(t *testing.T, client *gophercloud.ServiceClient, id string) {
	t.Logf("Attempting to delete cluster-template: %s", id)

	err := clustertemplates.Delete(client, id).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting cluster-template %s: %s:", id, err)
	}

	t.Logf("Successfully deleted cluster-template: %s", id)

	return
}
*/

// CreateQuotas will create a random quota. An error will be returned if the
// quota could not be created.
func CreateQuota(t *testing.T, client *gophercloud.ServiceClient) (*quotas.Quotas, error) {
	choices, err := clients.AcceptanceTestChoicesFromEnv()
	if err != nil {
		return nil, err
	}

	name := tools.RandomString("TESTACC-", 8)
	t.Logf("Attempting to create quota: %s", name)

	createOpts := quotas.CreateOpts{
		Resource:  "Cluster",
		ProjectID: choices.ProjectName,
		HardLimit: 10,
	}

	res := quotas.Create(client, createOpts)
	if res.Err != nil {
		return nil, res.Err
	}

	requestID := res.Header.Get("X-OpenStack-Request-Id")
	th.AssertEquals(t, true, requestID != "")

	t.Logf("Quota %s request ID: %s", name, requestID)

	quota, err := res.Extract()
	if err != nil {
		return nil, err
	}

	t.Logf("Successfully created quota: %s", quota.ProjectID)

	tools.PrintResource(t, quota)

	th.AssertEquals(t, name, quota.ProjectID)
	th.AssertEquals(t, "Cluster", quota.Resource)
	th.AssertEquals(t, 10, quota.HardLimit)

	return quota, nil
}

// DeleteQuota will delete a given quota. A fatal error will occur if the
// quota could not be deleted. This works best as a deferred function.
func DeleteQuota(t *testing.T, client *gophercloud.ServiceClient, id string, resource string) {
	t.Logf("Attempting to delete quota: %s", id)

	err := quotas.Delete(client, id, resource).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting quota %s: %s:", id, err)
	}

	t.Logf("Successfully deleted quota: %s", id)

	return
}
