/*
 * Twilio - Video
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
	"time"

	"github.com/patnunes/twilio-go/client"
)

func (c *ApiService) FetchRoomParticipant(RoomSid string, Sid string) (*VideoV1RoomParticipant, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoomParticipant'
type ListRoomParticipantParams struct {
	// Read only the participants with this status. Can be: `connected` or `disconnected`. For `in-progress` Rooms the default Status is `connected`, for `completed` Rooms only `disconnected` Participants are returned.
	Status *string `json:"Status,omitempty"`
	// Read only the Participants with this [User](https://www.twilio.com/docs/chat/rest/user-resource) `identity` value.
	Identity *string `json:"Identity,omitempty"`
	// Read only Participants that started after this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
	DateCreatedAfter *time.Time `json:"DateCreatedAfter,omitempty"`
	// Read only Participants that started before this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
	DateCreatedBefore *time.Time `json:"DateCreatedBefore,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoomParticipantParams) SetStatus(Status string) *ListRoomParticipantParams {
	params.Status = &Status
	return params
}
func (params *ListRoomParticipantParams) SetIdentity(Identity string) *ListRoomParticipantParams {
	params.Identity = &Identity
	return params
}
func (params *ListRoomParticipantParams) SetDateCreatedAfter(DateCreatedAfter time.Time) *ListRoomParticipantParams {
	params.DateCreatedAfter = &DateCreatedAfter
	return params
}
func (params *ListRoomParticipantParams) SetDateCreatedBefore(DateCreatedBefore time.Time) *ListRoomParticipantParams {
	params.DateCreatedBefore = &DateCreatedBefore
	return params
}
func (params *ListRoomParticipantParams) SetPageSize(PageSize int) *ListRoomParticipantParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoomParticipantParams) SetLimit(Limit int) *ListRoomParticipantParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RoomParticipant records from the API. Request is executed immediately.
func (c *ApiService) PageRoomParticipant(RoomSid string, params *ListRoomParticipantParams, pageToken, pageNumber string) (*ListRoomParticipantResponse, error) {
	path := "/v1/Rooms/{RoomSid}/Participants"

	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.DateCreatedAfter != nil {
		data.Set("DateCreatedAfter", fmt.Sprint((*params.DateCreatedAfter).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedBefore != nil {
		data.Set("DateCreatedBefore", fmt.Sprint((*params.DateCreatedBefore).Format(time.RFC3339)))
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

	ps := &ListRoomParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RoomParticipant records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoomParticipant(RoomSid string, params *ListRoomParticipantParams) ([]VideoV1RoomParticipant, error) {
	if params == nil {
		params = &ListRoomParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoomParticipant(RoomSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VideoV1RoomParticipant

	for response != nil {
		records = append(records, response.Participants...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRoomParticipantResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRoomParticipantResponse)
	}

	return records, err
}

// Streams RoomParticipant records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoomParticipant(RoomSid string, params *ListRoomParticipantParams) (chan VideoV1RoomParticipant, error) {
	if params == nil {
		params = &ListRoomParticipantParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoomParticipant(RoomSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VideoV1RoomParticipant, 1)

	go func() {
		for response != nil {
			for item := range response.Participants {
				channel <- response.Participants[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListRoomParticipantResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRoomParticipantResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRoomParticipantResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomParticipantResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateRoomParticipant'
type UpdateRoomParticipantParams struct {
	// The new status of the resource. Can be: `connected` or `disconnected`. For `in-progress` Rooms the default Status is `connected`, for `completed` Rooms only `disconnected` Participants are returned.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateRoomParticipantParams) SetStatus(Status string) *UpdateRoomParticipantParams {
	params.Status = &Status
	return params
}

func (c *ApiService) UpdateRoomParticipant(RoomSid string, Sid string, params *UpdateRoomParticipantParams) (*VideoV1RoomParticipant, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomParticipant{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
