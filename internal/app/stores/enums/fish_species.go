package enums

import "errors"

type FishSpeciesId uint8

const (
	FishSpeciesIdUnusedLower FishSpeciesId = iota
	FishSpeciesIdWalleye
	FishSpeciesIdCrappie
	FishSpeciesIdLargemouthBass
	FishSpeciesIdWhiteBass
	FishSpeciesIdWiper
	FishSpeciesIdStripedBass
	FishSpeciesIdRainbowTrout
	FishSpeciesIdUnusedUpper
)

type FishSpeciesName string

const (
	FishSpeciesNameWalleye        FishSpeciesName = "Walleye"
	FishSpeciesNameCrappie        FishSpeciesName = "Crappie"
	FishSpeciesNameLargemouthBass FishSpeciesName = "Largemouth Bass"
	FishSpeciesNameWhiteBass      FishSpeciesName = "White Bass"
	FishSpeciesNameWiper          FishSpeciesName = "Wiper"
	FishSpeciesNameStripedBass    FishSpeciesName = "Striped Bass"
	FishSpeciesNameRainbowTrout   FishSpeciesName = "Rainbow Trout"
)

func (fsId FishSpeciesId) FishSpeciesName() string {

	if fsId <= FishSpeciesIdUnusedLower || fsId >= FishSpeciesIdUnusedUpper {
		return ""
	}

	return string([]FishSpeciesName{
		FishSpeciesNameWalleye,
		FishSpeciesNameCrappie,
		FishSpeciesNameLargemouthBass,
		FishSpeciesNameWhiteBass,
		FishSpeciesNameWiper,
		FishSpeciesNameStripedBass,
		FishSpeciesNameRainbowTrout,
	}[fsId-FishSpeciesIdUnusedLower-1])
}

func (fsId FishSpeciesId) String() string {
	return fsId.FishSpeciesName()
}

func (fsId FishSpeciesId) IsValid() bool {
	if fsId <= FishSpeciesIdUnusedLower || fsId >= FishSpeciesIdUnusedUpper {
		return false
	}
	return true
}

func (fsId FishSpeciesId) MarshalText() ([]byte, error) {
	return []byte(fsId.String()), nil
}

var FishSpeciesIdMap map[string]FishSpeciesId

func init() {
	const size = FishSpeciesIdUnusedUpper - FishSpeciesIdUnusedLower - 1

	FishSpeciesIdMap = make(map[string]FishSpeciesId, size)

	for i := FishSpeciesIdUnusedLower + 1; i < FishSpeciesIdUnusedUpper; i++ {
		FishSpeciesIdMap[i.String()] = i
	}

}

var ErrInvalidFishSpeciesString = errors.New("fish_species_id: invalid fish species name string value")

func (fsId *FishSpeciesId) UnmarshalText(text []byte) error {

	v, ok := FishSpeciesIdMap[string(text)]
	if !ok {
		return ErrInvalidFishSpeciesString
	}

	*fsId = v

	return nil
}

func (fsName *FishSpeciesName) FishSpeciesId() (FishSpeciesId, error) {

	v, ok := FishSpeciesIdMap[string(*fsName)]
	if !ok {
		return v, ErrInvalidFishSpeciesString
	}

	return v, nil
}
