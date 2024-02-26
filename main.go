package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/go-playground/webhooks/v6/github"
)
const (
	path = "/webhooks"
)

func main() {
	fmt.Printf("Starting webhook server based on commit: %s", os.Getenv("WEBHOOK_COMMIT_SHA"))
	
	
	hook, _ := github.New(github.Options.Secret(os.Getenv("WEBHOOK_SECRET")))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn't one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
	http.ListenAndServe(":3000", nil)
}