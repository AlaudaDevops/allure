/*
Copyright 2025 The AlaudaDevops Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package allure

import (
	"regexp"

	messages "github.com/cucumber/messages/go/v21"
	"github.com/godogx/allure/report"
)

// AllureLabelReg is a regular expression pattern for matching allure label tags.
var AllureLabelReg = regexp.MustCompile(`^@?allure\.label\.([^.]+):(.+)$`)

// matchAllureLabel parses a tag string to extract allure label key and value.
//
// Examples:
//
//	matchAllureLabel("@allure.label.epic:hello world") // returns "epic", "hello world", true
//	matchAllureLabel("allure.label.story:play")        // returns "story", "play", true
//	matchAllureLabel("@other.tag")                     // returns "", "", false
func matchAllureLabel(tag string) (key, value string, ok bool) {
	matches := AllureLabelReg.FindStringSubmatch(tag)
	if len(matches) != 3 {
		return "", "", false
	}
	return matches[1], matches[2], true
}

// GetAllureLabelsFromTags parse Allure labels based on tags provided by the user
func GetAllureLabelsFromTags(tags []*messages.PickleTag) []report.Label {
	labels := []report.Label{}
	for _, tag := range tags {
		if key, value, ok := matchAllureLabel(tag.Name); ok {
			labels = append(labels, report.Label{Name: key, Value: value})
		}
	}
	return labels
}
