/*
Package quotas contains functionality for working with Magnum Quota API.

Example to Create a Quota

	createOpts := quotas.CreateOpts{
		ProjectID: "aa5436ab58144c768ca4e9d2e9f5c3b2",
		Resource:  "Cluster",
		HardLimit: 10,
	}

	quota, err := quotas.Create(serviceClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Get a Quota

	projectID := "aa5436ab58144c768ca4e9d2e9f5c3b2"
	resource := "Cluster"
	quota, err := quotas.Get(serviceClient, projectID, resource).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", quota)

Example to List Quotas

	allPages, err := quotas.ListDetail(serviceClient, nil).AllPages()
	if err != nil {
		panic(err)
	}

	allQuotas, err := quotas.ExtractQuotas(allPages)
	if err != nil {
		panic(err)
	}

	for _, quota := range allQuotas {
		fmt.Printf("%+v\n", quota)
	}

Example to Update a Quota

	updateOpts := quotas.UpdateOpts{
		ProjectID: "aa5436ab58144c768ca4e9d2e9f5c3b2",
		Resource:  "Cluster",
		HardLimit: 20,
	}
	quota, err := quotas.Update(serviceClient, "aa5436ab58144c768ca4e9d2e9f5c3b2", updateOpts).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", quota)

Example to Delete a Quota

	projectID := "dc6d336e3fc4c0a951b5698cd1236ee"
	resource := "Cluster"
	err := quotas.Delete(serviceClient, projectID, resource).ExtractErr()
	if err != nil {
		panic(err)
	}

*/
package quotas
