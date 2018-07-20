package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/quotas"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

const projectID = "aa5436ab58144c768ca4e9d2e9f5c3b2"
const resourceID = "746e779a-751a-456b-a3e9-c883d734946f"
const resourceType = "Cluster"

const requestUUID = "req-781e9bdc-4163-46eb-91c9-786c53188bbb"

var CreateResponse = fmt.Sprintf(`
										{
										   "resource": "Cluster",
										   "created_at": "2017-01-17T17:35:48+00:00",
										   "updated_at": null,
										   "hard_limit": 1,
										   "project_id": "%s",
										   "id": 26
										}`, projectID)

var UpdateResponse = fmt.Sprintf(`
										{
										   "resource": "Cluster",
										   "created_at": "2017-01-17T17:35:49+00:00",
										   "updated_at": "2017-01-17T17:38:20+00:00",
										   "hard_limit": 10,
										   "project_id": "%s",
										   "id": 26
										}`, projectID)

var ExpectedQuota = quotas.Quotas{
	Resource:  resourceType,
	ProjectID: projectID,
	HardLimit: 10,
	CreatedAt: time.Date(2016, 8, 29, 6, 51, 31, 0, time.UTC),
	UpdatedAt: time.Date(2016, 8, 29, 6, 53, 24, 0, time.UTC),
}

func HandleCreateQuotaSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/quotas", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, CreateResponse)
	})
}

func HandleDeleteQuotaSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/quotas/"+projectID+"/"+resourceType, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusNoContent)
	})
}

var ListResponse = fmt.Sprintf(`
{
	"quotas": [
		{
         "resource": "Cluster",
         "created_at": "2017-01-17T17:35:49+00:00",
         "updated_at": "2017-01-17T17:38:21+00:00",
         "hard_limit": 10,
         "project_id": "%s",
         "id": 26
		},
	]
}`, projectID)

var ExpectedQuotas = []quotas.Quotas{ExpectedQuota}

func HandleListQuotasSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/quotas", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})
}

func HandleUpdateQuotaSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/v1/quotas/"+projectID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("X-OpenStack-Request-Id", requestUUID)
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, UpdateResponse)
	})
}
