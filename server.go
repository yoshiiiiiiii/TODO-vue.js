package main

import (

	// ロギングを行うパッケージ
	"log"

	// HTTPを扱うパッケージ
	"net/http"

	// Gin
	"github.com/gin-gonic/gin"

	// MySQL用ドライバ
	_ "github.com/jinzhu/gorm/dialects/mysql"

	//controller
	controller "todo_ver2/controllers/controller"
)

func main() {
	// サーバーを起動する
	serve()
}

func serve() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	r := gin.Default()

	// 静的ファイルのパスを指定
	r.Static("/views", "./views")

	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す
	r.StaticFS("/todo", http.Dir("./views/static"))

	// 全ての予定情報のJSONを返す
	r.GET("/fetchAllPlans", controller.FetchAllPlans)

	// １つの予定情報の状態のJSONを返す
	r.GET("/fetchPlan", controller.FindPlan)

	// 予定情報をDBへ登録する
	r.POST("/addPlan", controller.AddPlan)

	// 予定情報の状態を変更する
	r.POST("/changeStatePlan", controller.ChangeStatePlan)

	// 予定情報を削除する
	r.POST("/deletePlan", controller.DeletePlan)

	if err := r.Run(":8888"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}