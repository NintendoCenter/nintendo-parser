package service

import (
	"NintendoCenter/nintendo-parser/internal/client"
	"NintendoCenter/nintendo-parser/internal/protos"
)

type Parser struct {
	client client.Client
}

func NewParser(client client.Client) *Parser {
	return &Parser{client: client}
}

func (p *Parser) ParseGames(offset int, limit int) []protos.Game {
	list := p.client.GetGameList(offset, limit)
	result := make([]protos.Game, 0, len(list))
	if len(list) == 0 {
		return result
	}

	for _, dto := range list {
		game := p.mapDtoToMode(dto)
		result = append(result, game)
	}

	return result
}

func (p *Parser) mapDtoToMode(dto client.NintendoGameDto) protos.Game {
	offer := protos.Offer{
		Shop:      protos.Shop_NINTENDO,
		IsDigital: true,
		IsUsed:    false,
		Link:      "https://nintendo.ru" + dto.Url,
		Price:     p.getPriceOutOfDto(dto),
	}

	game := protos.Game{
		Id:          dto.ID,
		Title:       dto.Title,
		Description: dto.Description,
		ImageUrl:    dto.ImageUrl,
	}

	game.Offers = append(game.Offers, &offer)

	return game
}

func (p *Parser) getPriceOutOfDto(dto client.NintendoGameDto) *protos.Price {
	var price *protos.Price = nil
	if dto.OriginalPrice > 0 {
		realPrice := dto.OriginalPrice
		if dto.DiscountedPrice > 0 {
			realPrice = dto.DiscountedPrice
		}
		if dto.OriginalPrice == dto.DiscountedPrice {
			dto.DiscountedPrice = 0
		}
		price = &protos.Price{
			Original:   float32(dto.OriginalPrice),
			Discounted: float32(dto.DiscountedPrice),
			Real:       float32(realPrice),
		}
	}

	return price
}