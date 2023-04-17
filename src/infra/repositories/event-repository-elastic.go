package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/matheusteodoro01/go-clean-arch/src/domain/models"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
)

type EventRepositoryElastic struct {
	Client *elasticsearch.Client
}

func NewEventRepositoryElastic(client *elasticsearch.Client) (*EventRepositoryElastic, error) {
	if client == nil {
		return nil, errors.New("client cannot be nil")
	}

	return &EventRepositoryElastic{Client: client}, nil
}

func (client *EventRepositoryElastic) Send(event *models.Event) error {
	// Codifica o evento para JSON
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Cria um buffer para armazenar o JSON codificado
	buf := bytes.NewBuffer(eventJSON)

	// Cria um reader a partir do buffer
	reader := bytes.NewReader(buf.Bytes())

	// Cria uma requisição de bulk com o reader criado acima
	req := esapi.BulkRequest{
		Index: "test",
		Body:  reader,
	}

	// Executa a requisição
	res, err := req.Do(context.Background(), client.Client)
	if err != nil {
		return err
	}

	// Verifica se houve algum erro na resposta
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[%s] Erro ao enviar evento: %s", res.Status(), res.String())
	}

	return nil
}
