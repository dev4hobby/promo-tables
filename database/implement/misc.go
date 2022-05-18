package implement

func (db *ORM) RowExists(model interface{}, limit int) int {
	rows := db.Limit(limit).Find(&model)
	return int(rows.RowsAffected)
}

func (db *ORM) InitTable(model interface{}) {
	err := db.AutoMigrate(&model)
	if err != nil {
		panic(err)
	}
}

func (db *ORM) FlushAll() {
	db.Exec("DROP TABLE `store`.`categories`, `store`.`orders`, `store`.`products`, `store`.`promo_categories`, `store`.`promo_controllers`, `store`.`promo_purchased_logs`, `store`.`users`, `store`.`user_store_subscriptions`")
}
