package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	client := &http.Client{}

	m := map[string]string{
		"foo": "One",
		"bar": "Test",
	}

	xb, _ := json.Marshal(m)

	req, _ := http.NewRequest("POST", "http://localhost:1323/foobar", bytes.NewReader(xb))
	req.Header.Add("Authorization", `Basic Zm9vOmJhcg==`)

	resp, _ := client.Do(req)
	fmt.Println(resp)

}
