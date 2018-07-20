/*
Package clusters contains functionality for working with Magnum Cluster resources.

Example to List Clusters

	listOpts := clusters.ListOpts{
		Limit: 20,
	}

	allPages, err := clusters.List(serviceClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allClusters, err := clusters.ExtractClusters(allPages)
	if err != nil {
		panic(err)
	}

	for _, cluster := range allClusters {
		fmt.Printf("%+v\n", cluster)
	}

*/
package quotas
