// Copyright (C) 2019-2023 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package lib

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/algorand/go-algorand/logging"
	"github.com/algorand/go-algorand/node"
)

// GenesisJSONText is initialized when the node starts.
var GenesisJSONText string

// HandlerFunc defines a wrapper for http.HandlerFunc that includes a context
type HandlerFunc func(ReqContext, echo.Context)

// Route type description
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc HandlerFunc
}

// Routes contains all routes
type Routes []Route

// ReqContext is passed to each of the handlers below via wrapCtx, allowing
// handlers to interact with the node
type ReqContext struct {
	Node     *node.AlgorandFullNode
	Log      logging.Logger
	Context  echo.Context
	Shutdown <-chan struct{}
}

// ErrorResponse sets the specified status code (should != 200), and fills in the
// a human readable error.
func ErrorResponse(w http.ResponseWriter, status int, internalErr error, publicErr string, logger logging.Logger) {
	logger.Info(internalErr)

	w.WriteHeader(status)
	_, err := w.Write([]byte(publicErr))
	if err != nil {
		logger.Errorf("algod failed to write response: %v", err)
	}
}
