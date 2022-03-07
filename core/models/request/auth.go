package request

type Auth struct {
	KeyAuth
	PasswordAuth
}

type KeyAuth struct {
	Key string `json:"license_key,omitempty"`
}

type PasswordAuth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (_ Auth) FromKey(licenseKey string) Auth {
	return Auth{
		KeyAuth{
			Key: licenseKey,
		},
		PasswordAuth{},
	}
}

func (_ Auth) FromUsername(username, password string) Auth {
	return Auth{
		KeyAuth{},
		PasswordAuth{
			Username: username,
			Password: password,
		},
	}
}
