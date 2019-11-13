new Vue({
    // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
    el: '#app',

    // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
    data: {
        // 予定情報
        plans: [],
        // 予定名
        planName: '',
        //　工数
        planCost: '',
        // 優先度
        planPriority: '1',
        // メモ
        planMemo: '',
        // 予定の状態
        current: 0,
        // 予定の状態一覧
        options: [
            { value: -1, label: 'すべて' },
            { value:  0, label: '未達成' },
            { value:  1, label: '達成済' }
        ],
        selected: [],
        prioritys: [
            { value: 1, label: '1(高)'},
            { value: 2, label: '2(中)'},
            { value: 3, label: '3(低)'},
        ],
        // true：入力済・false：未入力
        isEntered: false,
    },

    // 算出プロパティ
    computed: {
        // 予定の状態一覧を表示する
        labels() {
            return this.options.reduce(function (a, b) {
                return Object.assign(a, { [b.value]: b.label })
            }, {})
        },
        // 表示対象の予定情報を返却する
        computedPlans() {
            return this.plans.filter(function (el) {
                var option = this.current < 0 ? true : this.current === el.state
                return option
            }, this)
        },
        // 入力チェック
        validate() {
            var isEnteredPlanName = 0 < this.planName.length
            this.isEntered = isEnteredPlanName
            return isEnteredPlanName
        }
    },

    // インスタンス作成時の処理
    created: function() {
        this.doFetchAllPlans()
    },

    // メソッド定義
    methods: {
        // 全ての予定情報を取得する
        doFetchAllPlans() {
            axios.get('/fetchAllPlans')
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        var resultPlans = response.data

                        console.log(resultPlans)
                        // サーバから取得した予定情報をdataに設定する
                        this.plans = resultPlans
                    }
                })
        },
        // １つの予定情報を取得する
        doFetchPlan(plan) {
            axios.get('/fetchPlan', {
                params: {
                    planID: plan.id
                }
            })
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        var resultPlan = response.data

                        // 選択された予定情報のインデックスを取得する
                        var index = this.plans.indexOf(plan)

                        // spliceを使うとdataプロパティの配列の要素をリアクティブに変更できる
                        this.plans.splice(index, 1, resultPlan[0])

                        //this.selected = resultPlan[0].planPriority
                    }
                })
        },
        // 予定情報を登録する
        doAddPlan() {
            // サーバへ送信するパラメータ
            const params = new URLSearchParams();
            params.append('planName', this.planName)
            params.append('planCost', this.planCost)
            params.append('planPriority', this.planPriority)
            params.append('planMemo', this.planMemo)


            axios.post('/addPlan', params)
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        // 予定情報を取得する
                        this.doFetchAllPlans()

                        // 入力値を初期化する
                        this.initInputValue()
                    }
                })
        },
        // 予定情報の状態を変更する
        doChangePlanState(plan) {
            // サーバへ送信するパラメータ
            const params = new URLSearchParams();
            params.append('planID', plan.id)
            params.append('planState', plan.state)

            axios.post('/changeStatePlan', params)
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        // 予定情報を取得する
                        this.doFetchPlan(plan)
                    }
                })
        },
        // 予定情報を削除する
        doDeletePlan(plan) {
            // サーバへ送信するパラメータ
            const params = new URLSearchParams();
            params.append('planID', plan.id)

            axios.post('/deletePlan', params)
                .then(response => {
                    if (response.status != 200) {
                        throw new Error('レスポンスエラー')
                    } else {
                        // 予定情報を取得する
                        this.doFetchAllPlans()
                    }
                })
        },
        // 入力値を初期化する
        initInputValue() {
            this.current = 0
            this.planName = ''
            this.planCost = ''
            this.planPriority = '1'
            this.planMemo = ''
        }
    }
})