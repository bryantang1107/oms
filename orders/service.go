package main

import (
	"context"
	"log"

	common "github.com/bryantang1107/commons"
	pb "github.com/bryantang1107/commons/api"
)

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return common.ErrNoItems
	}
	mergedItems := mergeItemsQuantities(p.Items)
	log.Println(mergedItems)

	// validate with stock service
	return nil
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	// remove duplicate items
	merged := make([]*pb.ItemsWithQuantity, 0)
	hMap := make(map[string]int)
	for index, val := range items {
		i, exist := hMap[val.ID]
		if exist {
			merged[i].Quantity += val.Quantity
		} else {
			merged = append(merged, val)
		}
		hMap[val.ID] = index
	}
	return merged
}
