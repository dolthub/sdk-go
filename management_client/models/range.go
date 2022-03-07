package management_models

type Range string

func (_ Range) Last30() Range {
	return "last30"
}

func (_ Range) Last60() Range {
	return "last60"
}

func (_ Range) Last180() Range {
	return "last180"
}

func (_ Range) Last365() Range {
	return "last365"
}
