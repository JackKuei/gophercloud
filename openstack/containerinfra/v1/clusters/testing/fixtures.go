package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clusters"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

const clusterUUID = "746e779a-751a-456b-a3e9-c883d734946f"
const clusterUUID2 = "846e779a-751a-456b-a3e9-c883d734945f"
const requestUUID = "req-781e9bdc-4163-46eb-91c9-786c53188bbb"

var ClusterCreateResponse = fmt.Sprintf(`
											{
												"uuid":"%s"
											}`, clusterUUID)

var ClusterListResponse = fmt.Sprintf(`
{
	"clusters": [
		{
			"status":"CREATE_COMPLETE",
			"uuid":"%s",
			"links":[
			  {
				 "href":"http://10.164.180.104:9511/v1/clusters/746e779a-751a-456b-a3e9-c883d734946f",
				 "rel":"self"
			  },
			  {
				 "href":"http://10.164.180.104:9511/clusters/746e779a-751a-456b-a3e9-c883d734946f",
				 "rel":"bookmark"
			  }
			],
			"stack_id":"9c6f1169-7300-4d08-a444-d2be38758719",
			"created_at":"2016-08-29T06:51:31+00:00",
			"api_address":"https://172.24.4.6:6443",
			"discovery_url":"https://discovery.etcd.io/cbeb580da58915809d59ee69348a84f3",
			"updated_at":"2016-08-29T06:53:24+00:00",
			"master_count":1,
			"coe_version": "v1.2.0",
			"keypair":"my-keypair",
			"cluster_template_id":"0562d357-8641-4759-8fed-8173f02c9633",
			"master_addresses":[
			  "172.24.4.6"
			],
			"node_count":1,
			"node_addresses":[
			  "172.24.4.13"
			],
			"status_reason":"Stack CREATE completed successfully",
			"create_timeout":60,
			"name":"k8s"
		},
		{
			"status":"CREATE_COMPLETE",
			"uuid":"%s",
			"links":[
			  {
				 "href":"http://10.164.180.104:9511/v1/clusters/746e779a-751a-456b-a3e9-c883d734946f",
				 "rel":"self"
			  },
			  {
				 "href":"http://10.164.180.104:9511/clusters/746e779a-751a-456b-a3e9-c883d734946f",
				 "rel":"bookmark"
			  }
			],
			"stack_id":"9c6f1169-7300-4d08-a444-d2be38758719",
			"created_at":null,
			"api_address":"https://172.24.4.6:6443",
			"discovery_url":"https://discovery.etcd.io/cbeb580da58915809d59ee69348a84f3",
			"updated_at":null,
			"master_count":1,
			"coe_version": "v1.2.0",
			"keypair":"my-keypair",
			"cluster_template_id":"0562d357-8641-4759-8fed-8173f02c9633",
			"master_addresses":[
			  "172.24.4.6"
			],
			"node_count":1,
			"node_addresses":[
			  "172.24.4.13"
			],
			"status_reason":"Stack CREATE completed successfully",
			"create_timeout":60,
			"name":"k8s"
		}
	]
}`, clusterUUID, clusterUUID2)

var UpdateResponse = fmt.Sprintf(`
									{
										"uuid":"%s"
									}`, clusterUUID)

var ExpectedCluster = clusters.Cluster{
	Status: "CREATE_COMPLETE",
	UUID:   clusterUUID,
	Links: []gophercloud.Link{
		{
			Href: "http://10.164.180.104:9511/v1/clusters/746e779a-751a-456b-a3e9-c883d734946f",
			Rel:  "self",
		},
		{
			Href: "http://10.164.180.104:9511/clusters/746e779a-751a-456b-a3e9-c883d734946f",
			Rel:  "bookmark",
		},
	},
	StackID:           "9c6f1169-7300-4d08-a444-d2be38758719",
	CreatedAt:         time.Date(2016, 8, 29, 6, 51, 31, 0, time.UTC),
	APIAddress:        "https://172.24.4.6:6443",
	DiscoveryURL:      "https://discovery.etcd.io/cbeb580da58915809d59ee69348a84f3",
	UpdatedAt:         time.Date(2016, 8, 29, 6, 53, 24, 0, time.UTC),
	MasterCount:       1,
	COEVersion:        "v1.2.0",
	KeyPair:           "my-keypair",
	ClusterTemplateID: "0562d357-8641-4759-8fed-8173f02c9633",
	MasterAddresses:   []string{"172.24.4.6"},
	NodeCount:         1,
	NodeAddresses:     []string{"172.24.4.13"},
	StatusReason:      "Stack CREATE completed successfully",
	CreateTimeout:     60,
	Name:              "k8s",
}

var ExpectedClusterUUID = clusters.UUID{
	UUID: clusterUUID,
}

func HandleCreateClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ClusterCreateResponse)
	})
}

func HandleDeleteClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters/"+clusterUUID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusNoContent)
	})
}

var ExpectedClusters = []clusters.Cluster{ExpectedCluster}

func HandleListClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ClusterListResponse)
	})
}

func HandleUpdateClusterSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/clusters/"+clusterUUID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, UpdateResponse)
	})
}
