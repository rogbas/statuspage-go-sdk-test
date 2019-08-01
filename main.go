package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rogbas/statuspage-go-sdk-test/config"
	sp "github.com/statusio/statusio-go"
)

func main() {

	spAPI := sp.NewStatusioApi(config.OrgId, config.ApiKey)

	// statusSummary, err := spAPI.StatusSummary(config.PageId)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	incList, err := spAPI.IncidentList(config.PageId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Incidents: %#v", incList.Result.ActiveIncidents)

	os.Exit(0)
}
