// +build acceptance containerinfra

package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/quotas"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestQuotasCRUD(t *testing.T) {
	client, err := clients.NewContainerInfraV1Client()
	th.AssertNoErr(t, err)

	quota, err := CreateQuota(t, client)
	th.AssertNoErr(t, err)
	t.Log(quota.ProjectID)

	defer DeleteQuota(t, client, quota.ProjectID, "Cluster")

	// Test quotas list
	allPages, err := quotas.List(client, nil).AllPages()
	th.AssertNoErr(t, err)

	allQuotas, err := quotas.ExtractQuotas(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, v := range allQuotas {
		if v.ProjectID == quota.ProjectID {
			found = true
		}
	}

	th.AssertEquals(t, found, true)

	// Test quota update
	updateOpts := quotas.UpdateOpts{
		ProjectID: quota.ProjectID,
		Resource:  quota.Resource,
		HardLimit: quota.HardLimit - 1,
	}

	updateQuota, err := quotas.Update(client, quota.ProjectID, "Cluster", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, quota.ProjectID, updateQuota.ProjectID)
	th.AssertEquals(t, quota.Resource, updateQuota.Resource)
	th.AssertEquals(t, quota.HardLimit-1, updateQuota.HardLimit)
	tools.PrintResource(t, updateQuota)
}
