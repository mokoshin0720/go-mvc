// リクエストを変換し、modelのロジックを呼び出し、viewに変換する
package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tutorial/mvc/pkg/model"
)

var db = model.GetDBConn()

func Index(w http.ResponseWriter, r *http.Request) {
	sample := model.Sample{Message: "modelの導入できた"}
	res, err := json.Marshal(sample)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func AllUsers(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "全てのユーザーを取得する")
	users := []model.User{}
	db.Find(&users)

	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func Userid(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "指定したユーザーを取得する")
	users := []model.User{}

	v := r.URL.Query()
	if v == nil {
		fmt.Fprintf(w, "ユーザーを入力してください")
	}

	// QueryStringで指定したユーザーをDBから取得
	for _, vs := range v {
		db.Where("name = ?", vs[0]).Find(&users)
	}

	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ユーザーを作成（from body value）\n")

	e := r.ParseForm()
	log.Println(e)
	name := r.Form.Get("name")

	user := model.User{Name: name}
	db.Create(&user)
	fmt.Fprintf(w, "ユーザー（%v）を作成しました", name)
}