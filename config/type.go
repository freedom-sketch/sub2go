package config

var defaultConfig = Config{
	API:          API{},
	Panel:        PanelCredentials{},
	Subscription: SubscriptionMetaData{},
	Logging:      Logging{},
}

type PanelCredentials struct {
	URL      string `json:"url"`
	Username int    `json:"username"`
	Password string `json:"password"`
}

type SubscriptionMetaData struct {
	ProfileUpdateInterval int    `json:"profile-update-interval"`
	ProfileTitle          string `json:"profile-title"`
	SubscriptionUserinfo  string `json:"subscription-userinfo"`
	SupportUrl            string `json:"support-url"`
	ProfileWebPageUrl     string `json:"profile-web-page-url"`
	Announce              string `json:"announce"`
}

type API struct {
	Host    string `json:"host"`
	WebPath string `json:"web_path"`
}

type Logging struct {
	FileName string `json:"file_name"`
	Level    string `json:"level"`
}

type Config struct {
	API          API                  `json:"api"`
	Panel        PanelCredentials     `json:"panel"`
	Subscription SubscriptionMetaData `json:"subscription"`
	Logging      Logging              `json:"logging"`
}
