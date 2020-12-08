package nitwit

// LookupTweetsByIDs returns a variety of information about the Tweets specified by the requested ID or list of IDs.
func (c *Client) LookupTweetsByIDs(id string, opts ...RequestOption) (*Messages, error) {
	endpointUrl := "https://api.stocktwits.com/api/2/streams/symbol/" + id + ".json"
	req, err := NewRequest(endpointUrl, opts...)
	if err != nil {
		return nil, err
	}
	messages := new(Messages)
	err = c.MakeRequest(req, messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
