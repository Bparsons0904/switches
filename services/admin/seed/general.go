package seed

import (
	"time"

	"gorm.io/gorm"
)

func seedTypes(tx *gorm.DB) error { // {{{
	type Type struct {
		ID        int       `gorm:"primaryKey;autoIncrement"              json:"id"`
		Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
		Code      string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"`
		Category  string    `gorm:"type:varchar(50);not null;index"       json:"category"`
		CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
	}

	types := []Type{
		// Switch Types
		{Name: "Linear", Code: "linear", Category: "switch_type"},
		{Name: "Tactile", Code: "tactile", Category: "switch_type"},
		{Name: "Clicky", Code: "clicky", Category: "switch_type"},
		{Name: "Silent Linear", Code: "silent_linear", Category: "switch_type"},
		{Name: "Silent Tactile", Code: "silent_tactile", Category: "switch_type"},

		// Spring Types
		{Name: "Linear", Code: "spring_linear", Category: "spring_type"},
		{Name: "Progressive", Code: "progressive", Category: "spring_type"},
		{Name: "Slow", Code: "slow", Category: "spring_type"},
		{Name: "Fast", Code: "fast", Category: "spring_type"},
		{Name: "Variable", Code: "variable", Category: "spring_type"},
		{Name: "Balanced", Code: "balanced", Category: "spring_type"},
		{Name: "Soft", Code: "soft", Category: "spring_type"},
		{Name: "Heavy", Code: "heavy", Category: "spring_type"},
		{Name: "Custom", Code: "custom", Category: "spring_type"},
		{Name: "Dual-Stage", Code: "dual_stage", Category: "spring_type"},
		{Name: "Triple-Stage", Code: "triple_stage", Category: "spring_type"},

		// Spring Materials
		{Name: "Steel", Code: "steel", Category: "spring_material"},
		{
			Name:     "Gold-Plated Steel",
			Code:     "gold_plated_steel",
			Category: "spring_material",
		},
		{Name: "Copper", Code: "copper", Category: "spring_material"},
		{Name: "Phosphor Bronze", Code: "phosphor_bronze", Category: "spring_material"},
		{Name: "Stainless Steel", Code: "stainless_steel", Category: "spring_material"},

		// Plastic Materials
		{Name: "ABS", Code: "abs", Category: "material"},
		{Name: "POM", Code: "pom", Category: "material"},
		{Name: "Nylon", Code: "nylon", Category: "material"},
		{Name: "UHMWPE", Code: "uhmwpe", Category: "material"},
		{Name: "Polycarbonate", Code: "pc", Category: "material"},
		{Name: "PBT", Code: "pbt", Category: "material"},

		// Sound Levels
		{Name: "Very Low", Code: "very_low", Category: "sound_level"},
		{Name: "Low", Code: "low", Category: "sound_level"},
		{Name: "Medium", Code: "medium", Category: "sound_level"},
		{Name: "High", Code: "high", Category: "sound_level"},
		{Name: "Very High", Code: "very_high", Category: "sound_level"},

		// Sound Types
		{Name: "Clicky", Code: "sound_clicky", Category: "sound_type"},
		{Name: "Thocky", Code: "sound_thocky", Category: "sound_type"},
		{Name: "Quiet", Code: "sound_quiet", Category: "sound_type"},
		{Name: "Creamy", Code: "sound_creamy", Category: "sound_type"},
		{Name: "Clacky", Code: "sound_clacky", Category: "sound_type"},

		// Tactility Types
		{Name: "Leaf Spring", Code: "leaf_spring", Category: "tactility_type"},
		{Name: "Coil Spring", Code: "coil_spring", Category: "tactility_type"},
		{Name: "Click Bar", Code: "click_bar", Category: "tactility_type"},

		// Tactility Feedback
		{Name: "Sharp", Code: "sharp", Category: "tactility_feedback"},
		{Name: "Rounded", Code: "rounded", Category: "tactility_feedback"},
		{Name: "Crisp", Code: "crisp", Category: "tactility_feedback"},
		{Name: "Smooth", Code: "smooth", Category: "tactility_feedback"},

		// Colors
		{Name: "Red", Code: "red", Category: "color"},
		{Name: "Black", Code: "black", Category: "color"},
		{Name: "Blue", Code: "blue", Category: "color"},
		{Name: "Brown", Code: "brown", Category: "color"},
		{Name: "Clear", Code: "clear", Category: "color"},
		{Name: "Yellow", Code: "yellow", Category: "color"},
		{Name: "White", Code: "white", Category: "color"},
		{Name: "Transparent", Code: "transparent", Category: "color"},
		{Name: "Smokey", Code: "smokey", Category: "color"},
		{Name: "Green", Code: "green", Category: "color"},
		{Name: "Purple", Code: "purple", Category: "color"},
		{Name: "Orange", Code: "orange", Category: "color"},
		{Name: "Pink", Code: "pink", Category: "color"},
		{Name: "Gray", Code: "gray", Category: "color"},
		{Name: "Silver", Code: "silver", Category: "color"},
		{Name: "Gold", Code: "gold", Category: "color"},
		{Name: "Turquoise", Code: "turquoise", Category: "color"},
		{Name: "Teal", Code: "teal", Category: "color"},
		{Name: "Lavender", Code: "lavender", Category: "color"},
		{Name: "Magenta", Code: "magenta", Category: "color"},
		{Name: "Cyan", Code: "cyan", Category: "color"},
		{Name: "Ivory", Code: "ivory", Category: "color"},
		{Name: "Coral", Code: "coral", Category: "color"},
		{Name: "Maroon", Code: "maroon", Category: "color"},
		{Name: "Beige", Code: "beige", Category: "color"},
		{Name: "Mint", Code: "mint", Category: "color"},
		{Name: "Peach", Code: "peach", Category: "color"},
		{Name: "Tan", Code: "tan", Category: "color"},
		{Name: "Khaki", Code: "khaki", Category: "color"},

		// Pin Configuration
		{Name: "3-Pin", Code: "3_pin", Category: "pin_configuration"},
		{Name: "5-Pin", Code: "5_pin", Category: "pin_configuration"},
	}

	if err := tx.Create(&types).Error; err != nil {
		return err
	}

	return nil
} // }}}

