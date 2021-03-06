package main

import (
	"log"
	"os"
	"text/template"
)

var tpl2 *template.Template

func init() {
	// 현재경로에서 tpl.gohtml파일을 찾아서 변수 tpl에 저장한다.
	tpl2 = template.Must(template.ParseFiles("tpl2.gohtml"))
}

func main() {
	// tpl템플릿을 실행하고 3번째 인자로 data 42를 전해준다.
	// 3번째 인자의 타입은 인터페이스타입
	// gohtml은 {{.}}을 통해 가장 최근에 보낸 데이터를 받아온다.
	err := tpl2.ExecuteTemplate(os.Stdout, "tpl2.gohtml", `byungwook`)
	if err != nil {
		log.Fatalln(err)
	}
}
