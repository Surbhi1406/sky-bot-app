package Controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stanwar/sky-bot/Models"

	_ "github.com/mattn/go-sqlite3"
)

func sendRequestToSkybot(skyBotEngnURL, webhookURL string, task Models.Task) {
	fmt.Println("Started  sendDataToSkybot ")
	dataBytes, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Error marshalling data:", err)

	}

	resp, err := http.Post(skyBotEngnURL+"?webhook_url="+webhookURL, "application/json", bytes.NewBuffer(dataBytes))
	if err != nil {
		fmt.Println("Error sending data to webhook:", err)

	}

	defer resp.Body.Close()

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding webhook response:", err)

	}

	fmt.Println("Webhook response:", response)
}
