package main

import "payment-apps-backend/delivery"

func main() {
	delivery.NewServer().Run()
}
