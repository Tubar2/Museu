package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	"github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

var UpdateWorkArtist = tx.Transaction{
	Tag:         "updateWorkArtist",
	Label:       "Updates a work's artist",
	Description: "Changes a work's artist",
	Method:      "PUT",

	Args: tx.ArgList{
		{
			Tag:         "work",
			Label:       "Work",
			Description: "Work",
			DataType:    "->work",
			Required:    true,
		},
		{
			Tag:         "artist",
			Label:       "Artist",
			Description: "New artist of the work",
			DataType:    "->artist",
		},
	},

	Routine: func(sw *stubwrapper.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		workKey, ok := req["work"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter book must be an asset")
		}
		artistKey, ok := req["artist"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter tenant must be an asset")
		}

		// Returns work-assets from channel
		workAsset, err := workKey.Get(sw)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		// TODO: o que é isso: (*workAsset)
		workMap := (map[string]interface{})(*workAsset)

		// Returns artist-asset from channel
		artistAsset, err := artistKey.Get(sw)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}
		artistMap := (map[string]interface{})(*artistAsset)

		updatedArtistKey := make(map[string]interface{})
		updatedArtistKey["@assetType"] = "artist"
		updatedArtistKey["@key"] = artistMap["@key"]

		// Update work's artist
		workMap["artist"] = updatedArtistKey
		workMap, err = workAsset.Update(sw, workMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update asset")
		}

		// Marshal asset back to JSON format
		workJSON, nerr := json.Marshal(workMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return workJSON, nil
	},
}
