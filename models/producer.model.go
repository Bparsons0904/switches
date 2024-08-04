package models

import "github.com/google/uuid"

type Producer struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7();primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(50);not null"                       json:"name"`
	Alias     string    `gorm:"type:varchar(50);not null;uniqueIndex"           json:"alias"`
	SiteURL   string    `gorm:"type:varchar(50)"                                json:"siteURL,omitempty"`
	CreatedAt string    `gorm:"autoCreateTime"                                  json:"createdAt"`
	UpdateAt  string    `gorm:"autoUpdateTime"                                  json:"updatedAt"`
}

var producers = []Producer{
	{
		Name:    "Cherry",
		Alias:   "cherry",
		SiteURL: "https://www.cherrymx.de/",
	},
	{
		Name:    "Gateron",
		Alias:   "gateron",
		SiteURL: "https://www.gateron.com/",
	},
	{
		Name:    "Kailh",
		Alias:   "kailh",
		SiteURL: "https://www.kailh.com/",
	},
	{
		Name:    "Outemu",
		Alias:   "outemu",
		SiteURL: "",
	},
	{
		Name:    "ZealPC",
		Alias:   "zealpc",
		SiteURL: "https://zealpc.net/",
	},
	{
		Name:    "TTC",
		Alias:   "ttc",
		SiteURL: "",
	},
	{
		Name:    "Durock",
		Alias:   "durock",
		SiteURL: "https://www.durock.us/",
	},
	{
		Name:    "Akko",
		Alias:   "akko",
		SiteURL: "https://en.akkogear.com/",
	},
	{
		Name:    "NovelKeys",
		Alias:   "novelkeys",
		SiteURL: "https://novelkeys.com/",
	},
	{
		Name:    "Drop",
		Alias:   "drop",
		SiteURL: "https://drop.com/",
	},
	{
		Name:    "Glorious",
		Alias:   "glorious",
		SiteURL: "https://www.gloriousgaming.com/",
	},
	{
		Name:    "Keychron",
		Alias:   "keychron",
		SiteURL: "https://www.keychron.com/",
	},
	{
		Name:    "Epomaker",
		Alias:   "epomaker",
		SiteURL: "https://epomaker.com/",
	},
	{
		Name:    "Razer",
		Alias:   "razer",
		SiteURL: "https://www.razer.com/",
	},
	{
		Name:    "Logitech",
		Alias:   "logitech",
		SiteURL: "https://www.logitechg.com/",
	},
	{
		Name:    "SteelSeries",
		Alias:   "steelseries",
		SiteURL: "https://steelseries.com/",
	},
	{
		Name:    "Varmilo",
		Alias:   "varmilo",
		SiteURL: "https://en.varmilo.com/",
	},
	{
		Name:    "Ducky",
		Alias:   "ducky",
		SiteURL: "https://www.duckychannel.com.tw/",
	},
	{
		Name:    "Roccat",
		Alias:   "roccat",
		SiteURL: "https://en.roccat.com/",
	},
	{
		Name:    "Cooler Master",
		Alias:   "cooler_master",
		SiteURL: "https://www.coolermaster.com/",
	},
	{
		Name:    "AKKO X Varmilo",
		Alias:   "akko_varmilo",
		SiteURL: "https://en.akkogear.com/collections/varmilo/",
	},
	{
		Name:    "JWK",
		Alias:   "jwk",
		SiteURL: "http://www.jwk-tech.com/",
	},
	{
		Name:    "Tecsee",
		Alias:   "tecsee",
		SiteURL: "http://www.tecsee.com/",
	},
	{
		Name:    "Everglide",
		Alias:   "everglide",
		SiteURL: "https://everglide.com/",
	},
	{
		Name:    "SP-Star",
		Alias:   "sp_star",
		SiteURL: "",
	},
	{
		Name:    "Gazzew",
		Alias:   "gazzew",
		SiteURL: "",
	},
}
