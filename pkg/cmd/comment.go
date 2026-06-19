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

var commentSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Submit a user comment or checkin for a specific charging location",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "charge-point-id",
			Usage:    "This must be a valid POI ID",
			Required: true,
			BodyPath: "chargePointID",
		},
		&requestflag.Flag[int64]{
			Name:     "checkin-status-type-id",
			Usage:    "Optional valid CheckStatusTypeID to indicate overall catgeory and success/failure to use equipment e.g. 10 = Charged Successfully.",
			BodyPath: "checkinStatusTypeID",
		},
		&requestflag.Flag[string]{
			Name:     "comment",
			Usage:    "This is an optional comment to describe the charging experience, may include guidance for future users.",
			BodyPath: "comment",
		},
		&requestflag.Flag[int64]{
			Name:     "comment-type-id",
			Usage:    "This must be a valid Comment Type ID as per UserCommentTypes found in Core Reference Data. If left as null then General Comment will be used.",
			BodyPath: "commentTypeID",
		},
		&requestflag.Flag[int64]{
			Name:     "rating",
			Usage:    "Optional integer rating between 1 = Worst, 5 = Best.",
			BodyPath: "rating",
		},
		&requestflag.Flag[string]{
			Name:     "related-url",
			Usage:    "Optional website URL for related information",
			BodyPath: "relatedURL",
		},
		&requestflag.Flag[string]{
			Name:     "user-name",
			Usage:    "This is an optional name to associate with the submission, for authenticated users their profile username is used.",
			BodyPath: "userName",
		},
	},
	Action:          handleCommentSubmit,
	HideHelpCommand: true,
}

func handleCommentSubmit(ctx context.Context, cmd *cli.Command) error {
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

	params := ocm.CommentSubmitParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Comment.Submit(ctx, params, options...)
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
		Title:          "comment submit",
		Transform:      transform,
	})
}
