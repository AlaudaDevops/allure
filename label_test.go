package allure

import (
	"testing"

	messages "github.com/cucumber/messages/go/v21"
	"github.com/stretchr/testify/assert"
)

func TestMatchAllureLabel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantKey  string
		wantVal  string
		wantBool bool
	}{
		{
			name:     "tag with @ prefix",
			input:    "@allure.label.epic:hello world",
			wantKey:  "epic",
			wantVal:  "hello world",
			wantBool: true,
		},
		{
			name:     "tag without @ prefix",
			input:    "allure.label.story:play",
			wantKey:  "story",
			wantVal:  "play",
			wantBool: true,
		},
		{
			name:     "non allure tag",
			input:    "@other.tag",
			wantKey:  "",
			wantVal:  "",
			wantBool: false,
		},
		{
			name:     "invalid tag format",
			input:    "@allure.label.",
			wantKey:  "",
			wantVal:  "",
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, val, ok := matchAllureLabel(tt.input)
			assert.Equal(t, tt.wantKey, key)
			assert.Equal(t, tt.wantVal, val)
			assert.Equal(t, tt.wantBool, ok)
		})
	}
}

func TestGetAllureLabelsFromTags(t *testing.T) {
	tests := []struct {
		name          string
		inputTags     []*messages.PickleTag
		expectedCount int
		expectedTags  []struct {
			name  string
			value string
		}
	}{
		{
			name: "multiple valid tags",
			inputTags: []*messages.PickleTag{
				{Name: "@allure.label.epic:MyEpic"},
				{Name: "@allure.label.story:MyStory"},
				{Name: "allure.label.severity:critical"},
			},
			expectedCount: 3,
			expectedTags: []struct {
				name  string
				value string
			}{
				{"epic", "MyEpic"},
				{"story", "MyStory"},
				{"severity", "critical"},
			},
		},
		{
			name: "mixed valid and invalid tags",
			inputTags: []*messages.PickleTag{
				{Name: "@some.other.tag"},
				{Name: "@allure.label.epic:MyEpic"},
				{Name: "not a label tag"},
			},
			expectedCount: 1,
			expectedTags: []struct {
				name  string
				value string
			}{
				{"epic", "MyEpic"},
			},
		},
		{
			name:          "empty tag list",
			inputTags:     []*messages.PickleTag{},
			expectedCount: 0,
			expectedTags:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			labels := GetAllureLabelsFromTags(tt.inputTags)
			assert.Equal(t, tt.expectedCount, len(labels))

			if tt.expectedCount > 0 {
				for i, expected := range tt.expectedTags {
					assert.Equal(t, expected.name, labels[i].Name)
					assert.Equal(t, expected.value, labels[i].Value)
				}
			}
		})
	}
}
