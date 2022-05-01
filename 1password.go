package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

type opResponse struct {
	UUID    string            `json:"uuid"`
	Details opResponseDetails `json:"details"`
}
type opResponseDetails struct {
	Fields []opResponseField `json:"fields"`
	Title  string            `json:"title"`
}

type opResponseField struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Designation string `json:"designation"`
}

var defaultOpGet = func(itemName string) (*opResponse, error) {
	out, err := exec.Command("op", "item", "get", itemName, "--fields", "credential", "--format", "json").CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", out)
		return nil, err
	}
	resp := opResponse{Details: opResponseDetails{Fields: []opResponseField{opResponseField{}}}}
	err = json.Unmarshal(out, &resp.Details.Fields[0])
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func opgetter(itemName string) (string, error) {
	resp, err := defaultOpGet(itemName)
	if err != nil {
		return "", err
	}
	for _, v := range resp.Details.Fields {
		return v.Value, nil
	}
	return "", errors.New("unable to find password")
}

func opsetter(itemName, secret string) error {
	stdoutStderr, err := exec.Command("op", "item", "create", `--category=API Credential`,
		"credential="+secret, "--title="+itemName).CombinedOutput()

	fmt.Printf("%s\n", stdoutStderr)
	return err
}
