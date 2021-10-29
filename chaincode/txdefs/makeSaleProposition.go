package txdefs

import (
	"encoding/json"
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	"github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

func getMuseumWithWork(sw *stubwrapper.StubWrapper, work assets.Key) (map[string]interface{}, error) {
	query := fmt.Sprintf(`{
		"selector": {
			"works": {
				"$elemMatch": {
					"@key": "%s"
				}
			}
		}
		}`, work.Key())

	iterator, err_ := sw.GetQueryResult(query)
	if err_ != nil {
		return nil, errors.WrapErrorWithStatus(err_, "error getting query result", 500)
	}

	res, err := iterator.Next()
	if err != nil {
		return nil, errors.WrapErrorWithStatus(err, "error iterating response", 500)
	}
	response := make(map[string]interface{})
	err = json.Unmarshal(res.Value, &response)
	if err != nil {
		return nil, errors.WrapErrorWithStatus(err, "error during unmarshal of response", 500)
	}

	return response, nil
}

var MakeSaleProposition = tx.Transaction{
	Tag:         "makeSaleProposition",
	Label:       "Make a Sale Proposition",
	Description: "Makes a sale proposition of a work",
	Method:      "POST",

	Args: tx.ArgList{
		{
			Tag:         "work",
			Label:       "Work",
			Description: "Work you wish to make a proposition towards",
			DataType:    "->work",
			Required:    true,
		},
		{
			Tag:         "price",
			Label:       "Price",
			Description: "Price proposition",
			DataType:    "number",
			Required:    true,
		},
	},

	Routine: func(sw *stubwrapper.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		work, _ := req["work"].(assets.Key)
		price, _ := req["price"].(float64)

		res, err := getMuseumWithWork(sw, work)
		if err != nil {
			return nil, errors.WrapError(err, "couldn't get museum with work")
		}

		museumKey, err := assets.NewKey(res)

		saleMap := make(map[string]interface{})

		saleMap["@assetType"] = "sale12"
		saleMap["boughtWork"] = work
		saleMap["prevOwner"] = museumKey
		saleMap["price"] = price

		saleAsset, err := assets.NewAsset(saleMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = saleAsset.PutNew(sw)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		responseJSON, nerr := json.Marshal(saleAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return responseJSON, nil
	},
}
