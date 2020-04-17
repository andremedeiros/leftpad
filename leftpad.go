package main

import "strings"

type leftpad struct{}

func New() *leftpad {
	return &leftpad{}
}

func (*leftpad) Leftpad(s string, l int, padding string) string {
	if len(s) >= l {
		return s
	}

	return strings.Repeat(padding, l-len(s)) + s
}
