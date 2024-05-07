package migrations

import (
	"my-task-api/app/configs"
	"my-task-api/app/databases"
	"my-task-api/features/user/data"
)

func InitialMigration() {
	DB := databases.InitDBMysql(configs.InitConfig())

	// Migrate model-model yang Anda miliki
	DB.AutoMigrate(&data.User{})
}
