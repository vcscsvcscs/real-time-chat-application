package hub

type receivedMessage struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

type chatClient struct {
	Name string `json:"name"`
}
