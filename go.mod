module github.com/reapchain/ethermint

go 1.17

require (
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/go-bip39 v1.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/ethereum/go-ethereum v1.10.16
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/holiman/uint256 v1.2.0
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/miguelmota/go-ethereum-hdwallet v0.1.1
	github.com/onsi/ginkgo v1.16.5
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/reapchain/cosmos-sdk v0.45.1-reapsdkv2
	github.com/reapchain/ibc-go v0.3.0-reapibcv2
	github.com/reapchain/reapchain-core v0.1.5
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rs/cors v1.8.2
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.10.1
	github.com/status-im/keycard-go v0.0.0-20200402102358-957c09536969
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/tm-db v0.6.7
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f // indirect
	google.golang.org/genproto v0.0.0-20220504150022-98cd25cafc72
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
