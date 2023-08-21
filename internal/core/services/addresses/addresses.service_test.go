package addresses_test

import (
	addressesservice "bia-challenge/internal/core/services/addresses"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddressFromMeterID_HappyPath(t *testing.T) {
	addressesService := addressesservice.NewAddressesService()
	address, err := addressesService.GetAddressFromMeterID(1)
	assert.Nil(t, err)
	assert.Equal(t, "some address mock for meter id 1", address)
}
