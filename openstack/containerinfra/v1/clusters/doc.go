/*
Package clusters contains functionality for working with Magnum Cluster resources.

Example to Create a Cluster

	createOpts := clusters.CreateOpts{
		Name:              "k8s",
		DiscoveryURL:      "",
		MasterCount:       1,
		ClusterTemplateID: "0562d357-8641-4759-8fed-8173f02c9633",
		NodeCount:         2,
		CreateTimeout:     60,
		KeyPair:           "my_keypair",
		MasterFlavorID:    "",
		Labels:            map[string]string{},
		FlavorID:          "",
	}

	cluster, err := clusters.Create(serviceClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Get a Cluster

	clusterName := "cluster123"
	cluster, err := clusters.Get(serviceClient, clusterName).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cluster)

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

Example to Update a Cluster

	updateOpts := []clusters.UpdateOptsBuilder{
		clusters.UpdateOpts{
			Path:  "/master_lb_enabled",
			Value: "True",
			Op:    "replace",
		},
		clusters.UpdateOpts{
			Path:  "/registry_enabled",
			Value: "True",
			Op:    "replace",
		},
	}
	clusterUUID, err := clusters.Update(serviceClient, clusterUUID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", clusterUUID)

Example to Delete a Cluster

	clusterUUID := "dc6d336e3fc4c0a951b5698cd1236ee"
	requestUUID, err := clusters.Delete(serviceClient, clusterUUID).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", requestUUID)

*/
package clusters
