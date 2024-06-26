package ports

import "github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"

//go:generate mockery --name APIPort
type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
	GetOrder(id string) (domain.Order, error)
}
