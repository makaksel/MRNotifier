package replyCache

import "fmt"

func makeKey(projectPath string, mrID int) string {
	return fmt.Sprintf("reply:%s:%d", projectPath, mrID)
}
