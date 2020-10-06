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
	"io/ioutil"
	"net/http"
)

type Report struct {
	UserAgent   string `json:"user_agent"`
	Destination string `json:"destination"`
	Url         string `json:"url"`
	Type        string `json:"type"`
	Timestamp   int    `json:"timestamp"`
	Attempts    int    `json:"attempts"`
	Body        []byte `json:"body"`
}

type CspViolationReport struct {
	Blocked   string `json:"blocked"`
	Directive string `json:"directive"`
	Policy    string `json:"policy"`
	Status    int    `json:"status"`
	Referrer  string `json:"referrer"`
}

// https://www.w3.org/TR/network-error-logging/
type NelReport struct {
	SamplingFraction int    `json:"sampling_fraction"`
	Referrer         string `json:"referrer"`
	ServerIp         string `json:"server_ip"`
	Protocol         string `json:"protocol"`
	Method           string `json:"method"`
	StatusCode       int    `json:"status_code"`
	ElapsedTime      int    `json:"elapsed_time"`
	Type             string `json:"type"`
}

func handleReport(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/reports+json" {
		http.Error(w, "wrong content type header", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var rep Report
	if err := json.Unmarshal(body, &r); err != nil {
		http.Error(w, "unable to parse report", http.StatusBadRequest)
		return
	}

	switch rep.Type {
	case "csp-violation":
		if err := handleCspViolation(&rep); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	case "nel":
		if err := handleNelReport(&rep); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "unsupported report type", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleCspViolation(r *Report) error {
	var cspv CspViolationReport
	if err := json.Unmarshal(r.Body, &cspv); err != nil {
		return err
	}

	return nil
}

func handleNelReport(r *Report) error {

	return nil
}