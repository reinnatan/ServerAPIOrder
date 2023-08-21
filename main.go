package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Post("/processOrder", func(c *fiber.Ctx) error {
		orderRequest := new(OrderRequest)
		generalResponse := new(GeneralResponse)

		if err := c.BodyParser(orderRequest); err != nil {
			generalResponse.Message = "Invalid request body"
			marshalOrder, _ := json.Marshal(generalResponse)
			return c.Status(400).SendString(string(marshalOrder))
		}

		if orderRequest.OrderId == "" {
			generalResponse.Message = "Invalid order id"
			marshalOrder, _ := json.Marshal(generalResponse)
			return c.Status(400).SendString(string(marshalOrder))

		}

		if orderRequest.OrderStatus == "" {
			generalResponse.Message = "Invalid order status"
			marshalOrder, _ := json.Marshal(generalResponse)
			return c.Status(400).SendString(string(marshalOrder))
		}

		if orderRequest.LastUpdatedTimeStamp == 0 {
			generalResponse.Message = "Invalid last updated timestamp"
			marshalOrder, _ := json.Marshal(generalResponse)
			return c.Status(400).SendString(string(marshalOrder))
		}

		orderRequest.OrderStatus = "Processing"
		order, err := json.Marshal(orderRequest)
		if err != nil {
			generalResponse.Message = "Invalid processing data"
			marshalOrder, _ := json.Marshal(generalResponse)
			return c.Status(412).SendString(string(marshalOrder))
		}

		//currentTimeStamp := time.Now().UnixMilli()
		//orderRequest.LastUpdatedTimeStamp = currentTimeStamp

		return c.SendString(string(order))
	})

	app.Listen(":3000")
}
