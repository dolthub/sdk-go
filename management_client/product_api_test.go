package management_client

import (
	"github.com/stretchr/testify/assert"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	"testing"
)

const (
	testProductID = 1686069332136783
)

func TestProduct(t *testing.T) {
	c := Setup()

	t.Run("List", func(t *testing.T) {
		resp := c.ProductApi().ListProducts(management_request.SearchProductsRequest{})
		assert.NoError(t, resp.Error)
		assert.Equal(t, 2, resp.Count)
	})

	t.Run("Show", func(t *testing.T) {
		resp := c.ProductApi().ShowProduct(testProductID)
		assert.NoError(t, resp.Error)
	})
}

func TestInstallationFile(t *testing.T) {
	c := Setup()

	var id int

	t.Run("Create", func(t *testing.T) {
		resp := c.ProductApi().CreateInstallationFile(management_request.CreateInstallationFileRequest{
			FullLink: "link",
			Product:  testProductID,
			Version:  "v1",
		})
		assert.NoError(t, resp.Error)
		id = resp.Value.Id
	})

	t.Run("List", func(t *testing.T) {
		resp2 := c.ProductApi().ListInstallationFiles(management_request.SearchInstallationFilesRequest{
			Product: testProductID,
		})
		assert.NoError(t, resp2.Error)
	})

	t.Run("Show", func(t *testing.T) {
		resp3 := c.ProductApi().ShowInstallationFile(id)
		assert.NoError(t, resp3.Error)
	})
}

func TestCustomFields(t *testing.T) {
	c := Setup()

	var id int

	t.Run("Create", func(t *testing.T) {
		resp := c.ProductApi().CreateProductCustomField(management_request.CreateProductCustomFieldRequest{
			Name:         "field",
			DefaultValue: "test",
			Product:      testProductID,
		})
		assert.NoError(t, resp.Error)
		id = resp.Value.Id
	})

	t.Run("Show1", func(t *testing.T) {
		respShow := c.ProductApi().ShowProductCustomField(id)
		assert.NoError(t, respShow.Error)
		assert.Equal(t, "test", respShow.Value.DefaultValue)
	})

	t.Run("Update", func(t *testing.T) {
		respUpdate := c.ProductApi().UpdateProductCustomField(id, management_request.UpdateProductCustomFieldRequest{
			Name:         "field",
			DefaultValue: "abcd",
		})
		assert.NoError(t, respUpdate.Error)
	})

	t.Run("Show2", func(t *testing.T) {
		respShow := c.ProductApi().ShowProductCustomField(id)
		assert.NoError(t, respShow.Error)
		assert.Equal(t, "abcd", respShow.Value.DefaultValue)
	})

	t.Run("Delete", func(t *testing.T) {
		err := c.ProductApi().DeleteProductCustomField(id)
		assert.NoError(t, err)
	})

	t.Run("List", func(t *testing.T) {
		respList := c.ProductApi().ListProductCustomFields(management_request.SearchProductCustomFieldsRequest{
			Product: testProductID,
		})
		assert.NoError(t, respList.Error)
		assert.Equal(t, 0, respList.Count)
	})
}
