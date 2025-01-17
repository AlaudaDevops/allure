package allure

import (
	"regexp"

	messages "github.com/cucumber/messages/go/v21"
	"github.com/godogx/allure/report"
)

// AllureLabelReg is a regular expression pattern for matching allure label tags.
var AllureLabelReg = regexp.MustCompile(`^@?allure\.label\.([^.]+):(.+)$`)

// MatchAllureLabel parses a tag string to extract allure label key and value.
//
// Examples:
//
//	MatchAllureLabel("@allure.label.epic:hello world") // returns "epic", "hello world", true
//	MatchAllureLabel("allure.label.story:play")        // returns "story", "play", true
//	MatchAllureLabel("@other.tag")                     // returns "", "", false
func MatchAllureLabel(tag string) (key, value string, ok bool) {
	matches := AllureLabelReg.FindStringSubmatch(tag)
	if len(matches) != 3 {
		return "", "", false
	}

	return matches[1], matches[2], true
}

// GetAllureLabelsFromTags parse Allure labels based on tags provided by the user.
func GetAllureLabelsFromTags(tags []*messages.PickleTag) []report.Label {
	var labels []report.Label

	for _, tag := range tags {
		if key, value, ok := MatchAllureLabel(tag.Name); ok {
			labels = append(labels, report.Label{Name: key, Value: value})
		}
	}

	return labels
}
