package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"main/cmd/domain"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Data check
	user := new(domain.TMember)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	user.CustomDate = time.Now()
	insertMember(user)

	// TODO ::: 사용자 등록이 성공인지 실패인지 로직 추가해야함

	data, _ := json.Marshal(user)

	response := StandardResponse{
		Code:      http.StatusOK,
		Message:   "OK",
		Timestamp: time.Now(),
		Data:      data,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Data check
	user := new(domain.TMember)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	user.CustomDate = time.Now() // 추후 Update시간으로
	updateMember(user)

	// TODO ::: 사용자 등록이 성공인지 실패인지 로직 추가해야함

	data, _ := json.Marshal(user)

	response := StandardResponse{
		Code:      http.StatusOK,
		Message:   "OK",
		Timestamp: time.Now(),
		Data:      data,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// 전체 사용자 리스트
func GetUserListHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(selectMember()) // JSON 직렬화
	// log.Printf("전체 API %s", data)
	response := StandardResponse{
		Code:      http.StatusOK,
		Message:   "OK",
		Timestamp: time.Now(),
		Data:      data,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// 단일 사용자 리스트
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, _ := json.Marshal(selectMemberOne(vars["id"]))

	response := StandardResponse{
		Code:      http.StatusOK,
		Message:   "OK",
		Timestamp: time.Now(),
		Data:      data,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// 사용자 삭제
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := deleteMemberOne(vars["id"])
	log.Printf("삭제결과 : %t", result)
	if !result {
		response := StandardResponse{
			Code:      http.StatusInternalServerError,
			Message:   "삭제실패",
			Timestamp: time.Now(),
			Data:      nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	} else {
		response := StandardResponse{
			Code:      http.StatusOK,
			Message:   "OK",
			Timestamp: time.Now(),
			Data:      nil,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
