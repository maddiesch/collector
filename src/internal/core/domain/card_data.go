package domain

import (
	"encoding/json"
	"strconv"
)

type CardData struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	SetName         string        `json:"set_name"`
	Language        string        `json:"lang"`
	CollectorNumber string        `json:"collector_number"`
	ReleasedAt      string        `json:"released_at"`
	ManaCost        string        `json:"mana_cost"`
	HasFoil         bool          `json:"foil"`
	HasNormal       bool          `json:"nonfoil"`
	Image           CardDataImage `json:"image_uris"`
	Links           CardDataLink  `json:"related_uris"`
	Prices          CardDataPrice `json:"prices"`
}

type CardDataImage struct {
	Small  string `json:"small"`
	Normal string `json:"normal"`
}

type CardDataLink struct {
	Gatherer string `json:"gatherer"`
}

type CardDataPrice struct {
	NormalUSD int32
	FoilUSD   int32
}

func (c *CardDataPrice) UnmarshalJSON(data []byte) error {
	price := struct {
		USD     string `json:"usd"`
		USDFoil string `json:"usd_foil"`
	}{}
	if err := json.Unmarshal(data, &price); err != nil {
		return err
	}

	parse := func(s string) (int32, error) {
		value, err := strconv.ParseFloat(price.USD, 32)
		if err != nil {
			return 0, err
		}
		return int32(value * 100), nil
	}

	c.NormalUSD, _ = parse(price.USD)
	c.FoilUSD, _ = parse(price.USDFoil)

	return nil
}
