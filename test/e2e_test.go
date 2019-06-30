package test

import (
	"os"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

var url = os.Getenv("WS_ENDPOINT")

func TestConnectionCanBeEstablished(t *testing.T) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	defer c.Close()
	assert.Nil(t, err)
	assert.NotNil(t, c)
}

func TestTwoEchoMessagesAreReceived(t *testing.T) {
	c, _, _ := websocket.DefaultDialer.Dial(url, nil)
	defer c.Close()
	message := []byte("foo")

	c.WriteMessage(websocket.TextMessage, message)

	_, firstMessage, _ := c.ReadMessage()
	_, secondMessage, _ := c.ReadMessage()
	assert.Equal(t, message, firstMessage)
	assert.Equal(t, message, secondMessage)
}
