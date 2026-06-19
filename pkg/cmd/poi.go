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

var poiList = cli.Command{
	Name:    "list",
	Usage:   "Used to fetch a list of POIs (sites) within a geographic boundary or near a\nspecific latitude/longitude. This is the primary method for most applications\nand services to consume data from Open Charge Map.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]any]{
			Name:      "boundingbox",
			Usage:     "Filter results to a given bounding box. specify top left and bottom right box corners as: (lat,lng),(lat2,lng2)",
			QueryPath: "boundingbox",
		},
		&requestflag.Flag[bool]{
			Name:      "camelcase",
			Usage:     "Set to true to get a property names in camelCase format.",
			Default:   false,
			QueryPath: "camelcase",
		},
		&requestflag.Flag[string]{
			Name:      "chargepointid",
			Usage:     "Exact match on a given OCM POI ID (comma separated list)",
			QueryPath: "chargepointid",
		},
		&requestflag.Flag[string]{
			Name:      "client",
			Usage:     "String to identify your client application. Optional but recommended to distinguish your client from other bots/crawlers",
			QueryPath: "client",
		},
		&requestflag.Flag[bool]{
			Name:      "compact",
			Usage:     "Set to true to remove reference data objects from output (just returns IDs for common reference data such as DataProvider etc).",
			QueryPath: "compact",
		},
		&requestflag.Flag[[]any]{
			Name:      "connectiontypeid",
			Usage:     "Exact match on a given connection type id (comma separated list)",
			QueryPath: "connectiontypeid",
		},
		&requestflag.Flag[string]{
			Name:      "countrycode",
			Usage:     "2-character ISO Country code to filter to one specific country",
			QueryPath: "countrycode",
		},
		&requestflag.Flag[[]string]{
			Name:      "countryid",
			Usage:     "Exact match on a given numeric country id (comma separated list)",
			QueryPath: "countryid",
		},
		&requestflag.Flag[[]any]{
			Name:      "dataproviderid",
			Usage:     "Exact match on a given data provider id id (comma separated list). ",
			QueryPath: "dataproviderid",
		},
		&requestflag.Flag[float64]{
			Name:      "distance",
			Usage:     "Optionally filter results by a max distance from the given latitude/longitude",
			QueryPath: "distance",
		},
		&requestflag.Flag[string]{
			Name:      "distanceunit",
			Usage:     "`miles` or `km` distance unit",
			Default:   "Miles",
			QueryPath: "distanceunit",
		},
		&requestflag.Flag[string]{
			Name:      "greaterthanid",
			Usage:     "Filter to items with ID greater than given value",
			QueryPath: "greaterthanid",
		},
		&requestflag.Flag[bool]{
			Name:      "includecomments",
			Usage:     "If true, user comments and media items will be include in result set",
			QueryPath: "includecomments",
		},
		&requestflag.Flag[int64]{
			Name:      "latitude",
			Usage:     "Latitude for distance calculation and filtering",
			QueryPath: "latitude",
		},
		&requestflag.Flag[[]any]{
			Name:      "levelid",
			Usage:     "Exact match on a given charging level (1-3) id (comma separated list)",
			QueryPath: "levelid",
		},
		&requestflag.Flag[float64]{
			Name:      "longitude",
			Usage:     "Longitude for distance calculation and filtering",
			QueryPath: "longitude",
		},
		&requestflag.Flag[int64]{
			Name:      "maxresults",
			Usage:     "Limit on max number of results returned",
			Default:   100,
			QueryPath: "maxresults",
		},
		&requestflag.Flag[string]{
			Name:      "modifiedsince",
			Usage:     "Filter to results modified after the given date",
			QueryPath: "modifiedsince",
		},
		&requestflag.Flag[bool]{
			Name:      "opendata",
			Usage:     `Use opendata=true for only OCM provided ("Open") data.`,
			QueryPath: "opendata",
		},
		&requestflag.Flag[[]any]{
			Name:      "operatorid",
			Usage:     "Exact match on a given EVSE operator id (comma separated list)",
			QueryPath: "operatorid",
		},
		&requestflag.Flag[string]{
			Name:      "output",
			Usage:     "Optional output format `json`,`geojson`,`xml`,`csv`, JSON is the default and recommended as the highest fidelity.",
			Default:   "json",
			QueryPath: "output",
		},
		&requestflag.Flag[string]{
			Name:      "polygon",
			Usage:     "Filter results within a given Polygon. Specify an encoded polyline for the polygon shape. Polygon will be automatically closed from the last point to the first point.",
			QueryPath: "polygon",
		},
		&requestflag.Flag[string]{
			Name:      "polyline",
			Usage:     "Filter results along an encoded polyline, use with distance param to increase search distance along line. Polyline is expanded into a polygon to cover the search distance.",
			QueryPath: "polyline",
		},
		&requestflag.Flag[string]{
			Name:      "sortby",
			Usage:     "Default sort order is based on spatial index but you can optionally sort by  `modified_asc` for results in order of modification (oldest to newest), or ` id_asc` for results in order of ID",
			QueryPath: "sortby",
		},
		&requestflag.Flag[[]any]{
			Name:      "statustypeid",
			Usage:     "Exact match on a given status type id (comma separated list)",
			QueryPath: "statustypeid",
		},
		&requestflag.Flag[[]any]{
			Name:      "usagetypeid",
			Usage:     "Exact match on a given usage type id (comma separated list)",
			QueryPath: "usagetypeid",
		},
		&requestflag.Flag[bool]{
			Name:      "verbose",
			Usage:     "Set to false to get a smaller result set with null items removed.",
			QueryPath: "verbose",
		},
	},
	Action:          handlePoiList,
	HideHelpCommand: true,
}

func handlePoiList(ctx context.Context, cmd *cli.Command) error {
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

	params := ocm.PoiListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Poi.List(ctx, params, options...)
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
		Title:          "poi list",
		Transform:      transform,
	})
}
