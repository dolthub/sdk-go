package auth

type Auth struct {
	KeyAuth
	BasicAuth
}

type KeyAuth struct {
	Key string `json:"license_key,omitempty"`
}

type BasicAuth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func FromKey(licenseKey string) Auth {
	return Auth{
		KeyAuth{
			Key: licenseKey,
		},
		BasicAuth{},
	}
}

func FromUsername(username, password string) Auth {
	return Auth{
		KeyAuth{},
		BasicAuth{
			Username: username,
			Password: password,
		},
	}
}
