// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/andreibesleaga/ocm-cli/internal/mocktest"
)

func TestMediaitemCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"--bearer", "string",
			"mediaitem", "create",
			"--charge-point-id", "1234",
			"--image-data-base64", "data:image/jpeg;base64,<BASE64_ENCODED_DATA>",
			"--comment", "An example comment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"chargePointID: 1234\n" +
			"imageDataBase64: data:image/jpeg;base64,<BASE64_ENCODED_DATA>\n" +
			"comment: An example comment\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"--bearer", "string",
			"mediaitem", "create",
		)
	})
}
