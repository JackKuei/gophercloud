package clustertemplates

import (
	"github.com/gophercloud/gophercloud"
)

var apiVersion = "v1"
var apiName = "clustertemplates"

func commonURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiName)
}

func idURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiName, id)
}

func createURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return idURL(client, id)
}

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("clustertemplates")
}

func updateURL(client *gophercloud.ServiceClient, id string) string {
	return idURL(client, id)
}
