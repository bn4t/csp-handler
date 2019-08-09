/*
 *     Copyright (C) 2019  bn4t
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"os"
	"strconv"
	"strings"
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

var RateLimitMap = make(map[string]int)

func handleReport(c echo.Context) error {
	domain := c.Param("domain")
	var cspReportJson CSPJson
	ip := c.RealIP()

	// Validate the that requester is still in bounds of ratelimit
	// if not deny the request
	// This is to prevent spam since every CSP report sends a mail

	// check if an entry for the ip already exists and add 1 request for the ip
	// if it doesn't exist add it to the map
	if _, hasValue := RateLimitMap[ip]; hasValue {

		rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
		if err != nil {
			println("Invalid RATE_LIMIT value.")
			return c.String(500, "Internal server error")
		}

		// check if ratelimit is exceeded
		if RateLimitMap[ip] < rateLimit {
			RateLimitMap[ip]++
		} else {
			return c.String(492, "Too many requests")
		}

	} else {
		RateLimitMap[ip] = 1
	}

	// deny request if the content-type header is wrong
	if !strings.Contains(c.Request().Header.Get("Content-Type"), "application/csp-report") {
		return c.String(400, "Bad Request")
	}

	// unmarshal json
	err := json.NewDecoder(c.Request().Body).Decode(&cspReportJson)
	if err != nil {
		return c.String(500, "Internal server error")
	}

	// send mail in new goroutine so the http response is quicker
	go sendCSPMail(domain, cspReportJson.CspReport.DocumentURI, cspReportJson.CspReport.Referrer, cspReportJson.CspReport.ViolatedDirective,
		cspReportJson.CspReport.OriginalPolicy, cspReportJson.CspReport.BlockedURI)

	return c.NoContent(204)
}
