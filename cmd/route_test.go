package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestIndexHandlerTesting(t *testing.T) {
// 	res := httptest.NewRecorder()
// 	req := httptest.NewRequest("GET", "/", nil)

// 	indexHandler(res, req)

// 	if res.Code != http.StatusOK {
// 		t.Fatal("Failed!! ", res.Code)
// 	}
// }

// func TestIndexHandlerAssert(t *testing.T) {
// 	// assert 를 사용한 자동 테스트
// 	assert := assert.New(t)

// 	res := httptest.NewRecorder()
// 	req := httptest.NewRequest("GET", "/", nil)
// 	indexHandler(res, req)

// 	assert.Equal(http.StatusOK, res.Code)
// 	// 버퍼 갚을 다 읽어오자
// 	data, _ := io.ReadAll(res.Body)
// 	assert.Equal("indehandler 입니다.", string(data)) // 응답에 대한 기대값을 체크하는 로직

// }

func TestIndexPathHandlerMux(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHttpHandler()) // mock 서버
	defer ts.Close()                           // mock 서버를 꼭 닫아줘야함.

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode) // 응답에 대한 기대값을 체크하는 로직

}

// func TestCreateUser(t *testing.T) {
// 	assert := assert.New(t)

// 	ts := httptest.NewServer(NewHttpHandler()) // mock 서버
// 	defer ts.Close()                           // mock 서버를 꼭 닫아줘야함.

// 	resp, err := http.Post(ts.URL+"/api/v1/users", "application/json",
// 		strings.NewReader(`{"UserId":"23451","FirstName":"joonhyeok","LastName":"lim","Email":"joon95@metanet.co.kr"}`))

// 	assert.NoError(err)
// 	defer resp.Body.Close()

// 	assert.Equal(http.StatusCreated, resp.StatusCode) // 응답에 대한 기대값을 체크하는 로직
// }
