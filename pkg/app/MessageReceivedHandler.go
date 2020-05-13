package app

import (
	"fmt"
	"github.com/paked/messenger"
	"log"
	"strings"
	"time"
)

func (app *Application) InitMessageReceivedHandler(){
	app.Msng.HandleMessage(func(m messenger.Message, r *messenger.Response) {
		fmt.Printf("%v (Sent, %v)\n", m.Text, m.Time.Format(time.UnixDate))

		p, err := app.Msng.ProfileByID(m.Sender.ID, []string{"name", "first_name", "last_name", "profile_pic"})
		if err != nil {
			fmt.Println("Something went wrong!", err)
		}
		log.Println(p)
		//r.Text(fmt.Sprintf("Hello, %v!", p.FirstName), messenger.ResponseType)

		queryResult, err := app.Df.DetectIntentText(string(m.Sender.ID),m.Text)
		if err != nil {
			log.Fatal(err.Error())
		}

		if strings.Contains(queryResult.Action,"smalltalk") {
			err := r.Text(queryResult.FulfillmentText,messenger.ResponseType)
			if err != nil {
				log.Println(err)
				return // if there is an error, resp is empty struct, useless
			}
			//log.Println("Message ID", resp.MessageID, "sent to user", resp.RecipientID)
		} else {
			err := r.Text(m.Text,messenger.ResponseType) // echo, send back to user the same text he sent to us
			if err != nil {
				log.Println(err)
				return // if there is an error, resp is empty struct, useless
			}
			//log.Println("Message ID", resp.MessageID, "sent to user", resp.RecipientID)
		}
	})
}