package main

import "context"

type OrderService interface {
	CreateOrder(context context.Context) error
}

type OrderStore interface {
	Create(context context.Context) error
}
