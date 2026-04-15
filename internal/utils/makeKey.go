package utils

import "fmt"

func MakeKey(projectPath string, mrId int) string {
	return fmt.Sprintf("%s:%d", projectPath, mrId)
}
