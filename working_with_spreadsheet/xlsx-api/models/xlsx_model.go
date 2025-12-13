package models

type XLSXRequest struct {
	Header []string `json:"header"`
	Criteria []string `json:"criteria"`
	Students [][]any `json:"students"`
}