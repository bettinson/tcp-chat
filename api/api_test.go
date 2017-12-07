package api

import "testing"

func TestClientSerializeDeserialize(t *testing.T) {
	original := &Client{Name: "Matt", Channel: "Yeezys"}
	data, err := original.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	result := &Client{}
	err = result.Deserialize(data)
	if err != nil {
		t.Fatal(err)
	}

	if *original != *result {
		t.Fatalf("%v != %v", original, result)
	}
}

func TestMessageSerializeDeserialize(t *testing.T) {
	client := &Client{Name: "Matt", Channel: "Yeezys"}
	message := &Message{Client: client, message: "Yeezys"}
	data, err := message.Serialize()
	if err != nil {
		t.Fatal(err)
	}

	result := &Message{}
	err = result.Deserialize(data)
	if err != nil {
		t.Fatal(err)
	}

	if *message != *result {
		t.Fatalf("%v != %v", message, result)
	}
}
