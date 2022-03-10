package management_models

type OrderType string

func (_ OrderType) Normal() OrderType {
	return "normal"
}
