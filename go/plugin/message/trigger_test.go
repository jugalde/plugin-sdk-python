package message

import (
	"encoding/json"
	"testing"

	"github.com/orcalabs/plugin-sdk/go/plugin/connect"
)

func TestMarshalTriggerStartWithMessageEnvelope(t *testing.T) {

	expected := `{"version":"v1","type":"trigger_start","body":{"trigger_id":1,"connection":{"username":"hello","password":"blah"},"dispatcher_url":"http://abc123","trigger":"hello","input":{"param":"ok","param2":"blah"}}}`

	connJSON := `
	{
		"username": "hello",
		"password": "blah"
	}
	`
	msg := json.RawMessage([]byte(connJSON))
	conn := connect.Connection{
		RawMessage: msg,
		Contents:   nil,
	}

	inputJSON := `
	{
		"param": "ok",
		"param2": "blah"
	}
	`
	msg = json.RawMessage([]byte(inputJSON))
	input := Input{
		msg,
		nil,
	}

	trig := &TriggerStart{
		Connection:    conn,
		TriggerID:     1,
		DispatcherURL: "http://abc123",
		Trigger:       "hello",
		Input:         input,
	}

	m := Message{
		Header: Header{
			Version: "v1",
			Type:    "trigger_start",
		},
	}

	str, err := m.Marshal(trig)

	if err != nil {
		t.Fatal("Unable to marshal: ", err)
	}

	if string(str) != expected {
		t.Fatalf("Got %s but expected %s", str, expected)
	}

}

func TestMarshalTriggerStart(t *testing.T) {
	expected := `{"trigger_id":1,"connection":{"username":"hello","password":"blah"},"dispatcher_url":"http://abc123","trigger":"hello","input":{"param":"ok","param2":"blah"}}`

	connJSON := `
	{
		"username": "hello",
		"password": "blah"
	}
	`
	msg := json.RawMessage([]byte(connJSON))
	conn := connect.Connection{
		RawMessage: msg,
		Contents:   nil,
	}

	inputJSON := `
	{
		"param": "ok",
		"param2": "blah"
	}
	`
	msg = json.RawMessage([]byte(inputJSON))
	input := Input{
		msg,
		nil,
	}

	trig := &TriggerStart{
		Connection:    conn,
		TriggerID:     1,
		DispatcherURL: "http://abc123",
		Trigger:       "hello",
		Input:         input,
	}
	str, err := json.Marshal(trig)

	if err != nil {
		t.Fatal("Unable to marshal: ", err)
	}

	if string(str) != expected {
		t.Fatalf("Got %s but expected %s", str, expected)
	}
}
