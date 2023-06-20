```go
ctx := context.Background()


	resp := c.CustomerApi().CreateCustomer(ctx, management_request.CreateCustomerRequest{
		Email:     "ttt@ttt.ttt",
		FirstName: "Test",})
	require.NoError(t, resp.Error)
	id = resp.Value.Id



	resp := c.CustomerApi().ListCustomers(ctx, management_request.SearchCustomersRequest{
		Email: "ttt@ttt.ttt",})
	assert.NoError(t, resp.Error)
	assert.Equal(t, 1, resp.Count)



	resp := c.CustomerApi().EditCustomer(ctx, id, management_request.CreateCustomerRequest{
		LastName: "abcd",})
	assert.NoError(t, resp.Error)



	resp := c.CustomerApi().ShowCustomer(ctx, id)
	assert.NoError(t, resp.Error)
	assert.Equal(t, "abcd", resp.Value.LastName)



	err := c.CustomerApi().DeleteCustomer(ctx, id)
	assert.NoError(t, err)



```