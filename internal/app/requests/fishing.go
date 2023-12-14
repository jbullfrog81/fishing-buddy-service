package requests

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jbullfrog81/fishing-buddy-service/internal/app/stores"
	"github.com/jbullfrog81/fishing-buddy-service/internal/app/stores/enums"
)

type HeadersCommon struct {
	// in:header
	HeaderAuthorization string `json:"Authorization"`
}

type PostCatchRequestBody struct {
	FishSpeciesId enums.FishSpeciesId `json:"fish_species_id"`
	FishermanId   int                 `json:"fisherman_id"`
	Coordinates   stores.Coordinates  `json:"coordinates"`
}

type PostCatchRequest struct {
	HeadersCommon
	Body PostCatchRequestBody
}

func ValidateRequestHeader(headerValues string, supportedValue string) error {

	firstHeaderValue := strings.ToLower(strings.TrimSpace(strings.Split(headerValues, ";")[0]))
	if firstHeaderValue != supportedValue {
		return errors.New(fmt.Sprintf("header value %s does not match supported value %s", headerValues, supportedValue))
	}

	return nil
}
