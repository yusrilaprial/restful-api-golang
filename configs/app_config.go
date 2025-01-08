package configs

import "yusrilaprial/backend-api/utils"

var DB_HOST string = utils.GetEnv("DB_HOST", "127.0.0.1")
var DB_PORT string = utils.GetEnv("DB_PORT", "3306")
var DB_NAME string = utils.GetEnv("DB_NAME", "db_go_api")
var DB_USERNAME string = utils.GetEnv("DB_USERNAME", "root")
var DB_PASSWORD string = utils.GetEnv("DB_PASSWORD", "")
