package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clustertemplates"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestCreateClusterTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleCreateClusterTemplateSuccessfully(t)

	opts := clustertemplates.CreateOpts{
		Name:                "test-cluster-template",
		Labels:              map[string]string{},
		FixedSubnet:         "",
		MasterFlavorID:      "",
		NoProxy:             "10.0.0.0/8,172.0.0.0/8,192.0.0.0/8,localhost",
		HTTPSProxy:          "http://10.164.177.169:8080",
		TLSDisabled:         false,
		KeyPairID:           "kp",
		Public:              false,
		HTTPProxy:           "http://10.164.177.169:8080",
		DockerVolumeSize:    3,
		ServerType:          "vm",
		ExternalNetworkID:   "public",
		ImageID:             "fedora-atomic-latest",
		VolumeDriver:        "cinder",
		RegistryEnabled:     false,
		DockerStorageDriver: "devicemapper",
		NetworkDriver:       "flannel",
		FixedNetwork:        "",
		COE:                 "kubernetes",
		FlavorID:            "m1.small",
		MasterLBEnabled:     true,
		DNSNameServer:       "8.8.8.8",
	}

	res := clustertemplates.Create(fake.ServiceClient(), opts)
	th.AssertNoErr(t, res.Err)

	requestID := res.Header.Get("X-OpenStack-Request-Id")
	th.AssertEquals(t, "req-781e9bdc-4163-46eb-91c9-786c53188bbb", requestID)

	actual, err := res.Extract()
	th.AssertNoErr(t, err)

	actual.CreatedAt = actual.CreatedAt.UTC()
	th.AssertDeepEquals(t, ExpectedClusterTemplate, *actual)
}

func TestCreateClusterTemplateEmptyTime(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleCreateClusterTemplateEmptyTimeSuccessfully(t)

	opts := clustertemplates.CreateOpts{
		Name:                "test-cluster-template",
		Labels:              map[string]string{},
		FixedSubnet:         "",
		MasterFlavorID:      "",
		NoProxy:             "10.0.0.0/8,172.0.0.0/8,192.0.0.0/8,localhost",
		HTTPSProxy:          "http://10.164.177.169:8080",
		TLSDisabled:         false,
		KeyPairID:           "kp",
		Public:              false,
		HTTPProxy:           "http://10.164.177.169:8080",
		DockerVolumeSize:    3,
		ServerType:          "vm",
		ExternalNetworkID:   "public",
		ImageID:             "fedora-atomic-latest",
		VolumeDriver:        "cinder",
		RegistryEnabled:     false,
		DockerStorageDriver: "devicemapper",
		NetworkDriver:       "flannel",
		FixedNetwork:        "",
		COE:                 "kubernetes",
		FlavorID:            "m1.small",
		MasterLBEnabled:     true,
		DNSNameServer:       "8.8.8.8",
	}

	actual, err := clustertemplates.Create(fake.ServiceClient(), opts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedClusterTemplate_EmptyTime, *actual)
}

func TestDeleteClusterTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleDeleteClusterTemplateSuccessfully(t)

	uuid, err := clustertemplates.Delete(fake.ServiceClient(), "6dc6d336e3fc4c0a951b5698cd1236ee").Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "req-781e9bdc-4163-46eb-91c9-786c53188bbb", uuid)
}

func TestListClusterTemplates(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleListClusterTemplateSuccessfully(t)

	count := 0

	clustertemplates.List(fake.ServiceClient(), clustertemplates.ListOpts{Limit: 2}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := clustertemplates.ExtractClusterTemplates(page)
		th.AssertNoErr(t, err)
		for idx, _ := range actual {
			actual[idx].CreatedAt = actual[idx].CreatedAt.UTC()
		}
		th.AssertDeepEquals(t, ExpectedClusterTemplates, actual)

		return true, nil
	})
	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestUpdateClusterTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleUpdateClusterTemplateSuccessfully(t)

	updateOpts := []clustertemplates.UpdateOptsBuilder{
		clustertemplates.UpdateOpts{
			Path:  "/master_lb_enabled",
			Value: "True",
			Op:    "replace",
		},
		clustertemplates.UpdateOpts{
			Path:  "/registry_enabled",
			Value: "True",
			Op:    "replace",
		},
	}
	res := clustertemplates.Update(fake.ServiceClient(), "7d85f602-a948-4a30-afd4-e84f47471c15", updateOpts)
	th.AssertNoErr(t, res.Err)

	actual, err := res.Extract()
	th.AssertNoErr(t, err)
	actual.CreatedAt = actual.CreatedAt.UTC()
	th.AssertDeepEquals(t, ExpectedUpdateClusterTemplate, *actual)
}

func TestUpdateClusterTemplateEmptyTime(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleUpdateClusterTemplateEmptyTimeSuccessfully(t)

	updateOpts := []clustertemplates.UpdateOptsBuilder{
		clustertemplates.UpdateOpts{
			Path:  "/master_lb_enabled",
			Value: "True",
			Op:    "replace",
		},
		clustertemplates.UpdateOpts{
			Path:  "/registry_enabled",
			Value: "True",
			Op:    "replace",
		},
	}

	actual, err := clustertemplates.Update(fake.ServiceClient(), "7d85f602-a948-4a30-afd4-e84f47471c15", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedUpdateClusterTemplate_EmptyTime, *actual)
}
