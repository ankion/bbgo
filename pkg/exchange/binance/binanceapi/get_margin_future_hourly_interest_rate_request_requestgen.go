// Code generated by "requestgen -method GET -url /sapi/v1/margin/next-hourly-interest-rate -type GetMarginFutureHourlyInterestRateRequest -responseType []HourlyInterestRate"; DO NOT EDIT.

package binanceapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func (g *GetMarginFutureHourlyInterestRateRequest) Assets(assets string) *GetMarginFutureHourlyInterestRateRequest {
	g.assets = assets
	return g
}

func (g *GetMarginFutureHourlyInterestRateRequest) Isolated(isolated string) *GetMarginFutureHourlyInterestRateRequest {
	g.isolated = isolated
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetMarginFutureHourlyInterestRateRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetMarginFutureHourlyInterestRateRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check assets field -> json key assets
	assets := g.assets

	// assign parameter of assets
	params["assets"] = assets
	// check isolated field -> json key isolated
	isolated := g.isolated

	// assign parameter of isolated
	params["isolated"] = isolated

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetMarginFutureHourlyInterestRateRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetMarginFutureHourlyInterestRateRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetMarginFutureHourlyInterestRateRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetMarginFutureHourlyInterestRateRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetMarginFutureHourlyInterestRateRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetMarginFutureHourlyInterestRateRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetMarginFutureHourlyInterestRateRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

// GetPath returns the request path of the API
func (g *GetMarginFutureHourlyInterestRateRequest) GetPath() string {
	return "/sapi/v1/margin/next-hourly-interest-rate"
}

// Do generates the request object and send the request object to the API endpoint
func (g *GetMarginFutureHourlyInterestRateRequest) Do(ctx context.Context) ([]HourlyInterestRate, error) {

	// empty params for GET operation
	var params interface{}
	query, err := g.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	var apiURL string

	apiURL = g.GetPath()

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []HourlyInterestRate

	type responseUnmarshaler interface {
		Unmarshal(data []byte) error
	}

	if unmarshaler, ok := interface{}(&apiResponse).(responseUnmarshaler); ok {
		if err := unmarshaler.Unmarshal(response.Body); err != nil {
			return nil, err
		}
	} else {
		// The line below checks the content type, however, some API server might not send the correct content type header,
		// Hence, this is commented for backward compatibility
		// response.IsJSON()
		if err := response.DecodeJSON(&apiResponse); err != nil {
			return nil, err
		}
	}

	type responseValidator interface {
		Validate() error
	}

	if validator, ok := interface{}(&apiResponse).(responseValidator); ok {
		if err := validator.Validate(); err != nil {
			return nil, err
		}
	}
	return apiResponse, nil
}
