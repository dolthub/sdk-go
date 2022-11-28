package management_request

type PatchLicenseFeaturesRequest struct {
	ProductFeatures []LicenseFeaturePatch `json:"product_features,omitempty"`
}
