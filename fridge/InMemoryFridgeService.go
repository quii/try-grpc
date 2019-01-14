package fridge

import "context"

type InMemoryFridgeService struct {
	items []string
}

func NewInMemoryFridgeService(items []string) *InMemoryFridgeService {
	return &InMemoryFridgeService{items: items}
}

func (f *InMemoryFridgeService) GetItems(ctx context.Context, req *GetItemsRequest) (*ItemsResponse, error) {
	return f.makeItemsResponse(), nil
}

func (f *InMemoryFridgeService) AddItem(ctx context.Context, i *Item) (*ItemsResponse, error) {
	f.items = append(f.items, i.Name)
	return f.makeItemsResponse(), nil
}

func (f *InMemoryFridgeService) makeItemsResponse() *ItemsResponse {
	res := &ItemsResponse{}
	for _, i := range f.items {
		res.Items = append(res.Items, &Item{Name: i})
	}
	return res
}
