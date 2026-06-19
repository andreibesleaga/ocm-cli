// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ocm-cli/internal/mocktest"
)

func TestCommentSubmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--bearer", "string",
			"comment", "submit",
			"--charge-point-id", "0",
			"--checkin-status-type-id", "0",
			"--comment", "string",
			"--comment-type-id", "0",
			"--rating", "3",
			"--related-url", "string",
			"--user-name", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"chargePointID: 0\n" +
			"checkinStatusTypeID: 0\n" +
			"comment: string\n" +
			"commentTypeID: 0\n" +
			"rating: 3\n" +
			"relatedURL: string\n" +
			"userName: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--bearer", "string",
			"comment", "submit",
		)
	})
}
