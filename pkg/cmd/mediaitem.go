// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/andreibesleaga/ocm-cli/internal/apiquery"
	"github.com/andreibesleaga/ocm-cli/internal/requestflag"
	"github.com/andreibesleaga/ocm-go"
	"github.com/andreibesleaga/ocm-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var mediaitemCreate = cli.Command{
	Name:    "create",
	Usage:   "Submit a photo for a specific charging location",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "charge-point-id",
			Usage:    "ID value for the OCM site (POI) this image relates to.",
			Required: true,
			BodyPath: "chargePointID",
		},
		&requestflag.Flag[string]{
			Name:     "image-data-base64",
			Usage:    "BASE64 encoded data",
			Required: true,
			BodyPath: "imageDataBase64",
		},
		&requestflag.Flag[string]{
			Name:     "comment",
			Usage:    "Optional description of image or context",
			BodyPath: "comment",
		},
	},
	Action:          handleMediaitemCreate,
	HideHelpCommand: true,
}

func handleMediaitemCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := ocm.MediaitemNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Mediaitem.New(ctx, params, options...)
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
		Title:          "mediaitem create",
		Transform:      transform,
	})
}
