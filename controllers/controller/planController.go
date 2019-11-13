package controller

import (
	"log"
	// 文字列と基本データ型の変換パッケージ
	strconv "strconv"

	// Gin
	"github.com/gin-gonic/gin"

	// エンティティ(データベースのテーブルの行に対応)
	entity "todo_ver2/models/entity"

	// DBアクセス用モジュール
	db "todo_ver2/models/db"
)

// 曲状態を定義
const (
	// NotPurchased は 未達成
	NotPurchased = 0

	// Purchased は 達成済
	Purchased = 1
)

// FetchAllPlans は 全ての予定を取得する
func FetchAllPlans(c *gin.Context) {
	resultPlans := db.FindAllPlans()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultPlans)
}

// FindPlan は 指定したIDの予定を取得する
func FindPlan(c *gin.Context) {
	planIDStr := c.Query("planID")

	planID, _ := strconv.Atoi(planIDStr)

	resultPlan := db.FindPlan(planID)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultPlan)
}

// AddPlan は 予定をDBへ登録する
func AddPlan(c *gin.Context) {
	planName := c.PostForm("planName")
	planCostStr := c.PostForm("planCost")
	planPriorityStr := c.PostForm("planPriority")
	planMemo := c.PostForm("planMemo")

	planCost, _ := strconv.Atoi(planCostStr)
	planPriority, _ := strconv.Atoi(planPriorityStr)

	log.Print(c.PostForm)
	var plan = entity.Plan{
		Name:  planName,
		Cost: planCost,
		Priority: planPriority,
		Memo:  planMemo,
		State: NotPurchased,
	}

	log.Print(plan)
	db.InsertPlan(&plan)
}

// ChangeStatePlan は 予定の状態を変更する
func ChangeStatePlan(c *gin.Context) {
	reqPlanID := c.PostForm("planID")
	reqPlanState := c.PostForm("planState")

	planID, _ := strconv.Atoi(reqPlanID)
	planState, _ := strconv.Atoi(reqPlanState)
	changeState := NotPurchased

	// 予定が未達成の場合
	if planState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}

	db.UpdateStatePlan(planID, changeState)
}

// DeletePlan は 予定をDBから削除する
func DeletePlan(c *gin.Context) {
	planIDStr := c.PostForm("planID")

	planID, _ := strconv.Atoi(planIDStr)

	db.DeletePlan(planID)
}