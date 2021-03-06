package fbapi

type Api struct {
	AccessToken string
	ExpiresIn   int
}

type TokenData struct {
	ClientId     int64
	ClientSecret string
	Code         string
	RedirectUri  string
}

type GetTokenAns struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type GetAccountAns struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	AccessToken string `json:"access_token"`
}

type AdsAns struct {
	Data   []AdsAnsData `json:"data"`
	Paging Paging       `json:"paging"`
}

type Paging struct {
	Cursors  Cursors `json:"cursors"`
	Previous string  `json:"previous"`
	Next     string  `json:"next"`
}

type Cursors struct {
	After  string `json:"after"`
	Before string `json:"before"`
}

type AdsAnsData struct {
	ID          string                `json:"id"`
	Insights    AdsAnsDataInsights    `json:"insights"`
	Adcreatives AdsAnsDataAdcreatives `json:"adcreatives"`
}

type AdsAnsDataInsights struct {
	Data []AdsAnsDataInsightsData `json:"data"`
}

type AdsAnsDataInsightsData struct {
	DateStart   string `json:"date_start"`
	DateStop    string `json:"date_stop"`
	Impressions string `json:"impressions"`
	Spend       string `json:"spend"`
	AccountId   string `json:"account_id"`
	CampaignId  string `json:"campaign_id"`
	AdsetId     string `json:"adset_id"`
	AdId        string `json:"ad_id"`
}

type InsightsAns struct {
	Data []InsightsAnsData `json:"data"`
}

type InsightsAnsData struct {
	AdID        string `json:"ad_id"`
	Impressions string `json:"impressions"`
	Clicks      string `json:"clicks"`
	Spend       string `json:"spend"`
	DateStart   string `json:"date_start"`
	DateStop    string `json:"date_stop"`
}

type AdsAnsDataAdcreatives struct {
	Data []AdsAnsDataAdcreativesData `json:"data"`
}

type AdsAnsDataAdcreativesData struct {
	ObjectStorySpec AdsAnsDataAdcreativesDataObjectStorySpec `json:"object_story_spec"`
	Id              string                                   `json:"id"`
	Name            string                                   `json:"name"`
}

type AdsAnsDataAdcreativesDataObjectStorySpec struct {
	PageId   string                                           `json:"page_id"`
	LinkData AdsAnsDataAdcreativesDataObjectStorySpecLinkData `json:"link_data"`
}

type AdsAnsDataAdcreativesDataObjectStorySpecLinkData struct {
	Link    string `json:"link"`
	Message string `json:"message"`
	Name    string `json:"name"`
}
