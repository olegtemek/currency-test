package model

import (
	"encoding/xml"
	"time"
)

type Rates struct {
	XMLName xml.Name   `xml:"rates"`
	Items   []Currency `xml:"item"`
}

type Currency struct {
	Id    int
	Title string  `xml:"fullname"`
	Code  string  `xml:"title"`
	Value float64 `xml:"description"`
	Date  time.Time
}