func seedProducers(tx *gorm.DB) error { // {{{
	type Producer struct {
		ID        int       `gorm:"type:int;primaryKey;autoIncrement"     json:"id"`
		Name      string    `gorm:"type:varchar(50);not null"             json:"name"`
		Alias     string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"alias"`
		SiteURL   string    `gorm:"type:varchar(50)"                      json:"siteURL,omitempty"`
		CreatedAt time.Time `gorm:"autoCreateTime"                        json:"createdAt"`
		UpdateAt  time.Time `gorm:"autoUpdateTime"                        json:"updatedAt"`
	}
	producers := []Producer{
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
			Name:    "Kinetic Labs",
			Alias:   "kintic_labs",
			SiteURL: "https://kineticlabs.com/",
		},
		{
			Name:    "Chosfox",
			Alias:   "chosfox",
			SiteURL: "https://chosfox.com/",
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
			Name:    "Wuque Studio",
			Alias:   "wuque_studio",
			SiteURL: "https://shop.wuquestudio.com/",
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
		{
			Name:    "Keydous",
			Alias:   "keydous",
			SiteURL: "https://www.keydous.com/",
		},
		{
			Name:    "Cannon Keys",
			Alias:   "cannon_keys",
			SiteURL: "https://cannonkeys.com/",
		},
		{
			Name:    "KBDFans",
			Alias:   "kbd_fans",
			SiteURL: "https://kbdfans.com/",
		},
		{
			Name:    "HMX",
			Alias:   "hmx",
			SiteURL: "",
		},
		{
			Name:    "Tbcats Studio",
			Alias:   "tbcats",
			SiteURL: "",
		},
		{
			Name:    "C³ EQUALZ X TKC",
			Alias:   "c3xtkc",
			SiteURL: "https://c3equalz.com/",
		},
	}

	if err := tx.Create(&producers).Error; err != nil {
		return err
	}

	return nil
} // }}}
