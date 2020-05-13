package nlp

import (
	"context"
	"fmt"
	"log"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

type DialogflowProcessor struct {
	projectID        string
	authJSONFilePath string
	lang             string
	timeZone         string
	sessionClient    *dialogflow.SessionsClient
	ctx              context.Context
}

func Init(a ...string) *DialogflowProcessor {
	if a[0] == "" || a[1] == "" {
		log.Fatalf("Received empty project (%s) or JSON File path (%s)", a[0], a[1])
	}

	// Auth process: https://dialogflow.com/docs/reference/v2-auth-setup

	ctx := context.Background()
	sessionClient, err := dialogflow.NewSessionsClient(ctx, option.WithCredentialsFile(a[1]))
	if err != nil {
		log.Fatal("Error in auth with Dialogflow")
	}
	//defer sessionClient.Close()


	return &DialogflowProcessor{
		projectID:        a[0],
		authJSONFilePath: a[1],
		lang:             a[2],
		timeZone:         a[3],
		sessionClient:    sessionClient,
		ctx:              ctx,
	}

}
func(dp *DialogflowProcessor) DetectIntentText(sessionID, text string) (*dialogflowpb.QueryResult, error) {


	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", dp.projectID, sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: dp.lang}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := dp.sessionClient.DetectIntent(dp.ctx, &request)
	if err != nil {
		log.Fatal(err.Error())
	}

	queryResult := response.GetQueryResult()

	return queryResult, nil
}

func (dp *DialogflowProcessor) Close (){
	dp.sessionClient.Close()
}