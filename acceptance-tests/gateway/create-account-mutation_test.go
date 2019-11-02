package gateway

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

const url = "http://localhost:5000/query"
const requestBody = `
	{
		"operationName": null,
		"variables": {
			"email": "test@email.com",
			"name": "Test Account",
			"password": "test_password123"
		},
		"query": "mutation CreateAccount($email: String!, $name: String!, $password: String!) { createAccount(input: {email: $email, name: $name, password: $password}) { id } }"
	}
`

type Account struct {
	ID string
}

type Data struct {
	CreateAccount Account
}

type ResponseBody struct {
	Data Data
}

func unmarshalResponseBody(response *http.Response) ResponseBody {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	body := ResponseBody{}
	err = json.Unmarshal(blob, &body)
	if err != nil {
		panic(err)
	}
	return body
}

func sendRequest() *http.Response {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(requestBody)))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	return response
}

func TestRequest(t *testing.T) {
	response := sendRequest()
	defer response.Body.Close()

	t.Run("StatusCode", func(t *testing.T) {
		const expected = 200
		actual := response.StatusCode
		if actual != expected {
			t.Errorf(`response.StatusCode = %d; expected %d`, actual, expected)
		}
	})

	t.Run("ContentTypeHeader", func(t *testing.T) {
		expected := []string{"application/json"}
		actual := response.Header["Content-Type"]
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf(`response.StatusCode = %s; expected %s`, actual, expected)
		}
	})

	body := unmarshalResponseBody(response)

	t.Run("HasAccountID", func(t *testing.T) {
		_, err := uuid.Parse(body.Data.CreateAccount.ID)
		if err != nil {
			t.Error("error:", err)
		}
	})
}
