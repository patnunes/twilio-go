/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/patnunes/twilio-go/client"
)

// Optional parameters for the method 'CreateMessage'
type CreateMessageParams struct {
	//
	Attributes *string `json:"Attributes,omitempty"`
	//
	Body *string `json:"Body,omitempty"`
	//
	From *string `json:"From,omitempty"`
}

func (params *CreateMessageParams) SetAttributes(Attributes string) *CreateMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateMessageParams) SetBody(Body string) *CreateMessageParams {
	params.Body = &Body
	return params
}
func (params *CreateMessageParams) SetFrom(From string) *CreateMessageParams {
	params.From = &From
	return params
}

func (c *ApiService) CreateMessage(ServiceSid string, ChannelSid string, params *CreateMessageParams) (*IpMessagingV1Message, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteMessage(ServiceSid string, ChannelSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchMessage(ServiceSid string, ChannelSid string, Sid string) (*IpMessagingV1Message, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListMessage'
type ListMessageParams struct {
	//
	Order *string `json:"Order,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMessageParams) SetOrder(Order string) *ListMessageParams {
	params.Order = &Order
	return params
}
func (params *ListMessageParams) SetPageSize(PageSize int) *ListMessageParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMessageParams) SetLimit(Limit int) *ListMessageParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Message records from the API. Request is executed immediately.
func (c *ApiService) PageMessage(ServiceSid string, ChannelSid string, params *ListMessageParams, pageToken, pageNumber string) (*ListMessageResponse, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Message records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMessage(ServiceSid string, ChannelSid string, params *ListMessageParams) ([]IpMessagingV1Message, error) {
	if params == nil {
		params = &ListMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMessage(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []IpMessagingV1Message

	for response != nil {
		records = append(records, response.Messages...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListMessageResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListMessageResponse)
	}

	return records, err
}

// Streams Message records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMessage(ServiceSid string, ChannelSid string, params *ListMessageParams) (chan IpMessagingV1Message, error) {
	if params == nil {
		params = &ListMessageParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageMessage(ServiceSid, ChannelSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan IpMessagingV1Message, 1)

	go func() {
		for response != nil {
			for item := range response.Messages {
				channel <- response.Messages[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListMessageResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListMessageResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListMessageResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMessageResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateMessage'
type UpdateMessageParams struct {
	//
	Attributes *string `json:"Attributes,omitempty"`
	//
	Body *string `json:"Body,omitempty"`
}

func (params *UpdateMessageParams) SetAttributes(Attributes string) *UpdateMessageParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateMessageParams) SetBody(Body string) *UpdateMessageParams {
	params.Body = &Body
	return params
}

func (c *ApiService) UpdateMessage(ServiceSid string, ChannelSid string, Sid string, params *UpdateMessageParams) (*IpMessagingV1Message, error) {
	path := "/v1/Services/{ServiceSid}/Channels/{ChannelSid}/Messages/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ChannelSid"+"}", ChannelSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV1Message{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
