package assettypes

import "github.com/goledgerdev/cc-tools/assets"

var Work = assets.AssetType{
	Tag:         "work",
	Label:       "Work",
	Description: "Work is a object made by a artist",

	Props: []assets.AssetProp{
		{
			Tag:      "name",
			Label:    "Work Name",
			Required: true,
			IsKey:    true,
			DataType: "string",
		},
		{
			Tag:      "artist",
			Label:    "Artist",
			Required: true,
			IsKey:    true,
			DataType: "->artist",
		},
		{
			Tag:      "created-at",
			Label:    "When it was created",
			DataType: "datetime",
		},
	},
}
