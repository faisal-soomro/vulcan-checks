/*
Copyright 2023 Adevinta
*/

package main

import (
	"context"
	check "github.com/adevinta/vulcan-check-sdk"
	"github.com/adevinta/vulcan-check-sdk/state"
	report "github.com/adevinta/vulcan-report"
)

var (
	checkName = "vulcan-gcp-test"
	logger    = check.NewCheckLog(checkName)
)

func main() {
	c := check.NewCheckFromHandler(checkName, run)
	c.RunAndServe()
}

func run(ctx context.Context, target, assetType, optJSON string, state state.State) (err error) {
	logger.Printf("Starting the %v check for %v asset type", checkName, assetType)
	logger.Printf("Validating params. Target: %v Options: %v ...", target, optJSON)
	logger.Printf("Adding dummy vulnerability")

	vuln := report.Vulnerability{
		Summary:                "GCPProject Vulnerability",
		Description:            "Dummy vulnerability to test GCPProject asset type",
		Score:                  0.0,
		AffectedResource:       "dummyResource",
		AffectedResourceString: "Dummy Resource String",
		Labels:                 []string{"gcp"},
	}
	state.AddVulnerabilities(vuln)

	return ctx.Err()
}
