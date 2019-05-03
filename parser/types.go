package parser

type TelegramWebhoook struct {
	ResponseID  string `json:"responseId"`
	QueryResult struct {
		QueryText                string      `json:"queryText"`
		Parameters               interface{} `json:"parameters"`
		AllRequiredParamsPresent bool        `json:"allRequiredParamsPresent"`
		OutputContexts           []struct {
			Name          string      `json:"name"`
			LifespanCount float64     `json:"lifespanCount"`
			Parameters    interface{} `json:"parameters"`
		} `json:"outputContexts"`
		Intent struct {
			Name           string `json:"name"`
			DisplayName    string `json:"displayName"`
			EndInteraction bool   `json:"endInteraction"`
		} `json:"intent"`
		IntentDetectionConfidence float64 `json:"intentDetectionConfidence"`
		LanguageCode              string  `json:"languageCode"`
	} `json:"queryResult"`
	OriginalDetectIntentRequest struct {
		Source  string `json:"source"`
		Payload struct {
			Source string `json:"source"`
			Data   struct {
				UpdateID float64 `json:"update_id"`
				Message  struct {
					Text string  `json:"text"`
					Date float64 `json:"date"`
					From struct {
						FirstName    string  `json:"first_name"`
						IsBot        bool    `json:"is_bot"`
						ID           float64 `json:"id"`
						Username     string  `json:"username"`
						LanguageCode string  `json:"language_code"`
						LastName     string  `json:"last_name"`
					} `json:"from"`
					Chat struct {
						Type      string  `json:"type"`
						LastName  string  `json:"last_name"`
						ID        float64 `json:"id"`
						FirstName string  `json:"first_name"`
						Username  string  `json:"username"`
					} `json:"chat"`
					MessageID float64 `json:"message_id"`
				} `json:"message"`
			} `json:"data"`
		} `json:"payload"`
	} `json:"originalDetectIntentRequest"`
	Session string `json:"session"`
}

type Payload struct {
	Username   string
	Id         int
	Parameters interface{}
}
