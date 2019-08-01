package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rogbas/statuspage-go-sdk-test/config"
	sp "github.com/statusio/statusio-go"
)

func main() {

	spAPI := sp.NewStatusioApi(config.ApiId, config.ApiKey)

	statusSummary, err := spAPI.StatusSummary(config.PageId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(statusSummary.Result.Status[0].Containers[0].Status)

	os.Exit(0)
}
