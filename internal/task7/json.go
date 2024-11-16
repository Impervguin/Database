package task7

import (
	"DatabaseCourse/internal/task7/domain"
	"encoding/json"
	"io"
	"strings"
)

func (t7s *Task7Storage) GetShortClients() ([]domain.ShortClient, error) {
	clients := make([]domain.ShortClient, 0)
	res := t7s.db.
		Find(&clients)
	if res.Error != nil {
		return nil, res.Error
	}
	return clients, nil
}

func DumpClientsToJson(clients []domain.ShortClient, w io.Writer) error {
	encode := json.NewEncoder(w)
	encode.SetIndent("", "  ") // for pretty printing, remove for compact output
	err := encode.Encode(clients)
	if err != nil {
		return err
	}
	return nil
}

func ReadJsonClients(jsonFile io.Reader) ([]domain.ShortClient, error) {
	var clients []domain.ShortClient
	err := json.NewDecoder(jsonFile).Decode(&clients)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func SetAllNamesUpperJson(jsonIn io.Reader, jsonOut io.Writer) error {
	clients, err := ReadJsonClients(jsonIn)
	if err != nil {
		return err
	}

	for i, client := range clients {
		client.FirstName = strings.ToUpper(client.FirstName)
		client.LastName = strings.ToUpper(client.LastName)
		clients[i] = client
	}

	err = DumpClientsToJson(clients, jsonOut)
	if err != nil {
		return err
	}
	return nil
}

func AddShortClientJson(jsonIn io.Reader, jsonOut io.Writer, client *domain.ShortClient) error {
	clients, err := ReadJsonClients(jsonIn)
	if err != nil {
		return err
	}

	clients = append(clients, *client)

	err = DumpClientsToJson(clients, jsonOut)
	if err != nil {
		return err
	}
	return nil
}
