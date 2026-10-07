// Hotel Lobby AI Prompts creates short storyboard prompts for two-photo music videos.
//
// Homepage: https://hotellobbyai.net/
package main

import (
	"flag"
	"fmt"
	"strings"
)

func prompt(subjects, mood string) string {
	subjects = strings.TrimSpace(subjects)
	mood = strings.TrimSpace(mood)
	return fmt.Sprintf(
		"Animate %s as a playful two-person rap duet. Use a %s mood, clear lip-sync, gentle camera movement, and a clean final pose.",
		subjects,
		mood,
	)
}

func main() {
	subjects := flag.String("subjects", "two friends", "people or characters in the two photos")
	mood := flag.String("mood", "upbeat", "visual and musical mood")
	flag.Parse()
	fmt.Println(prompt(*subjects, *mood))
}
