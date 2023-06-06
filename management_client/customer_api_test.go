package management_client

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	"testing"
)

func TestCustomer(t *testing.T) {
	c := Setup()

	var id int64

	t.Run("Create", func(t *testing.T) {
		resp := c.CustomerApi().CreateCustomer(management_request.CreateCustomerRequest{
			Email:     "ttt@ttt.ttt",
			FirstName: "Test",
		})
		require.NoError(t, resp.Error)
		id = resp.Value.Id
	})

	t.Run("List", func(t *testing.T) {
		resp := c.CustomerApi().ListCustomers(management_request.SearchCustomersRequest{
			Email: "ttt@ttt.ttt",
		})
		assert.NoError(t, resp.Error)
		assert.Equal(t, 1, resp.Count)
	})

	t.Run("Edit", func(t *testing.T) {
		resp := c.CustomerApi().EditCustomer(id, management_request.CreateCustomerRequest{
			LastName: "abcd",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.CustomerApi().ShowCustomer(id)
		assert.NoError(t, resp.Error)
		assert.Equal(t, "abcd", resp.Value.LastName)
	})

	t.Run("Delete", func(t *testing.T) {
		err := c.CustomerApi().DeleteCustomer(id)
		assert.NoError(t, err)
	})
}

func TestCustomerLabels(t *testing.T) {
	c := Setup()

	var id int64

	t.Run("Create", func(t *testing.T) {
		resp := c.CustomerApi().CreateCustomerLabel(management_request.CreateCustomerLabelRequest{
			Label: "abcd",
			Color: "efgh",
		})
		require.NoError(t, resp.Error)
		id = resp.Value.Id
	})

	t.Run("List", func(t *testing.T) {
		resp := c.CustomerApi().ListCustomerLabels(management_request.SearchRequest{})
		assert.NoError(t, resp.Error)
		assert.Equal(t, 1, resp.Count)
	})

	t.Run("Edit", func(t *testing.T) {
		resp := c.CustomerApi().EditCustomerLabel(id, management_request.CreateCustomerLabelRequest{
			Color: "1234",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.CustomerApi().ShowCustomerLabel(id)
		assert.NoError(t, resp.Error)
		assert.Equal(t, "1234", resp.Value.Color)
	})

	t.Run("Delete", func(t *testing.T) {
		err := c.CustomerApi().DeleteCustomerLabel(id)
		assert.NoError(t, err)
	})
}

func TestCustomerAccounts(t *testing.T) {
	c := Setup()

	var id int64

	t.Run("Create", func(t *testing.T) {
		resp := c.CustomerApi().CreateCustomerAccount(management_request.CreateCustomerAccountRequest{
			Code: "abcd",
			Name: "efgh",
		})
		require.NoError(t, resp.Error)
		id = resp.Value.Id
	})

	t.Run("List", func(t *testing.T) {
		resp := c.CustomerApi().ListCustomerAccounts(management_request.SearchRequest{})
		assert.NoError(t, resp.Error)
		assert.Equal(t, 1, resp.Count)
	})

	t.Run("Edit", func(t *testing.T) {
		resp := c.CustomerApi().EditCustomerAccount(id, management_request.CreateCustomerAccountRequest{
			Description: "test",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.CustomerApi().ShowCustomerAccount(id)
		assert.NoError(t, resp.Error)
		assert.Equal(t, "test", resp.Value.Description)
	})

	t.Run("Delete", func(t *testing.T) {
		err := c.CustomerApi().DeleteCustomerAccount(id)
		assert.NoError(t, err)
	})
}
