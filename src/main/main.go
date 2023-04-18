package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/matheusteodoro01/go-clean-arch/src/domain/usecase"

	"github.com/matheusteodoro01/go-clean-arch/src/infra/controllers"
	"github.com/matheusteodoro01/go-clean-arch/src/infra/providers"
)

func main() {

	// Factory
	kafkaProducer := providers.NewKafkaMessagerProducer()
	sendEventUseCase := usecase.NewSendEventUseCase(kafkaProducer)
	eventController := controllers.NewEventController(sendEventUseCase)

	// Router
	router := chi.NewRouter()
	router.Post("/event", eventController.SendEventController)
	go http.ListenAndServe(":8000", router)

}
