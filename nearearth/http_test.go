package nearearth_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/smccutcheon-whs/solidinpractice/nearearth"
)

var (
	countHttpTestCases = []struct {
		name          string
		endpoint      string
		handler       http.HandlerFunc
		err           error
		expectedCode  int
		expectedCount int
	}{
		{
			name:         "no impacts",
			endpoint:     "/noimpacts",
			handler:      testHandler(nearearth.Impacts{}, nil),
			expectedCode: http.StatusOK,
		},
		{
			name:     "one impact",
			endpoint: "/oneimpact",
			handler: testHandler(nearearth.Impacts{
				{
					Designation:   "2021 AB",
					DiscoveryDate: "2021-01-01",
				},
			}, nil),
			expectedCode:  http.StatusOK,
			expectedCount: 1,
		},
		{
			name:     "two impacts",
			endpoint: "/twoimpacts",
			handler: testHandler(nearearth.Impacts{
				{
					Designation:   "2021 AB",
					DiscoveryDate: "2021-01-01",
				},
				{
					Designation:   "2021 AC",
					DiscoveryDate: "2021-01-02",
				},
			}, nil),
			expectedCode:  http.StatusOK,
			expectedCount: 2,
		},
		{
			name:         "error",
			endpoint:     "/error",
			handler:      testHandler(nearearth.Impacts{}, fmt.Errorf("expected error")),
			err:          fmt.Errorf("expected error"),
			expectedCode: http.StatusInternalServerError,
		},
	}
)

func testHandler(i nearearth.Impacts, err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
}

func newTestServer(h map[string]http.HandlerFunc) *testServer {
	s := &testServer{
		handlers: h,
	}
	s.start()
	return s
}

type testServer struct {
	server   *http.Server
	mux      *http.ServeMux
	handlers map[string]http.HandlerFunc
}

func (s *testServer) start() {
	s.mux = http.NewServeMux()
	for k, v := range s.handlers {
		s.mux.HandleFunc(k, v)
	}
	s.server = &http.Server{
		Addr:    ":8080",
		Handler: s.mux,
	}
	go s.server.ListenAndServe()
}

func (s *testServer) close() {
	s.server.Close()
}

func Test_HTTPCount(t *testing.T) {
	handlers := map[string]http.HandlerFunc{}
	for _, tc := range countHttpTestCases {
		handlers[tc.endpoint] = tc.handler
	}
	s := newTestServer(handlers)
	s.start()
	for _, tc := range countHttpTestCases {
		defer s.server.Close()
		t.Run(tc.name, func(t *testing.T) {
			svc := nearearth.NewImpactHTTPService(nearearth.WithTestURL("http://localhost:8080" + tc.endpoint))
			count, err := nearearth.Count(svc)
			if err != nil {
				if tc.err == nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
			if count != tc.expectedCount {
				t.Fatalf("expected count %d, got %d", tc.expectedCount, count)
			}
		})
	}
}
