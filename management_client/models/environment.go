package management_models

type Environment string

func (_ Environment) Windows() Environment {
	return "win"
}

func (_ Environment) Windows32() Environment {
	return "win32"
}

func (_ Environment) Windows64() Environment {
	return "win64"
}

func (_ Environment) Mac() Environment {
	return "mac"
}

func (_ Environment) Linux() Environment {
	return "linux"
}

func (_ Environment) Linux32() Environment {
	return "linux32"
}

func (_ Environment) Linux64() Environment {
	return "linux64"
}
