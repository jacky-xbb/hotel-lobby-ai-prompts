package main

import "testing"

func TestPromptUsesSubjectsAndMood(t *testing.T) {
	got := prompt("a chef and a cat", "jazzy")
	want := "Animate a chef and a cat as a playful two-person rap duet. Use a jazzy mood, clear lip-sync, gentle camera movement, and a clean final pose."
	if got != want {
		t.Fatalf("prompt() = %q, want %q", got, want)
	}
}
