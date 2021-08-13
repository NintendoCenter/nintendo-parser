package service

import (
	"fmt"

	"NintendoCenter/nintendo-parser/internal/client"
	"NintendoCenter/nintendo-parser/internal/protos"
	"go.uber.org/zap"
)

type Parser struct {
	logger *zap.Logger
	client client.Client
}

func NewParser(client client.Client, l *zap.Logger) *Parser {
	return &Parser{client: client, logger: l}
}

func (p *Parser) ParseGames(offset int, limit int) ([]*protos.Game, error) {
	list, err := p.client.GetGameList(offset, limit)
	if err != nil {
		return []*protos.Game{}, err
	}
	result := make([]*protos.Game, 0, len(list))
	if len(list) == 0 {
		p.logger.Info("got an empty list from the client")
		return result, nil
	}
	p.logger.Info(fmt.Sprintf("fetched %d items with offset %d and limit %d", len(list), offset, limit))
	for _, dto := range list {
		game := p.mapDtoToMode(dto)
		result = append(result, &game)
	}

	return result, nil
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