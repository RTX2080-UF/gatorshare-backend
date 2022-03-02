package main

import ("testing"
	"net/http/httptest"
	"gatorshare/middleware"
	"gatorshare/models"
)


func TestCreatePost(t *testing.T,base *Controller) {
  	// if AddNewpost(base.DB,  &posts) != 1  {
	// t.Error("Test for Create Post has failed!")
  	// }
  
	var jsonStr = []byte(`{
		"userId" : 6,
		"title" : "Test post 1",
		"description" : "Test Message",
		"userLimit" : 4,
		"status" : 2
	}`)

	req, err := http.NewRequest("POST", "/posts/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddNewpost)
	handler.ServeHTTP(rr, req)																																									
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// expected := `{"id":4,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }

}
