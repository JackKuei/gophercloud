package quotas

import (
	"github.com/gophercloud/gophercloud"
)

var apiName = "quotas"

func commonURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiName)
}

func resourceURL(client *gophercloud.ServiceClient, id string, resource string) string {
	return client.ServiceURL(apiName, id, resource)
}

func createURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
}

func getURL(client *gophercloud.ServiceClient, id string, resource string) string {
	return resourceURL(client, id, resource)
}

func listURL(client *gophercloud.ServiceClient) string {
	return commonURL(client)
}

func updateURL(client *gophercloud.ServiceClient, id string, resource string) string {
	return resourceURL(client, id, resource)
}

func deleteURL(client *gophercloud.ServiceClient, id string, resource string) string {
	return resourceURL(client, id, resource)
}
