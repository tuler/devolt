-include .env.develop

START_LOG = @echo "================================================= START OF LOG ==================================================="
END_LOG = @echo "================================================== END OF LOG ===================================================="

RPC_URL := http://localhost:8545
PRIVATE_KEY := 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

ifeq ($(NETWORK), localhost)
	BYTECODE_NETWORK_ARGS := script/Bytecode.s.sol --rpc-url $(RPC_URL) --private-key $(PRIVATE_KEY) --broadcast -v
	DEPLOY_NETWORK_ARGS := script/DeployProxy.s.sol --rpc-url $(RPC_URL) --private-key $(PRIVATE_KEY) --broadcast -v
else
	RPC_URL := $(TESTNET_RPC_URL)
	PRIVATE_KEY := $(TESTNET_PRIVATE_KEY)
	BYTECODE_NETWORK_ARGS := script/ExecutedVouchers.s.sol --rpc-url $(RPC_URL) --private-key $(PRIVATE_KEY) --broadcast --verify --etherscan-api-key $(TESTNET_BLOCKSCAN_API_KEY) -v
	DEPLOY_NETWORK_ARGS := script/DeployProxy.s.sol --rpc-url $(RPC_URL) --private-key $(PRIVATE_KEY) --broadcast --verify --etherscan-api-key $(TESTNET_BLOCKSCAN_API_KEY) -v
endif

.PHONY: env
env: ./.env.develop.tmpl
	cp ./.env.develop.tmpl ./.env.develop

.PHONY: infra
infra:
	$(START_LOG)
	@docker compose -f ./deployments/compose.infra.yaml up --build -d
	$(END_LOG)

.PHONY: dev
dev:
	$(START_LOG)
	@nonodo -- go run ./cmd/rollup/
	$(END_LOG)

.PHONY: build
build:
	$(START_LOG)
	@docker build \
		-t rollup \
		-f ./build/Dockerfile.rollup .
	@cartesi build --from-image rollup
	$(END_LOG)

.PHONY: iot
iot:
	$(START_LOG)
	@docker compose \
		-f ./deployments/compose.packages.yaml \
		--env-file ./.env.develop \
		up simulation streaming --build -d
	$(END_LOG)

.PHONY: prod
prod:
	$(START_LOG)
	@cartesi run --epoch-duration 60
	$(END_LOG)
	
.PHONY: generate
generate:
	$(START_LOG)
	@go run ./pkg/rollups-contracts/generate
	$(END_LOG)

.PHONY: test
test:
	@echo "TBD"

.PHONY: deploy
deploy:
	$(START_LOG)
	@cd contracts && forge script $(DEPLOY_NETWORK_ARGS)
	$(END_LOG)

.PHONY: bytecode
bytecode:
	$(START_LOG)
	@cd contracts && forge script $(BYTECODE_NETWORK_ARGS)
	$(END_LOG)