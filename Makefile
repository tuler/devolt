START_LOG = @echo "======================================================= START OF LOG ========================================================="
END_LOG = @echo "======================================================== END OF LOG =========================================================="

.PHONY: env
env: ./config/.env.develop.tmpl
	cp ./config/.env.develop.tmpl ./config/.env.develop

.PHONY: infra
infra:
	$(START_LOG)
	@docker compose -f ./deployments/compose.infra.yaml up --build -d
	$(END_LOG)

.PHONY: run
run:
	$(START_LOG)
	@docker compose \
		-f ./deployments/compose.packages.yaml \
		--env-file ./config/.env.develop \
		up simulation streaming --build -d
	@sunodo run
	$(END_LOG)

.PHONY: build
build:
	$(START_LOG)
	@docker build \
		-t rollup \
		-f ./build/Dockerfile.rollup .
	@sunodo build --from-image rollup
	$(END_LOG)