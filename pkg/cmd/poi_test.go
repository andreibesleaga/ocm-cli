// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/ocm-cli/internal/mocktest"
)

func TestPoiList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"poi", "list",
			"--boundingbox", "{}",
			"--camelcase=true",
			"--chargepointid", "chargepointid",
			"--client", "client",
			"--compact=true",
			"--connectiontypeid", "{}",
			"--countrycode", "countrycode",
			"--countryid", "string",
			"--dataproviderid", "{}",
			"--distance", "0",
			"--distanceunit", "distanceunit",
			"--greaterthanid", "greaterthanid",
			"--includecomments=true",
			"--latitude", "0",
			"--levelid", "{}",
			"--longitude", "0",
			"--maxresults", "0",
			"--modifiedsince", "modifiedsince",
			"--opendata=true",
			"--operatorid", "{}",
			"--output", "output",
			"--polygon", "polygon",
			"--polyline", "polyline",
			"--sortby", "sortby",
			"--statustypeid", "{}",
			"--usagetypeid", "{}",
			"--verbose=true",
		)
	})
}
