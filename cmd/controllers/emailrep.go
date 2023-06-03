// oreste v0.1


package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/b0gdan-iacob/oreste/cmd/utils"
)

func EmailRep(email string) {
	defer utils.ProgressBar.Add(10)
	var key string = utils.GetAPIKey("Emailrep")
	if key == "" {
		return
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", utils.EmailrepURL+email, nil)
	req.Header.Set("Key", key)
	req.Header.Set("User-Agent", "oreste")
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &utils.Emailrep_result)
}
