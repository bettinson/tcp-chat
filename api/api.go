package api

import "encoding/json"

type Client struct {
	Name    string `json:"name"`
	Channel string `json:"channel"`
}

func (c *Client) Serialize() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Client) Deserialize(data []byte) error {
	return json.Unmarshal(data, c)
}

type Message struct {
	client  Client `json:"client"`
	message string `json:"message"`
}

func (m *Message) Serialize() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) Deserialize(data []byte) error {
	return json.Unmarshal(data, m)
}
