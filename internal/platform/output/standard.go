package output

import (
	"encoding/json"
	"fmt"
)

type standardOutput struct{}

// NewStandardOutput creates an standardOutput renderer instance
func NewStandardOutput() Renderer {
	return &standardOutput{}
}

// Render renders the output struct to json and print to standard output
func (r *standardOutput) Render(output interface{}) error {
	jsonOutput, err := json.Marshal(output)
	if err != nil {
		return err
	}

	fmt.Println(string(jsonOutput))

	return nil
}

// Render renders the error to json and print to standard output
func (r *standardOutput) Error(err error) error {
	output := struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	}

	jsonOutput, mErr := json.Marshal(output)
	if mErr != nil {
		return mErr
	}

	fmt.Println(string(jsonOutput))

	return nil
}
