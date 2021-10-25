package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	"github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// TODO: Perguntar como fazer pra pegar todas as obras do artista
// sendo que artista n tem referênia pra obra
var GetWorksFromMuseum = tx.Transaction{
	Tag:         "getWorksFromMuseum",
	Label:       "Get Works From Museum",
	Description: "Get list of works of a specific museum",
	Method:      "GET",

	Args: []tx.Argument{
		{
			Tag:         "museum",
			Label:       "Museum",
			Description: "Museum name",
			DataType:    "->museum",
			Required:    true,
		},
	},

	Routine: func(sw *stubwrapper.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		// Get museum name from arguments
		museumKey, ok := req["museum"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter museum must be an asset")
		}

		museumMap, err := museumKey.GetMap(sw)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}

		obras, ok := museumMap["works"].([]assets.AssetType)
		if !ok {
			return nil, errors.WrapError(nil, "Unable to fetch works")
		}
		returnMap := make(map[string]interface{})
		returnMap["obras"] = obras

		returnJSON, nerr := json.Marshal(returnMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return returnJSON, nil
	},
}
