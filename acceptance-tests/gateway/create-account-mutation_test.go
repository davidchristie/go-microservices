package gateway

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

const url = "http://localhost:5000/query"

type Account struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

type Data struct {
	CreateAccount Account `json:"createAccount"`
}

type ResponseBody struct {
	Data Data `json:"data"`
}

type RequestBody struct {
	OperationName string    `json:"operationName"`
	Variables     Variables `json:"variables"`
	Query         string    `json:"query"`
}

type Variables struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func sendRequest(variables Variables) *http.Response {

	requestBody := RequestBody{
		Variables: variables,
		Query:     "mutation CreateAccount($email: String!, $name: String!, $password: String!) { createAccount(input: {email: $email, name: $name, password: $password}) { id } }",
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	return response
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

func TestRequest(t *testing.T) {
	response := sendRequest(Variables{
		Email:    "test.user+" + uuid.New().String() + "@email.com",
		Name:     "Test User",
		Password: "!*pZJpiiqpC6WVT*HBM3",
	})
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
	t.Logf("response body: %s", body)

	t.Run("HasAccountID", func(t *testing.T) {
		_, err := uuid.Parse(body.Data.CreateAccount.ID)
		if err != nil {
			t.Error("error:", err)
		}
	})
}
