package util

import (
	"strconv"
)

func GetDatabaseString(username string, password string, url string, port int, database string) string {
	return username + ":" + password + "@tcp" + "(" + url + ":" + strconv.Itoa(port) + ")/" + database + "?parseTime=true"
}
