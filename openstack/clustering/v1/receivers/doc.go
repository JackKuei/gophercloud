/*
Package receivers provides information and interaction with the receivers through
the OpenStack Clustering service.

Example to Create a Receiver

	createOpts := receivers.CreateOpts{
		Action:     "CLUSTER_DEL_NODES",
		ClusterID:  "b7b870ee-d3c5-4a93-b9d7-846c53b2c2dc",
		Name:       "test_receiver",
		Type:       "webhook",
	}

	receiver, err := receivers.Create(serviceClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Println("receiver", receiver)

*/
package receivers
