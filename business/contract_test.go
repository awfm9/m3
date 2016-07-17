package business

import "github.com/ethereum/go-ethereum/common"

type fakeContract struct{}

func (fc *fakeContract) Address() common.Address {
	return common.Address{}
}
