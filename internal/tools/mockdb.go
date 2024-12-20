package tools

import (
	"time"
)

type mockDb struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"karthikj": {
		AuthToken: "123abc",
		Username:  "karthikj",
	},
	"jasonstatham": {
		AuthToken: "456def",
		Username:  "jasonstatham",
	},
	"paulwalker": {
		AuthToken: "789ghi",
		Username:  "paulwalker",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"karthikj": {
		Coins:    5000,
		Username: "karthikj",
	},
	"jasonstatham": {
		Coins:    199,
		Username: "jasonstatham",
	},
	"paulwalker": {
		Coins:    599,
		Username: "paulwalker",
	},
}

func (d *mockDb) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}
func (d *mockDb) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}
	
func (d *mockDb) SetupDatabase() error {
	return nil
}
