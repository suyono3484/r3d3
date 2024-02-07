package delivery

import (
	"encoding/json"
	"errors"
	"github.com/suyono3484/r3d3"
	"go.uber.org/mock/gomock"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)

	a := r3d3.NewMockArmament(ctrl)
	a.EXPECT().Title().Return("Turbo Laser").AnyTimes()
	a.EXPECT().Qty().Return(66).AnyTimes()

	s := r3d3.NewMockSpaceCraft(ctrl)
	s.EXPECT().ID().DoAndReturn(func() int64 { return rand.Int63n(9) + 1 }).AnyTimes()
	s.EXPECT().Name().Return("Galileo").AnyTimes()
	s.EXPECT().Class().Return("Explorer").AnyTimes()
	s.EXPECT().Crew().Return(uint64(2400)).AnyTimes()
	s.EXPECT().ImageURL().Return("https://some/file").AnyTimes()
	s.EXPECT().Value().Return(64988.00).AnyTimes()
	s.EXPECT().Status().Return("Operational").AnyTimes()
	s.EXPECT().Armament().Return([]r3d3.Armament{a, a}).AnyTimes()

	h := NewMockFilterHelper(ctrl)

	type param struct {
		method string
		id     string
	}
	type useCaseResult struct {
		spacecraft r3d3.SpaceCraft
		err        error
	}
	jsonContentTypeString := jsonContentType
	tests := []struct {
		name            string
		param           param
		result          useCaseResult
		wantStatusCode  int
		wantContentType *string
		wantJSONStruct  any
	}{
		{
			name: "positive",
			param: param{
				method: http.MethodGet,
				id:     "1",
			},
			result: useCaseResult{
				spacecraft: s,
				err:        nil,
			},
			wantStatusCode:  http.StatusOK,
			wantContentType: &jsonContentTypeString,
			wantJSONStruct:  new(spacecraftOutContainer),
		},
		{
			name: "wrong id",
			param: param{
				method: http.MethodGet,
				id:     "wrong",
			},
			result: useCaseResult{
				spacecraft: s,
				err:        nil,
			},
			wantStatusCode:  http.StatusBadRequest,
			wantContentType: nil,
			wantJSONStruct:  nil,
		},
		{
			name: "wrong method",
			param: param{
				method: http.MethodPost,
				id:     "1",
			},
			result: useCaseResult{
				spacecraft: s,
				err:        nil,
			},
			wantStatusCode:  http.StatusMethodNotAllowed,
			wantContentType: nil,
			wantJSONStruct:  nil,
		},
		{
			name: "usecase returns error",
			param: param{
				method: http.MethodGet,
				id:     "1",
			},
			result: useCaseResult{
				spacecraft: nil,
				err:        errors.New("not found"),
			},
			wantStatusCode:  http.StatusInternalServerError,
			wantContentType: nil,
			wantJSONStruct:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewMockUseCase(ctrl)
			uc.EXPECT().Get(gomock.Any()).Return(tt.result.spacecraft, tt.result.err).AnyTimes()

			deli := NewDelivery(uc, h)

			req := httptest.NewRequest(tt.param.method, "/spacecraft/"+tt.param.id, nil)
			w := httptest.NewRecorder()

			deli.ServeHTTP(w, req)
			resp := w.Result()

			if resp.StatusCode != tt.wantStatusCode {
				t.Errorf("got status code %d, expected %d", resp.StatusCode, tt.wantStatusCode)
			}

			contentType := resp.Header.Get(contentTypeHeader)
			if tt.wantContentType != nil && *tt.wantContentType != contentType {
				t.Errorf("got Content-Type %s, expected %s", contentType, *tt.wantContentType)
			}

			if tt.wantJSONStruct != nil {
				body, _ := io.ReadAll(resp.Body)
				if err := json.Unmarshal(body, tt.wantJSONStruct); err != nil {
					t.Errorf("unexpected error when parsing JSON: %+v", err)
				}
			}
		})
	}
}
