package cmd

// 패키징 테스트
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla/v2"
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
	// CORS 추가
	mux.Use(CORSHandler)
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
	// Elastic APM 추가
	mux.Use(apmgorilla.Middleware())
	return mux
}

// CORSHandler는 모든 요청에 대해 CORS 헤더를 설정하는 미들웨어입니다.
func CORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// OPTIONS 요청에 대해 헤더만 설정하고 응답을 보냅니다.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// 다음 핸들러로 요청을 전달합니다.
		next.ServeHTTP(w, r)
	})
}
