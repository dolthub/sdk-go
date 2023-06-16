package management

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	"testing"
)

func TestCustomer(t *testing.T) {
	c := Setup()

	var id int64

	ctx := context.Background()

	t.Run("Create", func(t *testing.T) {
		resp := c.CustomerApi().CreateCustomer(ctx, management_request.CreateCustomerRequest{
			Email:     "ttt@ttt.ttt",
			FirstName: "Test",
		})
		require.NoError(t, resp.Error)
		id = resp.Value.Id
	})

	t.Run("List", func(t *testing.T) {
		resp := c.CustomerApi().ListCustomers(ctx, management_request.SearchCustomersRequest{
			Email: "ttt@ttt.ttt",
		})
		assert.NoError(t, resp.Error)
		assert.Equal(t, 1, resp.Count)
	})

	t.Run("Edit", func(t *testing.T) {
		resp := c.CustomerApi().EditCustomer(ctx, id, management_request.CreateCustomerRequest{
			LastName: "abcd",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.CustomerApi().ShowCustomer(ctx, id)
		assert.NoError(t, resp.Error)
		assert.Equal(t, "abcd", resp.Value.LastName)
	})

	t.Run("Delete", func(t *testing.T) {
		err := c.CustomerApi().DeleteCustomer(ctx, id)
		assert.NoError(t, err)
	})
}

func TestCustomerLabels(t *testing.T) {
	c := Setup()

	var id int64

	ctx := context.Background()

	t.Run("Create", func(t *testing.T) {
		resp := c.CustomerApi().CreateCustomerLabel(ctx, management_request.CreateCustomerLabelRequest{
			Label: "abcd",
			Color: "efgh",
		})
		require.NoError(t, resp.Error)
		id = resp.Value.Id
	})

	t.Run("List", func(t *testing.T) {
		resp := c.CustomerApi().ListCustomerLabels(ctx, management_request.SearchRequest{})
		assert.NoError(t, resp.Error)
		assert.Equal(t, 1, resp.Count)
	})

	t.Run("Edit", func(t *testing.T) {
		resp := c.CustomerApi().EditCustomerLabel(ctx, id, management_request.CreateCustomerLabelRequest{
			Color: "1234",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.CustomerApi().ShowCustomerLabel(ctx, id)
		assert.NoError(t, resp.Error)
		assert.Equal(t, "1234", resp.Value.Color)
	})

	t.Run("Delete", func(t *testing.T) {
		err := c.CustomerApi().DeleteCustomerLabel(ctx, id)
		assert.NoError(t, err)
	})
}

func TestCustomerAccounts(t *testing.T) {
	c := Setup()

	var id int64

	ctx := context.Background()

	t.Run("Create", func(t *testing.T) {
		resp := c.CustomerApi().CreateCustomerAccount(ctx, management_request.CreateCustomerAccountRequest{
			Code: "abcd",
			Name: "efgh",
		})
		require.NoError(t, resp.Error)
		id = resp.Value.Id
	})

	t.Run("List", func(t *testing.T) {
		resp := c.CustomerApi().ListCustomerAccounts(ctx, management_request.SearchRequest{})
		assert.NoError(t, resp.Error)
		assert.Equal(t, 1, resp.Count)
	})

	t.Run("Edit", func(t *testing.T) {
		resp := c.CustomerApi().EditCustomerAccount(ctx, id, management_request.CreateCustomerAccountRequest{
			Description: "test",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.CustomerApi().ShowCustomerAccount(ctx, id)
		assert.NoError(t, resp.Error)
		assert.Equal(t, "test", resp.Value.Description)
	})

	t.Run("Delete", func(t *testing.T) {
		err := c.CustomerApi().DeleteCustomerAccount(ctx, id)
		assert.NoError(t, err)
	})
}
