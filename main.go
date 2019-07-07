/*
 * AwesomeMusic API
 *
 * 音楽サービス AwesomeMusicのWebAPI
 *
 * API version: 0.0.1
 * Contact: adf1985adf@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "./go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
