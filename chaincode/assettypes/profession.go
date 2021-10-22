package assettypes

import (
	"github.com/goledgerdev/cc-tools-demo/chaincode/validators"
	"github.com/goledgerdev/cc-tools/assets"
)

var Profession = assets.AssetType{
	Tag:         "profession",
	Label:       "Profession",
	Description: "Job Especifications",

	Props: []assets.AssetProp{
		// {
		// 	// TODO: Ver se tem como fazer isso bonito, ou usar nome como key
		// 	IsKey:    true,
		// 	Required: true,
		// 	Tag:      "id",
		// 	Label:    "Profession identicator",
		// 	DataType: "id",
		// },
		{
			// Nome
			Tag:      "name",
			IsKey:    true,
			Required: true,
			Label:    "Profession Name",
			DataType: "string",
			Validate: validators.StringNotNull("profession name"),
		},
	},
}
