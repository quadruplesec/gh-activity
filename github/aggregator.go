package github

import "fmt"

type ActivityFeed []Event

func formatAggregatedPushes(count int, repo string) string {
	if count == 0 {
		return ""
	}
	word := "commits"
	if count == 1 {
		word = "commit"
	}

	action := Colorize(GetEventColor("PushEvent"), fmt.Sprintf("Pushed %d %s", count, word))
	targetRepo := Colorize(Bold, repo)
	return fmt.Sprintf(" - %s to %s", action, targetRepo)
}

func (feed ActivityFeed) Format() []string {
	var output []string
	pushCount := 0
	pushRepo := ""

	for _, event := range feed {
		if event.Type == "PushEvent" {
			if pushRepo == "" || pushRepo == event.Repo.Name {
				pushRepo = event.Repo.Name
				pushCount++
				continue
			}

			output = append(output, formatAggregatedPushes(pushCount, pushRepo))
			pushRepo = event.Repo.Name
			pushCount = 1
			continue
		}

		if pushCount > 0 {
			output = append(output, formatAggregatedPushes(pushCount, pushRepo))
			pushRepo = ""
			pushCount = 1
			continue
		}

		formatter, err := event.UnmarshalEventPayload()
		if err != nil {
			output = append(output, fmt.Sprintf("Error: %v", err))
			continue
		}
		output = append(output, formatter.FormatActivity(event.Actor, event.Repo))
	}

	// Flush any remaining pushes left over at the end
	if pushCount > 0 {
		output = append(output, formatAggregatedPushes(pushCount, pushRepo))
	}

	return output
}
