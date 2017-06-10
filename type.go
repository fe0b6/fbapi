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
