package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	"github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

var CreateNewMuseum = tx.Transaction{
	Tag:         "createNewMuseum",
	Label:       "Create New Museum",
	Description: "Creates a new Museum",
	Method:      "POST",

	Args: tx.ArgList{
		{
			Tag:         "name",
			Label:       "Name",
			Description: "Name of the Museum",
			DataType:    "string",
			Required:    true,
		},
	},

	Routine: func(sw *stubwrapper.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["name"].(string)

		museumMap := make(map[string]interface{})
		// TODO: Entender o que é esse @
		museumMap["@assetType"] = "museum"
		museumMap["name"] = name

		// Create asset from interface map
		museumAsset, err := assets.NewAsset(museumMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new museum on channel
		_, err = museumAsset.PutNew(sw)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		museumJSON, nerr := json.Marshal(museumAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return museumJSON, nil
	},
}
