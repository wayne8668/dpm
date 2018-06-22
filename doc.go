// Package classification DPM API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:8080
//     BasePath: /dpm/api/v1.0
//     Version: v1.0
//     Contact: xhtian<kook1001@126.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - jwt:
//
//     SecurityDefinitions:
//     jwt:
//          type: apiKey
//          name: Authorization
//          in: header
//
//
// swagger:meta
package main

//https://gowalker.org/github.com/go-validator/validator