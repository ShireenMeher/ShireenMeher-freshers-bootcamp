//package main
//
//import (
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"userproject/Models"
//)
//func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
//	req, _ := http.NewRequest(method, path, nil)
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	return w
//}
//func TestHelloWorld(t *testing.T) {
//	// Build our expected body
//	body := gin.H{
//		"id" : "125132",
//		"name" : "shireen",
//		"email" : "shireen@gmail.com",
//		"phone" : "12345678",
//		"address" : "ACBD",
//	}
//	// Grab our router
//	router := Routes.SetupRouter()
//	path:= "/user-api/user/125"
//	// Perform a GET request with that handler.
//	w := performRequest(router, "GET", path)
//	// Assert we encoded correctly,
//	// the request gives a 200
//	assert.Equal(t, http.StatusOK, w.Code)
//	// Convert the JSON response to a map
//	var response map[string]string
//	err := json.Unmarshal([]byte(w.Body.String()), &response)
//	// Grab the value & whether or not it exists
//	//temp:= string()
//	value1, exists1 := response["id"]
//	assert.Nil(t, err)
//	assert.True(t, exists1)
//	assert.Equal(t, body["id"], value1)
//
//	value2, exists2 := response["name"]
//	assert.True(t, exists2)
//	assert.Equal(t, body["id"], value2)
//
//	value3, exists3 := response["email"]
//	assert.True(t, exists3)
//	assert.Equal(t, body["id"], value3)
//	value4, exists4 := response["phone"]
//	assert.True(t, exists4)
//	assert.Equal(t, body["id"], value4)
//	value5, exists5 := response["address"]
//	assert.True(t, exists5)
//	assert.Equal(t, body["id"], value5)
//
//	// Make some assertions on the correctness of the response.
//
//
//}
//
//func TestGetEntryByID(t *testing.T) {
//	req, err := http.NewRequest("GET", "/entry", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	q := req.URL.Query()
//	q.Add("id", "1")
//	req.URL.RawQuery = q.Encode()
//	rr := httptest.NewRecorder()
//	handler := http.HandlerFunc("user-api/user/12532", GetUserByID)
//	handler.ServeHTTP(rr, req)
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//	// Check the response body is what we expect.
//	expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"}`
//	if rr.Body.String() != expected {
//		t.Errorf("handler returned unexpected body: got %v want %v",
//			rr.Body.String(), expected)
//	}
//}

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"userproject/Config"
	"userproject/Controllers"
	"userproject/Models"
	"userproject/Routes"

	"github.com/jinzhu/gorm"
)


func TestGet(t *testing.T) {
	//SQL Connection using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//Setting the router
	router := Routes.SetupRouter()
	router.GET("/user-api/user/", Controllers.GetUsers)

	//Get request
	request, _ := http.NewRequest("GET", "/user-api/user/", nil)

	//Recording the response
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput:= `[{"id":22,"name":"Test","email":"test@gmail.com","phone":"12345","address":"location"}]`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}

func TestPost(t *testing.T) {
	//SQL Connection using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.User{})


	//Setting the router
	router := Routes.SetupRouter()
	router.POST("/user-api/user/",Controllers.CreateUser)

	//send request
	newStud := Models.User{
		Id: 21,
		Name: "Test",
		Email: "test@gmail.com",
		Phone: "12345",
		Address: "location",

	}

	responseBody,_ := json.Marshal(newStud)
	req, _ := http.NewRequest("POST", "/user-api/user/", bytes.NewBuffer([]byte(responseBody)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedOutput := `{"id":21,"name":"Test","email":"test@gmail.com","phone":"12345","address":"location"}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}

}

func TestPut(t *testing.T) {
	//SQL database using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.User{})

	//setting up router
	router := Routes.SetupRouter()
	router.PUT("/user-api/user/22",Controllers.UpdateUser)

	//send request
	newStudent := Models.User{
		Name:  "test2",
		Address: "Location2",

	}

	responseBody,_ := json.Marshal(newStudent)
	req, _ := http.NewRequest("PUT", "/user-api/user/22", bytes.NewBuffer([]byte(responseBody)))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	expectedOutput := `{"id":22,"name":"test2","email":"test@gmail.com","phone":"12345","address":"Location2"}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expectedOutput)
	}
}

func TestDelete(t *testing.T) {
	//SQL Connection using GORM
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//setup router
	router := Routes.SetupRouter()
	router.DELETE("/user-api/user/22",Controllers.DeleteUser )

	//Get request
	req, _ := http.NewRequest("DELETE", "/user-api/user/22", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	//checking test case
	expectedOutput := `{"id22":"is deleted"}`
	if response.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Code, expectedOutput)
	}
}
