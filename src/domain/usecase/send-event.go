package usecase

import "github.com/matheusteodoro01/go-clean-arch/src/domain/models"

type SendEventInputDto struct {
	Service string
	Event   string
	Date    string
	Data    string
}

type SendEventUseCase struct {
	EventRepository models.EventRepository
}

func NewSendEventUseCase(eventRepository models.EventRepository) *SendEventUseCase {
	return &SendEventUseCase{EventRepository: eventRepository}
}

func (useCase *SendEventUseCase) SendEvent(inputDto *SendEventInputDto) error {
	event := models.NewEvent(inputDto.Service, inputDto.Event, inputDto.Date, inputDto.Data)

	err := useCase.EventRepository.Send(event)
	if err != nil {
		return err
	}

	return nil
}
