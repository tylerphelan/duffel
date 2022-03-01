package duffel

import (
	"context"
)

type (
	AirlinesClient interface {
		ListAirlines(ctx context.Context) *Iter[Airline]
		GetAirline(ctx context.Context, id string) (*Airline, error)
	}
)

func (a *API) ListAirlines(ctx context.Context) *Iter[Airline] {
	return newRequestWithAPI[EmptyPayload, Airline](a).
		Get("/air/airlines").
		All(ctx)
}

func (a *API) GetAirline(ctx context.Context, id string) (*Airline, error) {
	return newRequestWithAPI[EmptyPayload, Airline](a).
		Getf("/air/airlines/%s", id).
		One(ctx)
}

var _ AirlinesClient = (*API)(nil)
