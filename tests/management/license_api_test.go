package management

import (
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

	t.Run("Disable", func(t *testing.T) {
		resp := c.LicenseApi().DisableLicense(testLicenseID)
		assert.NoError(t, resp.Error)
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.LicenseApi().GetLicense(testLicenseID)
		assert.NoError(t, resp.Error)
		assert.False(t, resp.Value.Enabled)
	})

	t.Run("Enable", func(t *testing.T) {
		resp := c.LicenseApi().EnableLicense(testLicenseID)
		assert.NoError(t, resp.Error)
	})

	t.Run("Update", func(t *testing.T) {
		resp := c.LicenseApi().UpdateLicense(testLicenseID, management_request.UpdateLicenseRequest{
			MaxActivations: 5,
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("List", func(t *testing.T) {
		resp := c.LicenseApi().SearchLicenses(management_request.SearchLicensesRequest{})
		assert.NoError(t, resp.Error)
		assert.Equal(t, 2, resp.Count)
	})
}

func TestUser(t *testing.T) {
	c := Setup()

	var userID int

	t.Run("Assign", func(t *testing.T) {
		resp := c.LicenseApi().AssignUser(testUserLicenseID, management_request.AssignUserToLicenseRequest{
			Email:    "test@abcd.efgh",
			Password: "abcdefgh",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("List", func(t *testing.T) {
		resp := c.LicenseApi().ListLicenseUsers(management_request.SearchUsersRequest{
			EmailContains: "test@abcd.efgh",
		})
		assert.NoError(t, resp.Error)
		require.Equal(t, 1, resp.Count)
		userID = resp.Results[0].Id
	})

	t.Run("Password", func(t *testing.T) {
		resp := c.LicenseApi().SetUserPassword(userID, management_request.SetPasswordRequest{
			Password: "test",
		})
		assert.NoError(t, resp.Error)
	})

	t.Run("Unassign", func(t *testing.T) {
		resp := c.LicenseApi().UnassignUser(testUserLicenseID, management_request.UnassignRequest{
			LicenseUserId: userID,
		})
		assert.NoError(t, resp.Error)
	})
}

func TestLicenseCustomFields(t *testing.T) {
	c := Setup()

	var id int

	t.Run("List", func(t *testing.T) {
		resp := c.LicenseApi().SearchLicenseCustomFields(management_request.SearchLicenseCustomFieldsRequest{
			License: testLicenseID,
		})
		require.NoError(t, resp.Error)
		require.Equal(t, 1, resp.Count)
		id = resp.Results[0].Id
	})

	t.Run("Get", func(t *testing.T) {
		resp := c.LicenseApi().GetLicenseCustomField(id)
		assert.NoError(t, resp.Error)
	})

	t.Run("Update", func(t *testing.T) {
		resp := c.LicenseApi().UpdateLicenseCustomField(id, management_request.UpdateCustomLicenseFieldRequest{
			Value: "test",
		})
		assert.NoError(t, resp.Error)
	})
}
