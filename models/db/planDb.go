package db

import (
	// フォーマットI/O
	"fmt"
	"log"

	// Go言語のORM
	"github.com/jinzhu/gorm"

	// エンティティ(データベースのテーブルの行に対応)
	entity "todo_ver2/models/entity"
)

// DB接続する
func open() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "Todo"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&entity.Plan{})

	fmt.Println("db connected: ", &db)
	return db
}

// FindAllPlans は planテーブルのレコードを全件取得する
func FindAllPlans() []entity.Plan {
	plans := []entity.Plan{}

	db := open()
	// select
	db.Order("priority asc, ID asc").Find(&plans)

	log.Print([]entity.Plan{})
	// defer 関数がreturnする時に実行される
	defer db.Close()

	return plans
}

// FindPlan は planテーブルのレコードを１件取得する
func FindPlan(planID int) []entity.Plan {
	plan := []entity.Plan{}

	db := open()
	// select
	db.First(&plan, planID)
	defer db.Close()

	return plan
}

// InsertPlan は planテーブルにレコードを追加する
func InsertPlan(registerPlan *entity.Plan) {
	db := open()
	// insert
	db.Create(&registerPlan)
	defer db.Close()
}

// UpdateStatePlan は planテーブルの指定したレコードの状態を変更する
func UpdateStatePlan(planID int, planState int) {
	plan := []entity.Plan{}

	db := open()
	// update
	db.Model(&plan).Where("ID = ?", planID).Update("State", planState)
	defer db.Close()
}

// DeletePlan は planテーブルの指定したレコードを削除する
func DeletePlan(planID int) {
	plan := []entity.Plan{}

	db := open()
	// delete
	db.Delete(&plan, planID)
	defer db.Close()
}