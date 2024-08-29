run_local:
	go run main.go

run_db_migrate_up:
	go run database/migrate/up/up.go

run_db_migrate_down:
	go run database/migrate/down/down.go

run_db_seed_product_category:
	go run database/seed/product-category/product_category.go