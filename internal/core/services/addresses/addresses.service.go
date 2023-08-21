package addresses

import "fmt"

type addressesService struct{}

func NewAddressesService() *addressesService {
	return &addressesService{}
}

func (srv *addressesService) GetAddressFromMeterID(meterID int) (string, error) {
	// TODO: Implement logic to get address
	address := fmt.Sprintf("some address mock for meter id %v", meterID)
	return address, nil
}
