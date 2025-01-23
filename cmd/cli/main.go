package main

import (
	"fmt"
	"sync"

	"github.com/go/github_user_activity/internal/model"
	"github.com/go/github_user_activity/internal/usecase"
)

type ActivityResponse struct {
	Activity []model.Activity `json:"activity"`
	Err      error            `json:"error"`
}

func main() {
	var input string
	fmt.Print("github-username: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("invalid username error: ", err)
		return
	}

	chanel := make(chan ActivityResponse)
	var wg sync.WaitGroup

	fmt.Println("Fetching github activity...")
	wg.Add(1)
	go func() {
		defer wg.Done()
		activity, err := usecase.GetActivity(input)
		if err != nil {
			chanel <- ActivityResponse{Err: err}
			return
		}
		for _, act := range *activity {
			chanel <- ActivityResponse{Activity: []model.Activity{act}}
		}
	}()

	go func() {
		wg.Wait()
		close(chanel)
	}()

	for activity := range chanel {
		if activity.Err != nil {
			fmt.Println(activity.Err)
			return
		}
		act := activity.Activity[0]
		message := fmt.Sprintf("Type: %s, Actor: %s, Repo: %s, Payload: { PR: %v, Descrp: %s }",
			act.Type, act.Actor.Login, act.Repo.Name, act.Payload.PullRequest, act.Payload.Description)
		fmt.Println(message)
	}
}
