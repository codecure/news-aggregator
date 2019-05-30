package main

import (
	"testing"
	"time"
)

func TestParseRule(t *testing.T) {
	rule := parseRule("")
	emptyRule := feedRule{}
	if rule != emptyRule {
		t.Errorf("Result should be empty, got %s", rule)
	}
	rule = parseRule("date=someDateField&content=someContentField")
	correctRule := feedRule{date: "someDateField", content: "someContentField"}
	if rule != correctRule {
		t.Errorf("Result should be %s, got: %s", correctRule, rule)
	}
}

func TestGetField(t *testing.T) {
	type testStruct struct {
		Date    time.Time
		Content string
	}

	value := testStruct{Date: time.Now(), Content: "Test"}
	formattedDate := value.Date.Format(time.RFC3339)
	dateFieldValue := getField(value, "Date")
	contentFieldValue := getField(value, "Content")

	if formattedDate != dateFieldValue {
		t.Errorf("Result should be %s, got: %s", formattedDate, dateFieldValue)
	}
	if value.Content != contentFieldValue {
		t.Errorf("Result should be %s, got: %s", value.Content, contentFieldValue)
	}
}
