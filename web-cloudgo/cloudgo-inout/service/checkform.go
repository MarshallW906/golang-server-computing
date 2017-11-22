package service

import (
	"fmt"
	"html/template"
	"net/http"
)

func checkform(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "POST" {
		r.ParseForm()
		username := template.HTMLEscapeString(r.Form.Get("username"))
		password := template.HTMLEscapeString(r.Form.Get("password"))
		token := template.HTMLEscapeString(r.Form.Get("token"))
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		t := template.Must(template.New("checkform.html").ParseFiles("templates/checkform.html"))
		err := t.Execute(w, struct {
			Username string
			Password string
			Token    string
		}{Username: username, Password: password, Token: token})
		if err != nil {
			panic(err)
		}
	} else {
		http.Redirect(w, r, "/unknown", http.StatusTemporaryRedirect)
	}
}
