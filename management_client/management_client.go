package management_client

import (
	"strconv"
)

func idToParams(id int) map[string]string {
	return id64ToParams(int64(id))
}

func id64ToParams(id int64) map[string]string {
	return map[string]string{"id": strconv.FormatInt(id, 10)}
}
