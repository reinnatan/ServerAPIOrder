package main

type OrderRequest struct {
	OrderId              string `json:"order_id"`
	OrderDescription     string `json:"order_description"`
	OrderStatus          string `json:"order_status"`
	LastUpdatedTimeStamp int64  `json:"last_updated_timestamp"`
	SpecialOrder         bool   `json:"special_order"`
}
