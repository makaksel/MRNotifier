package redis

import "fmt"

func MRKey(projectPath string, mrID int) string {
	return fmt.Sprintf("mr:%s:%d", projectPath, mrID)
}
