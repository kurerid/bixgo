package bixgo

import "encoding/json"

type BitrixBoolean bool

func (b *BitrixBoolean) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*b = s == "Y" || s == "y"
	return nil
}
