package report

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStepFinishedAggregatesResultStatus verifies that skipped steps do not hide failure statuses.
func TestStepFinishedAggregatesResultStatus(t *testing.T) {
	t.Parallel()

	passedDetails := &StatusDetails{Message: "passed details"}
	skippedDetails := &StatusDetails{Message: "skipped details"}
	failedDetails := &StatusDetails{Message: "failed details"}
	brokenDetails := &StatusDetails{Message: "broken details"}

	testCases := []struct {
		name  string
		steps []struct {
			status  Status
			details *StatusDetails
		}
		wantStatus  Status
		wantDetails *StatusDetails
	}{
		{
			name: "passed followed by skipped",
			steps: []struct {
				status  Status
				details *StatusDetails
			}{
				{status: Passed, details: passedDetails},
				{status: Skipped, details: skippedDetails},
			},
			wantStatus:  Skipped,
			wantDetails: skippedDetails,
		},
		{
			name: "failed followed by skipped",
			steps: []struct {
				status  Status
				details *StatusDetails
			}{
				{status: Failed, details: failedDetails},
				{status: Skipped, details: skippedDetails},
			},
			wantStatus:  Failed,
			wantDetails: failedDetails,
		},
		{
			name: "broken followed by skipped",
			steps: []struct {
				status  Status
				details *StatusDetails
			}{
				{status: Broken, details: brokenDetails},
				{status: Skipped, details: skippedDetails},
			},
			wantStatus:  Broken,
			wantDetails: brokenDetails,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result := &Result{}
			for _, step := range testCase.steps {
				StepFinished(result, string(step.status), step.status, step.details, nil, 0)
			}

			assert.Equal(t, testCase.wantStatus, result.Status)
			assert.Same(t, testCase.wantDetails, result.StatusDetails)
		})
	}
}
