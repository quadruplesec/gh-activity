package github

import (
	"os"
)

// This function should only return true if 3 conditions are met:
// 1. The NO_COLOR environment variable is non-empty
// 2. The TERM environment variable is not set to "dumb"
// 3. stdout is an interactive terminal (TTY)
func ShouldUseColors() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}

	if os.Getenv("TERM") == "dumb" {
		return false
	}

	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}

	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
)

func GetEventColor(eventType string) string {
	switch eventType {
	case "PushEvent", "ReleaseEvent":
		return Cyan
	case "CreateEvent", "WatchEvent", "ForkEvent", "MemberEvent":
		return Green
	case "PullRequestEvent", "PullRequestReviewEvent", "PullRequestReviewCommentEvent":
		return Magenta
	case "IssuesEvent", "IssueCommentEvent", "CommitCommentEvent", "DiscussionEvent", "GollumEvent":
		return Yellow
	case "DeleteEvent":
		return Red
	default:
		return Reset
	}
}

// Colors text with ANSI escape codes if the terminal supports it
func Colorize(colorCode, text string) string {
	if !ShouldUseColors() {
		return text
	}
	return colorCode + text + Reset
}
