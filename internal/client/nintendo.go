package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"go.uber.org/zap"
)

type Client interface {
	GetGameList(offset int, limit int) ([]NintendoGameDto, error)
	FindByName(name string) (*NintendoGameDto, error)
}

type NintendoResponse struct {
	Response struct {
		Docs []NintendoGameDto `json:"docs"`
	} `json:"response"`
}

type NintendoGameDto struct {
	ID              string  `json:"fs_id"`
	Title           string  `json:"title"`
	Description     string  `json:"excerpt"`
	ImageUrl        string  `json:"image_url"`
	Url             string  `json:"url"`
	Type            string  `json:"type"`
	OriginalPrice   float64 `json:"price_regular_f"`
	DiscountedPrice float64 `json:"price_sorting_f"`
	DiscountSize    float64 `json:"price_discount_percentage_f,omitempty"`
}

type Nintendo struct {
	baseUrl string
	filter  string
	logger  *zap.Logger
}

func NewNintendoClient(l *zap.Logger) *Nintendo {
	return &Nintendo{
		baseUrl: "https://searching.nintendo-europe.com/ru/select",
		filter:  `type:GAME AND ((playable_on_txt:"HAC")) AND sorting_title:* AND *:*`,
		logger:  l,
	}
}

func (c *Nintendo) GetGameList(offset int, limit int) ([]NintendoGameDto, error) {
	query := c.prepareUrl(offset, limit, nil)
	r, err := c.fetch(query)
	if err != nil {
		return []NintendoGameDto{}, err
	}
	return r.Response.Docs, nil
}

func (c *Nintendo) FindByName(name string) (*NintendoGameDto, error) {
	query := c.prepareUrl(0, 1, &name)
	r, err := c.fetch(query)
	if err != nil {
		return nil, err
	}
	game := r.Response.Docs[0]
	return &game, nil
}

func (c *Nintendo) fetch(requestUrl string) (*NintendoResponse, error) {
	c.logger.Debug("fetching url: " + requestUrl)
	resp, err := http.Get(requestUrl)
	if err != nil {
		c.logger.Error("error while making request", zap.Error(err))
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.logger.Error("error on closing body stream", zap.Error(err))
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("cannot read response body", zap.Error(err))
		return nil, err
	}

	var response NintendoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.logger.Error("cannot unmarshal response to json", zap.Error(err))
		return nil, err
	}

	return &response, nil
}

func (c *Nintendo) prepareUrl(offset int, limit int, name *string) string {
	searchingName := "*"
	if name != nil {
		searchingName = *name
	}

	v := url.Values{}
	v.Set("q", searchingName)
	v.Set("start", strconv.Itoa(offset))
	v.Set("rows", strconv.Itoa(limit))
	v.Set("wt", "json")
	v.Set("fq", c.filter)

	return c.baseUrl + "?" + v.Encode()
}
