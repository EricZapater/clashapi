package environment

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Site is the data structure that represents an Site.
type Environment struct {
	Bearer   string
	Endpoint string
	From     string
	Password string
	T_Token  string
	T_ChatID int
}

// LoadEnvironment creates a new Environment.
func LoadEnvironment() Environment {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	godotenv.Load(".env")

	// getting env variables SITE_TITLE and DB_HOST
	bearer := os.Getenv("BEARER")
	endpoint := os.Getenv("ENDPOINT")
	from := os.Getenv("FROM")
	password := os.Getenv("PASSWORD")
	t_Token := os.Getenv("TELEGRAM_TOKEN")
	t_ChatID, err := strconv.Atoi(os.Getenv("TELEGRAM_CHATID"))
	CheckConversionError("TELEGRAM_CHATID", err)

	return Environment{
		Bearer:   bearer,
		Endpoint: endpoint,
		From:     from,
		Password: password,
		T_Token:  t_Token,
		T_ChatID: t_ChatID,
	}
}

func CheckConversionError(varName string, err error) {
	if err != nil {
		log.Fatalf("%s environment variable not found", varName)
	}
}
