package assettypes

import "github.com/goledgerdev/cc-tools/assets"

var Sale12 = assets.AssetType{
	Tag:         "sale12",
	Label:       "Sale12",
	Description: "Defines a sale between org1 and org2",
	Readers:     []string{"org1MSP", "org2MSP"},

	Props: []assets.AssetProp{
		{
			IsKey:    true,
			Tag:      "boughtWork",
			Label:    "Bought Work",
			DataType: "->work",
			Writers:  []string{"org1MSP", "org2MSP"},
		},
		{
			Required: true,
			Tag:      "prevOwner",
			Label:    "Owner before sale",
			DataType: "->museum",
		},
		{
			Required: true,
			Tag:      "price",
			Label:    "Price",
			DataType: "number",
		},
	},
}
