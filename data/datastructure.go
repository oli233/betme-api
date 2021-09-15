package data

type Data struct {
	Key           string `json:"key"`
	Active        bool   `json:"active"`
	Group         string `json:"group"`
	Details       string `json:"details"`
	Title         string `json:"title"`
	Has_outrights bool   `json:"has_outrights"`
}

type Sports struct {
	Success   bool   `json:"success"`
	DataArray []Data `json:"data"`
}

// H2h OddsData Odds structures
type H2h struct {
	H2hData []float32 `json:"h2h"`
}

type Site struct {
	SiteKey    string  `json:"site_key"`
	SiteNice   string  `json:"site_nice"`
	LastUpdate float32 `json:"last_update"`
	Odd        H2h     `json:"odds"`
}

type OddsData struct {
	Id           string   `json:"id"`
	SportKey     string   `json:"sport_key"`
	SportNice    string   `json:"sport_nice"`
	Teams        []string `json:"teams"`
	CommenceTime float32  `json:"commence_time"`
	HomeTeam     string   `json:"home_team"`
	Sites        []Site   `json:"sites"`
	SitesCount   int      `json:"sites_count"`
}

type Odds struct {
	Success       bool       `json:"success"`
	OddsDataArray []OddsData `json:"data"`
}
