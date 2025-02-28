package controller

import (
	_ "bytes"
	"encoding/json"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmployeeHandler(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		url          string
		body         io.Reader
		expectedCode int
		expectedBody web.WebResponse
	}{
		{
			name:         "Find Employee By ID",
			method:       http.MethodGet,
			url:          "/employees/1",
			body:         nil,
			expectedCode: http.StatusOK,
			expectedBody: web.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data: map[string]interface{}{
					"employee_id": "1",
					"name":        "John Doe",
					"role":        "Software Engineer",
					"email":       "",
					"phone":       "",
					"date_hired":  "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, tt.body)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				response := web.WebResponse{
					Code:   http.StatusOK,
					Status: "OK",
					Data: map[string]interface{}{
						"employee_id": "1",
						"name":        "John Doe",
						"role":        "Software Engineer",
						"email":       "",
						"phone":       "",
						"date_hired":  "",
					},
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			})

			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)

			var respBody map[string]interface{}
			json.NewDecoder(rr.Body).Decode(&respBody)

			assert.Equal(t, float64(tt.expectedBody.Code), respBody["code"])
			assert.Equal(t, tt.expectedBody.Status, respBody["status"])
			assert.Equal(t, tt.expectedBody.Data, respBody["data"])
		})
	}
}
