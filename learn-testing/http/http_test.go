package http

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Figure out two things
// What are inputs (parameters)
// Expected output

func TestDoubleHandler(t *testing.T) {
	tt := [...]struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "OK",
			queryParam:     "10",
			expectedStatus: http.StatusOK,
			expectedBody:   "20",
		},
		{
			name:           "Fail_MissingValue",
			queryParam:     "", // Missing `v` parameter
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "missing value",
		},

		{
			name:           "Fail_NotANumber",
			queryParam:     "abc", // `v` is not a number
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "not a number: abc",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			// NewRecorder returns an initialized [ResponseRecorder].
			// ResponseRecorder is an implementation of [http.ResponseWriter]
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/double?v="+tc.queryParam, nil)
			doubleHandler(w, r)

			require.Equal(t, tc.expectedStatus, w.Code)
			require.Equal(t, tc.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
