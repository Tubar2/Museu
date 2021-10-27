package txdefs

import (
	"encoding/json"
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	"github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

func getProfessions(sw *stubwrapper.StubWrapper, profession assets.Key) ([]interface{}, error) {
	query := fmt.Sprintf(`{
		"selector": {
			"profession.@key": "%s"
		}
	}`, profession.Key())

	iterator, err := sw.GetQueryResult(query)
	if err != nil {
		return nil, errors.WrapErrorWithStatus(err, "error getting query result", 500)
	}
	var response []interface{}

	for iterator.HasNext() {
		res, err := iterator.Next()
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error iterating response", 500)
		}

		professionRes := make(map[string]interface{})

		err = json.Unmarshal(res.Value, &professionRes)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error getting query result", 500)
		}
		response = append(response, professionRes)
	}

	return response, nil
}

var GetArtistByProfession = tx.Transaction{
	Tag:         "getArtistByProfession",
	Label:       "get a artist by Profession",
	Description: "Return a list a list of artist by profession",
	ReadOnly:    true,

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "profession",
			Label:    "Profession reference",
			DataType: "->profession",
		},
	},

	Routine: func(sw *stubwrapper.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		profession, _ := req["profession"].(assets.Key)

		res, err := getProfessions(sw, profession)
		if err != nil {
			return nil, errors.WrapError(err, "couldn't get artists with profession")
		}

		responseJSON, nerr := json.Marshal(res)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return responseJSON, nil
	},
}
