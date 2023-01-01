package gnomelook

type DataTheme struct {
	Url     string  `json:"url"`
	Sources Sources `json:"sources"`
	IdWeb   string  `json:"id_web"`
}

type Sources struct {
	Repository string      `json:"repository"`
	LastCommit interface{} `json:"last_commit"`
	ZipURL     string      `json:"zip_url"`
}

type KEY string

var DATA_THEMES map[KEY]DataTheme
