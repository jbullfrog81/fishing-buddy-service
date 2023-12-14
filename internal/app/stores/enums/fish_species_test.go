package enums_test

import (
	"fmt"
	"testing"

	"github.com/jbullfrog81/fishing-buddy-service/internal/app/stores/enums"
	"github.com/stretchr/testify/assert"
)

func TestFishSpeciesName(t *testing.T) {

	// Healthy Test Cases
	testCases := []struct {
		fishSpeciesId       enums.FishSpeciesId
		expFishSpeciesId    uint8
		fishSpeciesName     enums.FishSpeciesName
		expFishSpeciesName  string
		expFishSpeciesIdErr error
	}{
		{
			enums.FishSpeciesIdUnusedLower,
			0,
			"",
			"",
			enums.ErrInvalidFishSpeciesString,
		},
		{
			enums.FishSpeciesIdWalleye,
			1,
			enums.FishSpeciesNameWalleye,
			enums.FishSpeciesIdWalleye.String(),
			nil,
		},
		{
			enums.FishSpeciesIdCrappie,
			2,
			enums.FishSpeciesNameCrappie,
			enums.FishSpeciesIdCrappie.String(),
			nil,
		},
		{
			enums.FishSpeciesIdLargemouthBass,
			3,
			enums.FishSpeciesNameLargemouthBass,
			enums.FishSpeciesIdLargemouthBass.String(),
			nil,
		},
		{
			enums.FishSpeciesIdWhiteBass,
			4,
			enums.FishSpeciesNameWhiteBass,
			enums.FishSpeciesIdWhiteBass.String(),
			nil,
		},
		{
			enums.FishSpeciesIdWiper,
			5,
			enums.FishSpeciesNameWiper,
			enums.FishSpeciesIdWiper.String(),
			nil,
		},
		{
			enums.FishSpeciesIdStripedBass,
			6,
			enums.FishSpeciesNameStripedBass,
			enums.FishSpeciesIdStripedBass.String(),
			nil,
		},
		{
			enums.FishSpeciesIdRainbowTrout,
			7,
			enums.FishSpeciesNameRainbowTrout,
			enums.FishSpeciesIdRainbowTrout.String(),
			nil,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Happy path enums fish species Id %d and name %s", tc.fishSpeciesId, tc.expFishSpeciesName)
		t.Run(name, func(t *testing.T) {
			fishSpeciesName := tc.fishSpeciesId.FishSpeciesName()
			assert.Equal(t, tc.expFishSpeciesName, fishSpeciesName)
			fishSpeciesId, err := tc.fishSpeciesName.FishSpeciesId()
			assert.Equal(t, tc.fishSpeciesId, fishSpeciesId)
			assert.Equal(t, tc.expFishSpeciesIdErr, err)
		})
	}
}
