// oreste v0.1


package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alpkeskin/oreste/cmd/utils"
)

func Psbdmp(email string) {
	defer utils.ProgressBar.Add(10)
	var key string = utils.GetAPIKey("Psbdmp")
	if key == "" {
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", utils.PsbdmpAPI+email+"?key="+key, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "oreste")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &utils.Psbdmp_result)
}
