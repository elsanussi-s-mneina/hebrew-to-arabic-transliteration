package main

import "testing"

func TestTrimLeftOneRune_hello(t *testing.T) {
	got := trimLeftOneRune("hello")
	if got != "ello" {
		t.Errorf("failed with: %s", got)
	}
}

func TestTrimLeftOneRune_emptyStringCase(t *testing.T) {
	got := trimLeftOneRune("")
	if got != "" {
		t.Errorf("failed with: %s", got)
	}
}

func TestTransliterate_threeLetters(t *testing.T) {
	got := transliterate("שגר")
	if got != "شجر" {
		t.Errorf("failed to transliterate Arabic trees, got: %s", got)
	}
}
