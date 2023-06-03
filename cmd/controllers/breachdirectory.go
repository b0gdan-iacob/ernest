// oreste v0.1


package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/b0gdan-iacob/oreste/cmd/utils"
)

func BreachDirectory(email string) {
	defer utils.ProgressBar.Add(10)
	var key string = utils.GetAPIKey("Breachdirectory")
	if key == "" {
		return
	}
	url := utils.BreachDirectoryAPI + email

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("x-rapidapi-host", "breachdirectory.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &utils.Breachdirectory_result)
}
