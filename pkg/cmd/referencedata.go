// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/andreibesleaga/ocm-go"
	"github.com/andreibesleaga/ocm-go/option"
	"github.com/stainless-sdks/ocm-cli/internal/apiquery"
	"github.com/stainless-sdks/ocm-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var referencedataRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the core reference data used for looking up IDs such as Connection\nTypes, Operators, Countries etc.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]any]{
			Name:      "countryid",
			Usage:     "Optional filter on countryid, exact match on a given numeric country id (comma separated list)",
			QueryPath: "countryid",
		},
	},
	Action:          handleReferencedataRetrieve,
	HideHelpCommand: true,
}

func handleReferencedataRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := ocm.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := ocm.ReferencedataGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Referencedata.Get(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "referencedata retrieve",
		Transform:      transform,
	})
}
