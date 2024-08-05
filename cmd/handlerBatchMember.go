package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/cmd/domain"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Response 구조체 정의
type Response struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Timestamp time.Time       `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}

func ApiGetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Input Data check Query
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Missing 'id' parameter")
		return
	}
	// API 호출
	host := "http://localhost:3000/"
	uri := "api/v1/users/" + id

	responseData, err := http.Get(host + uri)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error calling external API:%s", err)
		return
	}
	defer responseData.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(responseData.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error reading response body:%s", err)
		return
	}
	// 응답에서 data 영역만 추출해서 return
	// JSON을 Response 구조체로 언마샬링
	var response Response
	err2 := json.Unmarshal((body), &response)
	if err2 != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// Data 필드의 실제 구조에 따라 언마샬링을 진행
	var dataMember *domain.TMemberTemp
	err4 := json.Unmarshal(response.Data, &dataMember)
	if err4 != nil {
		fmt.Println("Error decoding JSON:", err4)
		return
	}

	w.WriteHeader(http.StatusOK)
	if dataMember.UserId == "" {
		json.NewEncoder(w).Encode(nil)
	} else {
		// response data를 temp 테이블에 삽입
		insertMemberTemp((*domain.TMemberTemp)(dataMember))

		json.NewEncoder(w).Encode(dataMember)
	}
}

func ApiGetUserAllHandler(w http.ResponseWriter, r *http.Request) {
	// API 호출
	host := "http://localhost:3000/"
	uri := "api/v1/users"

	responseData, err := http.Get(host + uri)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error calling external API:%s", err)
		return
	}
	defer responseData.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(responseData.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error reading response body:%v", err)
		return
	}
	log.Printf("body : %s", body)
	// 응답에서 data 영역만 추출해서 return
	// JSON을 Response 구조체로 언마샬링
	var response Response
	err2 := json.Unmarshal(body, &response)
	if err2 != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// Data 필드의 실제 구조에 따라 언마샬링을 진행
	var dataMember *[]domain.TMember
	err4 := json.Unmarshal(response.Data, &dataMember)
	if err4 != nil {
		fmt.Println("Error decoding JSON:", err4)
		return
	}

	w.WriteHeader(http.StatusOK)
	if len(*dataMember) == 0 {
		json.NewEncoder(w).Encode(nil)
	} else {
		// response data를 temp 테이블에 삽입
		cnt := insertAllMemberTemp(*dataMember)
		fmt.Fprintf(w, "Batch Row successed: %d 건", cnt)
		// json.NewEncoder(w).Encode(cnt)
	}
}
