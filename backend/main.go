// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// HealthCheck just returns true if the service is up.
func HealthCheck(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("🚑 healthcheck ok!")
	w.WriteHeader(http.StatusOK)
}

// SimpleGet just returns true no matter what
func SimpleGet(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("Get Succeeded")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("back-end"))
}

// ServiceDiscoveryGet just returns true no matter what
func ServiceDiscoveryGet(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("Get on ServiceDiscovery endpoint successful!!! Or does it???")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("back-end-service-discovery: %v", time.Now())))
}

func main() {
	router := httprouter.New()
	router.GET("/back-end/", SimpleGet)
	router.GET("/service-discovery/", ServiceDiscoveryGet)

	// Health Check
	router.GET("/", HealthCheck)

	log.Fatal(http.ListenAndServe(":80", router))
}
