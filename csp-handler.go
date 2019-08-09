package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type CSPJson struct {
	CspReport struct {
		DocumentURI       string `json:"document-uri"`
		Referrer          string `json:"referrer"`
		ViolatedDirective string `json:"violated-directive"`
		OriginalPolicy    string `json:"original-policy"`
		BlockedURI        string `json:"blocked-uri"`
	} `json:"csp-report"`
}


func handleReport(c echo.Context) error {
	domain := c.Param("domain")
	var cspReportJson CSPJson

	// unmarshal json
	err := json.NewDecoder(c.Request().Body).Decode(&cspReportJson)
	if err != nil {
		log.Info(err)
		return c.String(500, "Internal server error")
	}

	// send mail in new goroutine so the http response is quicker
	go sendCSPMail(domain, cspReportJson.CspReport.DocumentURI, cspReportJson.CspReport.Referrer, cspReportJson.CspReport.ViolatedDirective,
		cspReportJson.CspReport.OriginalPolicy, cspReportJson.CspReport.BlockedURI)

	return c.NoContent(204)
}