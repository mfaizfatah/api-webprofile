package config

import (
	"os"
	"strings"

	"github.com/mfaizfatah/api-webprofile/app/adapter"
)

/*LoadConfig is for load all configuration, connection etc
 * @parameters
 * service is for service name
 */
func LoadConfig(service string) {
	dbdriver := strings.ToLower(os.Getenv("DB_DRIVER"))
	if dbdriver == "mysql" {
		sql := os.Getenv("DB_URI")
		adapter.LoadMySQL(sql)
	}
}
