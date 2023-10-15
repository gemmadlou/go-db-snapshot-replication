package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if os.Getenv("ENVIRONMENT") != "production" && err != nil {
		log.Fatal("Error loading .env file")
	}

	currentTime := time.Now()
	file := fmt.Sprintf("replay_%s_%s", currentTime.Format("20060102"), os.Getenv("SOURCE_DB_NAME"))

	cmdString := fmt.Sprintf(`
	export MYSQL_PWD=%s;

	mysqldump \
	-h %s \
	-u %s \
	--complete-insert \
	--compress \
	--quick \
	--single-transaction \
	--no-tablespaces \
	--max_allowed_packet=1G \
	--set-gtid-purged=OFF \
	--column-statistics=0 \
	--create-options \
	--ignore-table=%s.stream \
	%s | gzip > %s.sql.gz;

	export MYSQL_PWD=%s;

	mysql -h %s -u %s -e "CREATE DATABASE IF NOT EXISTS %s";
	`,
		os.Getenv("SOURCE_DB_PASS"),
		os.Getenv("SOURCE_DB_HOST"),
		os.Getenv("SOURCE_DB_USER"),
		os.Getenv("SOURCE_DB_NAME"),
		os.Getenv("SOURCE_DB_NAME"),
		file,
		os.Getenv("DESTINATION_DB_PASS"),
		os.Getenv("DESTINATION_DB_HOST"),
		os.Getenv("DESTINATION_DB_USER"),
		file,
	)

	cmdString += fmt.Sprintf(`
	pv %s.sql.gz | gunzip | mysql -h %s -u %s %s
	`,
		file,
		os.Getenv("DESTINATION_DB_HOST"),
		os.Getenv("DESTINATION_DB_USER"),
		file,
	)

	fmt.Println(cmdString)

}
