/*
 * Twilio - Verify
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

// Optional parameters for the method 'CreateEntity'
type CreateEntityParams struct {
	// The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user's UUID, GUID, or SID. It can only contain dash (-) separated alphanumeric characters.
	Identity *string `json:"Identity,omitempty"`
}

func (params *CreateEntityParams) SetIdentity(Identity string) *CreateEntityParams {
	params.Identity = &Identity
	return params
}

// Create a new Entity for the Service
func (c *ApiService) CreateEntity(ServiceSid string, params *CreateEntityParams) (*VerifyV2Entity, error) {
	path := "/v2/Services/{ServiceSid}/Entities"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Entity{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Entity.
func (c *ApiService) DeleteEntity(ServiceSid string, Identity string) error {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific Entity.
func (c *ApiService) FetchEntity(ServiceSid string, Identity string) (*VerifyV2Entity, error) {
	path := "/v2/Services/{ServiceSid}/Entities/{Identity}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Identity"+"}", Identity, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VerifyV2Entity{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEntity'
type ListEntityParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEntityParams) SetPageSize(PageSize int) *ListEntityParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEntityParams) SetLimit(Limit int) *ListEntityParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Entity records from the API. Request is executed immediately.
func (c *ApiService) PageEntity(ServiceSid string, params *ListEntityParams, pageToken, pageNumber string) (*ListEntityResponse, error) {
	path := "/v2/Services/{ServiceSid}/Entities"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListEntityResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Entity records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEntity(ServiceSid string, params *ListEntityParams) ([]VerifyV2Entity, error) {
	if params == nil {
		params = &ListEntityParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEntity(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VerifyV2Entity

	for response != nil {
		records = append(records, response.Entities...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListEntityResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListEntityResponse)
	}

	return records, err
}

// Streams Entity records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEntity(ServiceSid string, params *ListEntityParams) (chan VerifyV2Entity, error) {
	if params == nil {
		params = &ListEntityParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEntity(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VerifyV2Entity, 1)

	go func() {
		for response != nil {
			for item := range response.Entities {
				channel <- response.Entities[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListEntityResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListEntityResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListEntityResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEntityResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
