package usecase

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/matheusteodoro01/go-clean-arch/src/domain/models"
	"github.com/matheusteodoro01/go-clean-arch/src/domain/providers"
)

type SendEventInputDto struct {
	Service string
	Event   string
	Date    string
	Data    string
}

type SendEventUseCase struct {
	MessageSender providers.MessageSender
}

func NewSendEventUseCase(messageSender providers.MessageSender) *SendEventUseCase {
	return &SendEventUseCase{MessageSender: messageSender}
}

func (useCase *SendEventUseCase) Execute(inputDto SendEventInputDto) error {
	event := models.NewEvent(inputDto.Service, inputDto.Event, inputDto.Date, inputDto.Data)

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(event)
	if err != nil {
		log.Fatal(err)
	}

	useCase.MessageSender.Send(buf.Bytes(), "test")

	return nil
}
