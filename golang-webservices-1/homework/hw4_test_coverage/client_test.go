package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	InvalidToken                = "invalid_token"
	TimeoutErrorQuery           = "timeout_query"
	InternalErrorQuery          = "fatal_query"
	BadRequestErrorQuery        = "bad_request_query"
	BadRequestUnknownErrorQuery = "bad_request_unknown_query"
	InvalidJSONErrorQuery       = "invalid_json_query"
)

type XMLRow struct {
	ID     int    `xml:"id"`
	Name   string `xml:"first_name"`
	Age    int    `xml:"age"`
	About  string `xml:"about"`
	Gender string `xml:"gender"`
}

type Users struct {
	Rows []XMLRow `xml:"row"`
}

func SearchServer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("AccessToken") == InvalidToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch q := r.FormValue("query"); q {
	case TimeoutErrorQuery:
		time.Sleep(time.Second * 2)
		w.WriteHeader(http.StatusFound)
		return
	case InternalErrorQuery:
		w.WriteHeader(http.StatusInternalServerError)
		return
	case InvalidJSONErrorQuery:
		w.Write([]byte("invalid_json"))
		w.WriteHeader(http.StatusOK)
		return
	case BadRequestErrorQuery:
		w.WriteHeader(http.StatusBadRequest)
		return
	case BadRequestUnknownErrorQuery:
		resp, _ := json.Marshal(SearchErrorResponse{"UnknownError"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}

	orderField := r.FormValue("order_field")
	if orderField != "Id" && orderField != "Age" && orderField != "Name" && orderField != "" {
		resp, _ := json.Marshal(SearchErrorResponse{"ErrorBadOrderField"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}

	xmlFile, err := os.Open("dataset.xml")
	if err != nil {
		fmt.Println("cant open file:", err)
		return
	}
	defer xmlFile.Close()

	var data Users
	byteData, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteData, &data)

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		fmt.Println("offset to int convertion error: ", err)
		return
	}

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		fmt.Println("limit to int convertion error: ", err)
		return
	}

	resp, err := json.Marshal(data.Rows[offset:limit])
	if err != nil {
		fmt.Println("result json packing error:", err)
		return
	}

	w.Write(resp)
}

type TestCaseValid struct {
	Request SearchRequest
}

type TestCaseInvalid struct {
	Request       SearchRequest
	URL           string
	AccessToken   string
	ErrorExact    string
	ErrorContains string
}

func TestFindUsersInvalid(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(SearchServer))
	defer server.Close()

	testCases := []TestCaseInvalid{
		{
			Request:    SearchRequest{Limit: -1},
			ErrorExact: "limit must be > 0",
		},
		{
			Request:    SearchRequest{Offset: -1},
			ErrorExact: "offset must be > 0",
		},
		{
			Request:       SearchRequest{Limit: 1},
			URL:           "http://",
			ErrorContains: "unknown error",
		},
		{
			Request:       SearchRequest{Query: TimeoutErrorQuery},
			ErrorContains: "timeout for",
		},
		{
			AccessToken: InvalidToken,
			ErrorExact:  "Bad AccessToken",
		},
		{
			Request:    SearchRequest{Query: InternalErrorQuery},
			ErrorExact: "SearchServer fatal error",
		},
		{
			Request:       SearchRequest{Query: BadRequestErrorQuery},
			ErrorContains: "cant unpack error json",
		},
		{
			Request:       SearchRequest{Query: BadRequestUnknownErrorQuery},
			ErrorContains: "unknown bad request error",
		},
		{
			Request:    SearchRequest{OrderField: "order_field"},
			ErrorExact: "OrderFeld order_field invalid",
		},
		{
			Request:       SearchRequest{Query: InvalidJSONErrorQuery},
			ErrorContains: "cant unpack result json",
		},
	}

	for n, testCase := range testCases {
		url := server.URL
		if testCase.URL != "" {
			url = testCase.URL
		}

		client := SearchClient{
			URL:         url,
			AccessToken: testCase.AccessToken,
		}
		response, err := client.FindUsers(testCase.Request)

		if err == nil || response != nil {
			t.Errorf("[%d] expected error, got nil", n)
		}

		if testCase.ErrorExact != "" && err.Error() != testCase.ErrorExact {
			t.Errorf("[%d] wrong result, expected %#v, got %#v", n, testCase.ErrorExact, err.Error())
		}

		if testCase.ErrorContains != "" && !strings.Contains(err.Error(), testCase.ErrorContains) {
			t.Errorf("[%d] wrong result, expected %#v to contain %#v", n, err.Error(), testCase.ErrorContains)
		}
	}
}

func TestFindUsersValid(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(SearchServer))
	defer server.Close()

	testCases := []TestCaseValid{
		{
			SearchRequest{Limit: 1},
		},
		{
			SearchRequest{Limit: 30},
		},
		{
			SearchRequest{Limit: 25, Offset: 1},
		},
	}

	for n, testCase := range testCases {
		client := SearchClient{
			URL: server.URL,
		}
		response, err := client.FindUsers(testCase.Request)

		if err != nil || response == nil {
			t.Errorf("[%d] expected response, got error", n)
		}
	}
}
