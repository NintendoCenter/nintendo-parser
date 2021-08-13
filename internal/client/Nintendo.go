package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	GetGameList(offset int, limit int) []NintendoGameDto
	FindByName(name string) *NintendoGameDto
}

type NintendoResponse struct {
	Response struct{
		Docs []NintendoGameDto `json:"docs"`
	} `json:"response"`
}

type NintendoGameDto struct {
	ID string `json:"fs_id"`
	Title string `json:"title"`
	Description string `json:"excerpt"`
	ImageUrl string `json:"image_url"`
	Url string `json:"url"`
	Type string `json:"type"`
	OriginalPrice float64 `json:"price_regular_f"`
	DiscountedPrice float64 `json:"price_sorting_f"`
	DiscountSize float64 `json:"price_discount_percentage_f,omitempty"`
}

type Nintendo struct{
	baseUrl string
	filter string
}

func NewNintendoClient() *Nintendo {
	return &Nintendo{
		baseUrl: "https://searching.nintendo-europe.com/ru/select",
		filter: `type:GAME AND ((playable_on_txt:"HAC")) AND sorting_title:* AND *:*`,
	}
}

func (c *Nintendo) GetGameList(offset int, limit int) []NintendoGameDto {
	query := c.prepareUrl(offset, limit, nil)
	if r, err := c.fetch(query); err == nil {
		return r.Response.Docs
	}
	return []NintendoGameDto{}
}

func (c *Nintendo) FindByName(name string) *NintendoGameDto {
	query := c.prepareUrl(0, 1, &name)
	if r, err := c.fetch(query); err == nil {
		game := r.Response.Docs[0]
		return &game
	}
	return nil
}

func (c *Nintendo) fetch(requestUrl string) (*NintendoResponse, error) {
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response NintendoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
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
