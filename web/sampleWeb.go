package web

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func SampleWeb() {
	Init() //数据库初始化
	http.HandleFunc("/index", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("html/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:"+k, "value:"+strings.Join(v, ""))
	}
	//输出到客户端
	fmt.Fprintf(w, "hello loren")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(token)
		t, _ := template.ParseFiles("html/login.html")
		fmt.Println(t.Execute(w, token))
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
			fmt.Println(token)
		} else {
			//不存在token报错
		}
		//
		//fmt.Println("username:", r.Form["username"])
		//fmt.Println("password:", r.Form["password"])
		//
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("password")))
		//
		//t, err1 := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		//err1 = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
		//fmt.Println(t,err1)
		template.HTMLEscape(w, []byte(r.Form.Get("username")+": login success!")) //输出到客户端
	}
}
