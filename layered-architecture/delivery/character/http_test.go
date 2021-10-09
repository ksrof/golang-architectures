// layered-architecture/delivery/character/http_test.go

package character

import (
	"bytes"
	"errors"
	"layered-architecture/entities"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestCharacterHandler_Handler(t *testing.T) {
	testcases := []struct{
		method 							string
		expectedStatusCode 	int
	}{
		{"GET", http.StatusOK},
		{"POST", http.StatusOK},
		{"DELETE", http.StatusMethodNotAllowed},
	}

	for _, v := range testcases {
		req := httptest.NewRequest(v.method, "/character", nil)
		w := httptest.NewRecorder()

		c := New(mockDatastore{})
		c.Handler(w, req)

		if w.Code != v.expectedStatusCode {
			t.Errorf("Expected %v\tGot %v", v.expectedStatusCode, w.Code)
		}
	}
}

func TestCharacterGet(t *testing.T) {
	testcases := []struct{
		id string
		response []byte
	}{
		{"1", []byte("could not retrieve character")},
		{"1a", []byte("invalid parameter id")},
		{"2", []byte(`[{"ID": 2, "Name": "Sylvanas Windrunner", "Level": "??(Boss)", "Gender": "Female", "Race": "Forsaken", "Class": "Ranger", "Faction": "Horde"}]`)},
		{"0", []byte(`[{"ID": 1, "Name": "Leeroy Jenkins", "Level": 70", "Gender": "Male", "Race": "Humanoid", "Class": "Paladin", "Faction": "Alliance"}, {"ID": 2, "Name": "Sylvanas Windrunner", "Level": "??(Boss)", "Gender": "Female", "Race": "Forsaken", "Class": "Ranger", "Faction": "Horde"}]`)},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/character?id="+v.id, nil)
		w := httptest.NewRecorder()

		c := New(mockDatastore{})

		c.Get(w, req)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.response)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.response))
		}
	}
}

func TestCharacterCreate(t *testing.T) {
	testcases := []struct{
		reqBody []byte
		resBody []byte
	}{
		{[]byte(`{"Name": "Sylvanas Windrunner", "Level": "??(Boss)", "Gender": "Female", "Race": "Forsaken", "Class": "Ranger", "Faction": "Horde"}`), []byte(`could not create character`)},
		{[]byte(`{"Name": "Leeroy Jenkins", "Level": "70", "Gender": "Male", "Race": "Humanoid", "Class": "Paladin", "Faction": "Alliance"}`), []byte(`{"ID": 12, "Name": "Leeroy Jenkins", "Level": "70", "Gender": "Male", "Race": "Humanoid", "Class": "Paladin", "Faction": "Alliance"}`)},
		{[]byte(`{"Name": "Leeroy Jenkins", "Level": "70", "Gender": "Male", "Race": "Humanoid", "Class": "Paladin", "Faction": "Alliance"}`), []byte(`invalid body`)},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("POST", "/character", bytes.NewReader(v.reqBody))
		w := httptest.NewRecorder()

		c := New(mockDatastore{})

		c.Create(w, req)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.resBody)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.resBody))
		}
	}
}

type mockDatastore struct{}

func (m mockDatastore) Get(id int) ([]entities.Character, error) {
	if id == 1 {
		return nil, errors.New("db error")
	} else if id == 2 {
		return []entities.Character{{2, "Sylvanas Windrunner", "??(Boss)", "Female", "Forsaken", "Ranger", "Horde"}}, nil
	}

	return []entities.Character{{1, "Leeroy Jenkins", "70", "Male", "Humanoid", "Paladin", "Alliance"}, {2, "Sylvanas Windrunner", "??(Boss)", "Female", "Forsaken", "Ranger", "Horde"}}, nil
}

func (m mockDatastore) Create(character entities.Character) (entities.Character, error) {
	if strings.Compare(character.Name, "Sylvanas Windrunner") == 1 {
		return entities.Character{}, errors.New("db error")
	}

	return entities.Character{12, "Leeroy Jenkins", "70", "Male", "Humanoid", "Paladin", "Alliance"}, nil
}