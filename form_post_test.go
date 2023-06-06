package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFormPost(t *testing.T) {

	requestBody := strings.NewReader("firstName=Muhammad&lastName=Rivaldo")
	request := httptest.NewRequest("POST", "http://localhost/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()
	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstNameValue := r.PostFormValue("firstName")

	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")
	fmt.Fprintf(w, "Hello %s %s %s", firstName, lastName, firstNameValue)
}
