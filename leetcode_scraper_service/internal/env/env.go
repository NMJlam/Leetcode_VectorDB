package env 

import ( 
	"github.com/joho/godotenv"
	"os"
	"log"
)

func Get(auth string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failure to load .env", err)
	}
	key := os.Getenv(auth)

	if key == "" {
		log.Fatal("Auth is not set in .env file")
	}
	return key 
}
