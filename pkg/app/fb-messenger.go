package app

import (
	messenger "github.com/paked/messenger"
	"messengerBot/pkg/config"
	"messengerBot/pkg/nlp"
)

type Application struct {
	Msng  *messenger.Messenger
	Df *nlp.DialogflowProcessor
}

const (
	language   = "it"
	errorMessage = "Non so rispondere alla tua domanda"
	timezone = "Europe/Madrid"
)


func Init(cfg config.Config) (*Application, error) {

	client := messenger.New(messenger.Options{
		Verify: false,
		AppSecret: cfg.GetPageID(),
		VerifyToken: cfg.GetVerifyToken(),
		Token: cfg.GetAccessToken(),
	})

	df := nlp.Init("pizzabot-hqfwxu","pizzabot-hqfwxu-42470a011a39.json","it","Europe/Madrid")
	//defer df.Close()
	return &Application{
		Msng: client,
		Df: df,
	}, nil


}






/*func (app *Application) InitPostbackReceivedHandler(f func(msng *messenger.Messenger, userID int64, m messenger.FacebookPostback) ){
	app.Msng.PostbackReceived = f
}
func (app *Application) InitDeliveryReceivedHandler(f func(msng *messenger.Messenger, userID int64, m messenger.FacebookDelivery) ){
	app.Msng.DeliveryReceived = f
}*/

