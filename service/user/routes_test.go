package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/richieoscar/bigshop/service/user"
	"github.com/richieoscar/bigshop/service/user/mocks"
	"github.com/richieoscar/bigshop/types"
)

func TestUserHandlerService(test *testing.T) {
	userStore := mocks.NewUserStorMock()
	handler := user.NewHandler(userStore)

	test.Run("Should fail with invlaid payload", func(t *testing.T) {

		request := types.RegisterRequest{
			FirsName: "Oscar",
			LastName: "Richie",
			Email:    "hey@email.com",
			Password: "password",
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
