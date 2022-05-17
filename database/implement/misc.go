package implement

func (db *ORM) RowExists(model interface{}) bool {
	row := db.Limit(1).Find(&model)
	return row.RowsAffected > 0
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
