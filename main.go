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
	"flag"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	flag.StringVar(&Config.ConfigPath, "config", "./config.toml", "The config file for csp-handler.")
	flag.Parse()

	readConfig()

	// start the loop to reset the rate limit
	go rateLimitLoop()

	e := echo.New()
	e.POST("/report-uri/:domain", handleReport)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is a csp report handler. See https://git.bn4t.me/bn4t/csp-handler for more info.")
	})

	e.Logger.Fatal(e.Start(Config.BindTo))
}
