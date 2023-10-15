package configurations

import "os"

var (
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	CurrentSite      string
	BaseUrl          string
	EmailTokenSecret string
)

func SetEnvVariable() {
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	CurrentSite = os.Getenv("CURRENT_SITE")
	BaseUrl = os.Getenv("BASE_URL")
	EmailTokenSecret = os.Getenv("EMAIL_TOKEN_SECRET")
}
