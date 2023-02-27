package seed

import (
	"context"
	"time"

	"github.com/maddiesch/collector/internal/core/domain"
	"github.com/maddiesch/collector/internal/repositories/database"
	"github.com/stretchr/testify/require"
)

func Collection(t require.TestingT, db *database.Database) {
	cards := []domain.CollectedCard{
		{
			ID:              "01GT8H40W06W2Y7WGEWNFVAKGZ",
			GroupName:       "01GT8FQ4H1FXHC026PPRBDWMFDt",
			SetName:         "Phyrexia: All Will Be One",
			Name:            "Plains",
			CollectorNumber: "272",
			IsFoil:          false,
			Language:        "English",
			Condition:       domain.CardConditionMint,
			CreatedAt:       time.Unix(1677469661, 0),
			UpdatedAt:       time.Unix(1677469661, 0),
		},
		{
			ID:              "01GT8H488T11CE897YWYPJ18RP",
			GroupName:       "01GT8FQ4H1FXHC026PPRBDWMFDt",
			SetName:         "Phyrexia: All Will Be One",
			Name:            "Cutthroat Centurion",
			CollectorNumber: "89",
			IsFoil:          false,
			Language:        "English",
			Condition:       domain.CardConditionMint,
			CreatedAt:       time.Unix(1677469688, 0),
			UpdatedAt:       time.Unix(1677469688, 0),
		},
		{
			ID:              "01GT8H4FAEYSKS33PNQVDHV64P",
			GroupName:       "01GT8FQ4H1FXHC026PPRBDWMFDt",
			SetName:         "Phyrexia: All Will Be One",
			Name:            "Forest",
			CollectorNumber: "369",
			IsFoil:          true,
			Language:        "Phyrexian",
			Condition:       domain.CardConditionMint,
			CreatedAt:       time.Unix(1677469711, 0),
			UpdatedAt:       time.Unix(1677469711, 0),
		},
		{
			ID:              "01GT8MKN90SKSS33YARXH4GS74",
			GroupName:       "named-import",
			SetName:         "Phyrexia: All Will Be One",
			Name:            "Glissa Sunslayer",
			CollectorNumber: "318",
			IsFoil:          false,
			Language:        "English",
			Condition:       domain.CardConditionNearMint,
			CreatedAt:       time.Unix(1677474893, 0),
			UpdatedAt:       time.Unix(1677474893, 0),
		},
		{
			ID:              "01GT8MM4RXVBGREGY2X77Z7NND",
			GroupName:       "last-import",
			SetName:         "Phyrexia: All Will Be One",
			Name:            "Mountain",
			CollectorNumber: "275",
			IsFoil:          false,
			Language:        "English",
			Condition:       domain.CardConditionMint,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}
	ctx := context.Background()
	for _, card := range cards {
		err := db.InsertCollectedCard(ctx, card)

		require.NoError(t, err)
	}
}
