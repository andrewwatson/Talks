// START ONE OMIT
func SendMessage(req *http.Request, to, from, message string) (err error) {
	c := appengine.NewContext(req)
	client := urlfetch.Client(c)
	t := twilioclient.NewTwilioClient(twilioSID, twilioSecret)

	err = t.SendMessage(*client, to, from, message)

	c.Debugf("SENDING %s %s %s", to, from, message)
	if err != nil {
		c.Debugf("Twilio Error: %s", err.Error())
	}

	return
}
// STOP ONE OMIT

// START TWO OMIT
func (t *TwilioClient) SendMessage(client http.Client, toNumber, fromNumber, message string) (err error) {

	...
	twilioRequest.SetBasicAuth(t.AccountSid, t.AuthToken)
	twilioRequest.Header.Add("Content-type", "application/x-www-form-urlencoded")
	resp, clientError := client.Do(twilioRequest)
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		err = errors.New(fmt.Sprintf("TWILIO ERROR: %d", resp.StatusCode))		
	}
	return clientError
}
// STOP TWO OMIT