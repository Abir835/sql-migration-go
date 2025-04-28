package db

func InitDB() {
	ConnectDB()
	MigrationDB()
}
