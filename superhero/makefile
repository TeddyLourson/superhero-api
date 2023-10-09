sqlg:
	@sqlc generate
prod-up:
	@dbmate -e USR_DATABASE_URL up
test-up:
	@dbmate -e TEST_USR_DATABASE_URL up
prod-down:
	@dbmate -e USR_DATABASE_URL down
test-down:
	@dbmate -e TEST_USR_DATABASE_URL down
migration:
	@dbmate new $(NAME)