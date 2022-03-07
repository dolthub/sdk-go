package order

type OrderType string

func (_ OrderType) Normal() OrderType {
	return "normal"
}
