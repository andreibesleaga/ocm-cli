// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/ocm-cli/internal/apiquery"
	"github.com/stainless-sdks/ocm-cli/internal/requestflag"
	"github.com/stainless-sdks/ocm-go"
	"github.com/stainless-sdks/ocm-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var profileAuthenticate = cli.Command{
	Name:    "authenticate",
	Usage:   "Perform user authentication, returning a model which includes the users profile\nand a JWT auth token to re-use in subsequent requests.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "emailaddress",
			BodyPath: "emailaddress",
		},
		&requestflag.Flag[string]{
			Name:     "password",
			BodyPath: "password",
		},
	},
	Action:          handleProfileAuthenticate,
	HideHelpCommand: true,
}

func handleProfileAuthenticate(ctx context.Context, cmd *cli.Command) error {
	client := ocm.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := ocm.ProfileAuthenticateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Profile.Authenticate(ctx, params, options...)
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
		Title:          "profile authenticate",
		Transform:      transform,
	})
}
