package types_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	"github.com/reapchain/cosmos-sdk/codec"
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
	sdk "github.com/reapchain/cosmos-sdk/types"
	authtypes "github.com/reapchain/cosmos-sdk/x/auth/types"

	cryptocodec "github.com/tharsis/ethermint/crypto/codec"
	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
	ethermintcodec "github.com/tharsis/ethermint/encoding/codec"
	"github.com/tharsis/ethermint/types"
)

func init() {
	amino := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(amino)
}

type AccountTestSuite struct {
	suite.Suite

	account *types.EthAccount
	cdc     codec.JSONCodec
}

func (suite *AccountTestSuite) SetupTest() {
	privKey, err := ethsecp256k1.GenerateKey()
	suite.Require().NoError(err)
	pubKey := privKey.PubKey()
	addr := sdk.AccAddress(pubKey.Address())
	baseAcc := authtypes.NewBaseAccount(addr, pubKey, 10, 50)
	suite.account = &types.EthAccount{
		BaseAccount: baseAcc,
		CodeHash:    common.Hash{}.String(),
	}

	interfaceRegistry := codectypes.NewInterfaceRegistry()
	ethermintcodec.RegisterInterfaces(interfaceRegistry)
	suite.cdc = codec.NewProtoCodec(interfaceRegistry)
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
