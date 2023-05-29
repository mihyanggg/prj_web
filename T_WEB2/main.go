package main

import (
	"fmt"
	"net/http"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	// 전송된 파일은 r에 실려서 오니까, 읽어줘야 함..
	// input form으로 날라오는 파일을 읽는 method
	// input param으로 사용되는 key는, form명
	file, header, err := r.FormFile("upload_file")

	fmt.Println(file, header, err)
	w.Header("aaa", http.StatusOK)
}

func main() {

	http.HandleFunc("/uploads", uploadsHandler)

	// create.. static web server
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)

}
