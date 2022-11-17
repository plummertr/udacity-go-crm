package api

import (
	"github.com/google/uuid"
	"github.com/plummertr/udacity-go-crm/internal/data"
)

var InitData = make(map[string]data.Customer)

func (a *App) SeedData() *map[string]data.Customer {
	var initID1 = uuid.NewString()
	var initID2 = uuid.NewString()

	InitData[initID1] = data.Customer{
		ID:   initID1,
		Name: "Tom Plummer",
	}

	InitData[initID2] = data.Customer{
		ID:   initID2,
		Name: "Plom Tummer",
	}

	return &InitData
}
