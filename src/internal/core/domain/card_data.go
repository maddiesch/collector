package domain

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"

	"cloud.google.com/go/civil"
)

type CardData struct {
	ID              string        `json:"id"`
	TCGPlayerID     int64         `json:"tcgplayer_id"`
	Name            string        `json:"name"`
	SetCode         string        `json:"set"`
	SetName         string        `json:"set_name"`
	LanguageTag     string        `json:"lang"`
	CollectorNumber string        `json:"collector_number"`
	ReleasedAt      ReleaseDate   `json:"released_at"`
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
	NormalUSD int
	FoilUSD   int
}

func (c *CardDataPrice) UnmarshalJSON(data []byte) error {
	price := struct {
		USD     string `json:"usd"`
		USDFoil string `json:"usd_foil"`
	}{}
	if err := json.Unmarshal(data, &price); err != nil {
		return err
	}

	parse := func(s string) (int, error) {
		value, err := strconv.ParseFloat(price.USD, 32)
		if err != nil {
			return 0, err
		}
		return int(value * 100), nil
	}

	c.NormalUSD, _ = parse(price.USD)
	c.FoilUSD, _ = parse(price.USDFoil)

	return nil
}

type ReleaseDate civil.Date

func (r *ReleaseDate) UnmarshalJSON(data []byte) error {
	var d civil.Date

	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	*r = ReleaseDate(d)

	return nil
}

func (r *ReleaseDate) Scan(src any) error {
	srcStr, ok := src.(string)
	if !ok {
		return driver.ErrBadConn
	}

	d, err := civil.ParseDate(srcStr)
	if err != nil {
		return err
	}

	*r = ReleaseDate(d)

	return nil
}

func (r ReleaseDate) Value() (driver.Value, error) {
	return civil.Date(r).String(), nil
}
