package core_models

type LicenseType string

func (_ LicenseType) Perpetual() LicenseType {
	return "perpetual"
}

func (_ LicenseType) Subscription() LicenseType {
	return "subscription"
}

func (_ LicenseType) TimeLimited() LicenseType {
	return "time-limited"
}

func (_ LicenseType) Consumption() LicenseType {
	return "consumption"
}

func (_ LicenseType) Unknown() LicenseType {
	return "unknown"
}
