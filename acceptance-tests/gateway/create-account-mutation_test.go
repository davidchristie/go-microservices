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

const url = "http://api.go-microservices.local/query"

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

func readResponseBody(response *http.Response) []byte {
	blob, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return blob
}

func unmarshalResponseBody(blob []byte) ResponseBody {
	body := ResponseBody{}
	if err := json.Unmarshal(blob, &body); err != nil {
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

	expectedStatusCode := 200
	expectedContentTypeHeader := []string{"application/json"}

	t.Logf("response.StatusCode = %v", response.StatusCode)
	t.Logf("response.Header = %v", response.Header)

	blob := readResponseBody(response)

	t.Logf("response.Body = %v", string(blob))

	body := unmarshalResponseBody(blob)

	// Verify status code
	actualStatusCode := response.StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("response.StatusCode = %d; expected %d", actualStatusCode, expectedStatusCode)
	}

	// Verify content type header
	actualContentTypeHeader := response.Header["Content-Type"]
	if !reflect.DeepEqual(actualContentTypeHeader, expectedContentTypeHeader) {
		t.Errorf(`response.Header["Content-Type"] = %s; expected %s`, actualContentTypeHeader, expectedContentTypeHeader)
	}

	// Verify account ID
	t.Run("HasAccountID", func(t *testing.T) {
		_, err := uuid.Parse(body.Data.CreateAccount.ID)
		if err != nil {
			t.Error("error:", err)
		}
	})
}
