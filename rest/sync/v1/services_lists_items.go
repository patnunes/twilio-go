/*
 * Twilio - Sync
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

// Optional parameters for the method 'CreateSyncListItem'
type CreateSyncListItemParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item's parent Sync List expires (time-to-live) and is deleted.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
	// A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
	Data *map[string]interface{} `json:"Data,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
	ItemTtl *int `json:"ItemTtl,omitempty"`
	// An alias for `item_ttl`. If both parameters are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *CreateSyncListItemParams) SetCollectionTtl(CollectionTtl int) *CreateSyncListItemParams {
	params.CollectionTtl = &CollectionTtl
	return params
}
func (params *CreateSyncListItemParams) SetData(Data map[string]interface{}) *CreateSyncListItemParams {
	params.Data = &Data
	return params
}
func (params *CreateSyncListItemParams) SetItemTtl(ItemTtl int) *CreateSyncListItemParams {
	params.ItemTtl = &ItemTtl
	return params
}
func (params *CreateSyncListItemParams) SetTtl(Ttl int) *CreateSyncListItemParams {
	params.Ttl = &Ttl
	return params
}

func (c *ApiService) CreateSyncListItem(ServiceSid string, ListSid string, params *CreateSyncListItemParams) (*SyncV1SyncListItem, error) {
	path := "/v1/Services/{ServiceSid}/Lists/{ListSid}/Items"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ListSid"+"}", ListSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}
	if params != nil && params.Data != nil {
		v, err := json.Marshal(params.Data)

		if err != nil {
			return nil, err
		}

		data.Set("Data", string(v))
	}
	if params != nil && params.ItemTtl != nil {
		data.Set("ItemTtl", fmt.Sprint(*params.ItemTtl))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncListItem{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSyncListItem'
type DeleteSyncListItemParams struct {
	// If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
	IfMatch *string `json:"If-Match,omitempty"`
}

func (params *DeleteSyncListItemParams) SetIfMatch(IfMatch string) *DeleteSyncListItemParams {
	params.IfMatch = &IfMatch
	return params
}

func (c *ApiService) DeleteSyncListItem(ServiceSid string, ListSid string, Index int, params *DeleteSyncListItemParams) error {
	path := "/v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ListSid"+"}", ListSid, -1)
	path = strings.Replace(path, "{"+"Index"+"}", fmt.Sprint(Index), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchSyncListItem(ServiceSid string, ListSid string, Index int) (*SyncV1SyncListItem, error) {
	path := "/v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ListSid"+"}", ListSid, -1)
	path = strings.Replace(path, "{"+"Index"+"}", fmt.Sprint(Index), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncListItem{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncListItem'
type ListSyncListItemParams struct {
	// How to order the List Items returned by their `index` value. Can be: `asc` (ascending) or `desc` (descending) and the default is ascending.
	Order *string `json:"Order,omitempty"`
	// The `index` of the first Sync List Item resource to read. See also `bounds`.
	From *string `json:"From,omitempty"`
	// Whether to include the List Item referenced by the `from` parameter. Can be: `inclusive` to include the List Item referenced by the `from` parameter or `exclusive` to start with the next List Item. The default value is `inclusive`.
	Bounds *string `json:"Bounds,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSyncListItemParams) SetOrder(Order string) *ListSyncListItemParams {
	params.Order = &Order
	return params
}
func (params *ListSyncListItemParams) SetFrom(From string) *ListSyncListItemParams {
	params.From = &From
	return params
}
func (params *ListSyncListItemParams) SetBounds(Bounds string) *ListSyncListItemParams {
	params.Bounds = &Bounds
	return params
}
func (params *ListSyncListItemParams) SetPageSize(PageSize int) *ListSyncListItemParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSyncListItemParams) SetLimit(Limit int) *ListSyncListItemParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SyncListItem records from the API. Request is executed immediately.
func (c *ApiService) PageSyncListItem(ServiceSid string, ListSid string, params *ListSyncListItemParams, pageToken, pageNumber string) (*ListSyncListItemResponse, error) {
	path := "/v1/Services/{ServiceSid}/Lists/{ListSid}/Items"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ListSid"+"}", ListSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Bounds != nil {
		data.Set("Bounds", *params.Bounds)
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

	ps := &ListSyncListItemResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SyncListItem records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncListItem(ServiceSid string, ListSid string, params *ListSyncListItemParams) ([]SyncV1SyncListItem, error) {
	if params == nil {
		params = &ListSyncListItemParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncListItem(ServiceSid, ListSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SyncV1SyncListItem

	for response != nil {
		records = append(records, response.Items...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncListItemResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSyncListItemResponse)
	}

	return records, err
}

// Streams SyncListItem records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncListItem(ServiceSid string, ListSid string, params *ListSyncListItemParams) (chan SyncV1SyncListItem, error) {
	if params == nil {
		params = &ListSyncListItemParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncListItem(ServiceSid, ListSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SyncV1SyncListItem, 1)

	go func() {
		for response != nil {
			for item := range response.Items {
				channel <- response.Items[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncListItemResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSyncListItemResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSyncListItemResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncListItemResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSyncListItem'
type UpdateSyncListItemParams struct {
	// If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
	IfMatch *string `json:"If-Match,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item's parent Sync List expires (time-to-live) and is deleted. This parameter can only be used when the List Item's `data` or `ttl` is updated in the same request.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
	// A JSON string that represents an arbitrary, schema-less object that the List Item stores. Can be up to 16 KiB in length.
	Data *map[string]interface{} `json:"Data,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the List Item expires (time-to-live) and is deleted.
	ItemTtl *int `json:"ItemTtl,omitempty"`
	// An alias for `item_ttl`. If both parameters are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *UpdateSyncListItemParams) SetIfMatch(IfMatch string) *UpdateSyncListItemParams {
	params.IfMatch = &IfMatch
	return params
}
func (params *UpdateSyncListItemParams) SetCollectionTtl(CollectionTtl int) *UpdateSyncListItemParams {
	params.CollectionTtl = &CollectionTtl
	return params
}
func (params *UpdateSyncListItemParams) SetData(Data map[string]interface{}) *UpdateSyncListItemParams {
	params.Data = &Data
	return params
}
func (params *UpdateSyncListItemParams) SetItemTtl(ItemTtl int) *UpdateSyncListItemParams {
	params.ItemTtl = &ItemTtl
	return params
}
func (params *UpdateSyncListItemParams) SetTtl(Ttl int) *UpdateSyncListItemParams {
	params.Ttl = &Ttl
	return params
}

func (c *ApiService) UpdateSyncListItem(ServiceSid string, ListSid string, Index int, params *UpdateSyncListItemParams) (*SyncV1SyncListItem, error) {
	path := "/v1/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"ListSid"+"}", ListSid, -1)
	path = strings.Replace(path, "{"+"Index"+"}", fmt.Sprint(Index), -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}
	if params != nil && params.Data != nil {
		v, err := json.Marshal(params.Data)

		if err != nil {
			return nil, err
		}

		data.Set("Data", string(v))
	}
	if params != nil && params.ItemTtl != nil {
		data.Set("ItemTtl", fmt.Sprint(*params.ItemTtl))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncListItem{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
