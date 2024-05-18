package test_test

import (
	"bytes"
	"encoding/json"
	"github.com/richieoscar/bigshop/handler"
	"github.com/richieoscar/bigshop/test/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/richieoscar/bigshop/dto"
)

func TestUserHandlerService(test *testing.T) {
	userStore := mocks.NewUserStorMock()
	handler := handler.NewHandler(userStore)

	test.Run("Should fail with invlaid payload", func(t *testing.T) {

		request := dto.RegisterRequest{
			FirstName: "Oscar",
			LastName:  "Richie",
			Email:     "hey@email.com",
			Password:  "password",
		}

		marshal, _ := json.Marshal(request)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal()
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.HandleRegister)
		if router != nil {
			router.ServeHTTP(rr, req)
		}

		//assert

		if rr.Code != http.StatusBadRequest {
			test.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

	})

}
