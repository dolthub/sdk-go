package management

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	management_request "gitlab.com/l3178/sdk-go/management_client/models/request"
	"testing"
)

const (
	testLicenseID     = 1686073793260799
	testUserLicenseID = 1686074681611829
)

func TestLicense(t *testing.T) {
	c := Setup()

	ctx := context.Background()

	t.Run("Disable", func(t *testing.T) {
		resp := c.LicenseApi().DisableLicense(ctx, testLicenseID)
		assert.NoError(t, resp.Error)
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.LicenseApi().GetLicense(ctx, testLicenseID)
		assert.NoError(t, resp.Error)
		assert.False(t, resp.Value.Enabled)
	})

	t.Run("Enable", func(t *testing.T) {
		resp := c.LicenseApi().EnableLicense(ctx, testLicenseID)
		assert.NoError(t, resp.Error)
	})

	t.Run("Update", func(t *testing.T) {
		resp := c.LicenseApi().UpdateLicense(ctx, testLicenseID, management_request.UpdateLicenseRequest{
			MaxActivations: 5,
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("List", func(t *testing.T) {
		resp := c.LicenseApi().SearchLicenses(ctx, management_request.SearchLicensesRequest{})
		assert.NoError(t, resp.Error)
		assert.Equal(t, 3, resp.Count)
	})
}

func TestUser(t *testing.T) {
	c := Setup()

	var userID int

	ctx := context.Background()

	t.Run("Assign", func(t *testing.T) {
		resp := c.LicenseApi().AssignUser(ctx, testUserLicenseID, management_request.AssignUserToLicenseRequest{
			Email:    "test@abcd.efgh",
			Password: "abcdefgh",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("List", func(t *testing.T) {
		resp := c.LicenseApi().ListLicenseUsers(ctx, management_request.SearchUsersRequest{
			EmailContains: "test@abcd.efgh",
		})
		assert.NoError(t, resp.Error)
		require.Equal(t, 1, resp.Count)
		userID = resp.Results[0].Id
	})

	t.Run("Password", func(t *testing.T) {
		resp := c.LicenseApi().SetUserPassword(ctx, userID, management_request.SetPasswordRequest{
			Password: "test",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Unassign", func(t *testing.T) {
		resp := c.LicenseApi().UnassignUser(ctx, testUserLicenseID, management_request.UnassignRequest{
			LicenseUserId: userID,
		})
		assert.NoError(t, resp.Error)
	})
}

func TestLicenseCustomFields(t *testing.T) {
	c := Setup()

	var id int

	ctx := context.Background()

	t.Run("List", func(t *testing.T) {
		resp := c.LicenseApi().SearchLicenseCustomFields(ctx, management_request.SearchLicenseCustomFieldsRequest{
			License: testLicenseID,
		})
		require.NoError(t, resp.Error)
		require.Equal(t, 1, resp.Count)
		id = resp.Results[0].Id
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.LicenseApi().GetLicenseCustomField(ctx, id)
		assert.NoError(t, resp.Error)
	})

	t.Run("Update", func(t *testing.T) {
		resp := c.LicenseApi().UpdateLicenseCustomField(ctx, id, management_request.UpdateCustomLicenseFieldRequest{
			Value: "test",
		})
		assert.NoError(t, resp.Error)
	})
}
