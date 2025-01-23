package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go/github_user_activity/internal/model"
)

func GetActivity(username string) (*[]model.Activity, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	activity := []model.Activity{}
	if err := json.Unmarshal(body, &activity); err != nil {
		log.Fatal(err)
	}
	return &activity, nil
}
