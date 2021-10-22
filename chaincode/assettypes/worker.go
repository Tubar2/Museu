package assettypes

import (
	"github.com/goledgerdev/cc-tools-demo/chaincode/validators"
	"github.com/goledgerdev/cc-tools/assets"
)

var Artist = assets.AssetType{
	Tag:         "artist",
	Label:       "Artist",
	Description: "Personal and Professional information about someone",

	Props: []assets.AssetProp{
		{
			// Cpf
			Tag:      "cpf",
			Required: true,
			Label:    "CPF",
			DataType: "cpf",
			IsKey:    true,
		},
		{
			// Nome
			Tag:      "name",
			Required: true,
			Label:    "Artist's name",
			DataType: "string",
			Validate: validators.StringNotNull("artist name"),
		},
		{
			// Asset's profession reference
			Tag:      "profession",
			Label:    "Artist's profession",
			DataType: "->profession",
		},
	},
}
