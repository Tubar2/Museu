package assettypes

import "github.com/goledgerdev/cc-tools/assets"

var Museum = assets.AssetType{
	Tag:         "museum",
	Label:       "Museum",
	Description: "Museum is a collection of works",

	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "name",
			Label:    "Museum's name",
			DataType: "string",
		},
		{
			Tag:      "works",
			Label:    "All works in museum",
			DataType: "[]->work",
		},
	},
}
