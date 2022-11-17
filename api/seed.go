package api

import (
	"github.com/google/uuid"
	"github.com/plummertr/udacity-go-crm/internal/data"
)

var InitData = make(map[string]data.Customer)

func (a *App) SeedData() *map[string]data.Customer {
	var initID1 = uuid.NewString()
	var initID2 = uuid.NewString()
	var initID3 = uuid.NewString()

	InitData[initID1] = data.Customer{
		ID:        initID1,
		Name:      "Tom Plummer",
		Role:      "The Creator",
		Email:     "fake@fake.com",
		Phone:     "818-555-1234",
		Contacted: true,
	}

	InitData[initID2] = data.Customer{
		ID:        initID2,
		Name:      "Plom Tummer",
		Role:      "Purchaser",
		Email:     "Plom@fake.com",
		Phone:     "323-555-5678",
		Contacted: false,
	}

	InitData[initID3] = data.Customer{
		ID:        initID3,
		Name:      "John Smith",
		Role:      "Purchaser",
		Email:     "John.Smith.32343232@fake.com",
		Phone:     "213-555-0987",
		Contacted: true,
	}
	return &InitData
}
