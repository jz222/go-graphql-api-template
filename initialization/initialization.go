package initialization

import (
	"graphql_api/keys"
	"graphql_api/libs/mongodb"
	"log"
)

func InitEnv() {
	keys.GetKeys()

	log.Println("✅ environment variables initialized successfully")
}

func InitDatabase() {
	mongodb.InitiateDatabase()
}
