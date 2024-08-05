package cmd

// 패키징 테스트
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Main Page")
}

// default content-type setting
func jsonResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set Content-Type header to application/json for all responses
		w.Header().Set("Content-Type", "application/json")
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// 라우터 등록
func NewHttpHandler() http.Handler {
	mux := mux.NewRouter() // gorilla/mux
	mux.Use(jsonResponseMiddleware)
	mux.NotFoundHandler = http.HandlerFunc(notFoundHandler)       // 404
	mux.HandleFunc("/api/example", exampleHandler).Methods("GET") // 200

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/v1/users", GetUserListHandler).Methods("GET")                  // 사용자 검색
	mux.HandleFunc("/api/v1/users", CreateUserHandler).Methods("PUT")                   // 사용자 등록
	mux.HandleFunc("/api/v1/users", UpdateUserHandler).Methods("POST")                  // 사용자 수정
	mux.HandleFunc("/api/v1/users/{id:[A-z0-9]+}", DeleteUserHandler).Methods("DELETE") // 사용자 삭제
	mux.HandleFunc("/api/v1/users/{id:[A-z0-9]+}", GetUserHandler).Methods("GET")       // 사용자 개별 검색
	// API 호출 BATCH
	mux.HandleFunc("/api/v1/batch", ApiGetUserAllHandler).Methods("GET")             // 사용자 전체 검색 API 를 호출하여 temp table 적재
	mux.HandleFunc("/api/v1/batch/{id:[A-z0-9]+}", ApiGetUserHandler).Methods("GET") // 사용자 개별 검색 API 를 호출하여 temp table 적재

	return mux
}
