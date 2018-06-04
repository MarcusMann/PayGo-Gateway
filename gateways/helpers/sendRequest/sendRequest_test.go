package sendRequest

import (
	"encoding/json"
	"testing"
)

type TestResquest struct {
	ID     int    `json:"id, omitempty"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int64  `json:"userId"`
	Response map[string]interface{} `json:"response, omitempty"`
}

func (t TestResquest) GetUrl() (string, error) {
	return "https://jsonplaceholder.typicode.com/posts/", nil
}

func (t TestResquest) GetJSON() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (t *TestResquest) SetResponse(b []byte) error {
	err := json.Unmarshal(b, &t.Response)
	if err != nil {
		return err
	}
	return nil
}

func TestSendRequest(t *testing.T) {
	tr := new (TestResquest)
	tr.Title = "foo"
	tr.Body = "bar"
	tr.UserId = 1

	err := SendJSONRequest(tr, "POST")
	if err != nil {
		t.Errorf("error on send Request: %s\n", err)
	}
}
