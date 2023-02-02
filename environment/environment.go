package environment

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Site is the data structure that represents an Site.
type Environment struct {
	Bearer     string
	Endpoint   string
	From       string
	Password   string
	T_Token    string
	T_ChatID   int
	HoraAvis   int
	MinutAvis  int
	HoraFinal  int
	MinutFinal int
	DbHost     string
	DbUser     string
	DbPass     string
	DbPort     int
	DbName     string
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
	horaFinal, err := strconv.Atoi(os.Getenv("HORAFINAL"))
	CheckConversionError("HORAFINAL", err)
	minutFinal, err := strconv.Atoi(os.Getenv("MINUTFINAL"))
	CheckConversionError("MINUTFINAL", err)
	horaAvis, err := strconv.Atoi(os.Getenv("HORAAVIS"))
	CheckConversionError("HORAAVIS", err)
	minutAvis, err := strconv.Atoi(os.Getenv("MINUTAVIS"))
	CheckConversionError("MINUTAVIS", err)
	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbPort, err := strconv.Atoi(os.Getenv("DBPORT"))
	CheckConversionError("DBPORT", err)
	dbName := os.Getenv("DBNAME")

	return Environment{
		Bearer:     bearer,
		Endpoint:   endpoint,
		From:       from,
		Password:   password,
		T_Token:    t_Token,
		T_ChatID:   t_ChatID,
		HoraAvis:   horaAvis,
		MinutAvis:  minutAvis,
		HoraFinal:  horaFinal,
		MinutFinal: minutFinal,
		DbHost:     dbHost,
		DbUser:     dbUser,
		DbPass:     dbPass,
		DbPort:     dbPort,
		DbName:     dbName,
	}
}

func CheckConversionError(varName string, err error) {
	if err != nil {
		log.Fatalf("%s environment variable not found", varName)
	}
}
