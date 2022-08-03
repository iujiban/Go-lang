package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//strconv -> 변환
func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()    // 쿼리 인수 가져오기
	name := values.Get("name") //특정 키 값이 있는지 확인
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id")) //id값을 가져와서 int 타입 변환
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}
func MakeWebJandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Bar")
	})
	return mux
}
func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) //경로 테스트

	mux := MakeWebJandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code) //code 확인
	data, _ := io.ReadAll(res.Body)       //데이터를 읽어서 확인
	assert.Equal("Hello World", string(data))
}
func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil) //bar 경로 테스트

	mux := MakeWebJandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("HeLLO Bar", string(data))
}
func main() {
	/*
					// 웹 핸들러 등록
					// http request 클라이언트에는 보낸 메서드(method), 헤더 (header), 바디 (body) 같은 http 요청 정보
					http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
						fmt.Fprint(w, "Hello World")
					})

				http.HandleFunc("/bar", barHandler)
				//웹 서버 시작 : 함수를 호추해 웹 서버를 실행
				http.ListenAndServe(":3000", nil)


			//Multiplexer
			mux := http.NewServeMux() //serveMux 인스턴스 생성
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello World") //인스턴스에 핸들러 등록
			})
			mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello Bar")
			})

			http.ListenAndServe(":3000", mux) //mux 인스턴스 사용

		//StripPrefix: /static/ 제거"
		//http.FileServer(http.dir("static")) static 폴더 안에 있는 파일 서버 접근
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
		http.ListenAndServe(":3000", nil)
	*/
	http.ListenAndServe(":3000", MakeWebJandler())

}
