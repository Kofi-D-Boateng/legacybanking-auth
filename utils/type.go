package utils

import "encoding/json"

type Response struct {
	StatusCode int         		`json:"status"`
	Body       json.RawMessage 	`json:"body,omitempty"`
}

type Request struct {
	Function string 			`json:"function"`
	Payload  json.RawMessage 	`json:"payload"`
}