package main

import (
	"main/cmd"
	"net/http"
)

func main() {
	cmd.Env()                                          // 환경파일
	cmd.LogFormat()                                    // log 레벨
	cmd.InitDB()                                       // DB Connection
	http.ListenAndServe(":3000", cmd.NewHttpHandler()) // Server And Route
}
