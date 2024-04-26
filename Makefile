START_LOG = @echo "======================================================= START OF LOG ========================================================="
END_LOG = @echo "======================================================== END OF LOG =========================================================="

.PHONY: env
env: ./config/.env.develop.tmpl
	cp ./config/.env.develop.tmpl ./config/.env.develop

.PHONY: infra
infra:
	$(START_LOG)
	@docker compose -f ./deployments/compose.yaml up --build -d
	$(END_LOG)

.PHONY: run
run:
	$(START_LOG)
	@docker compose \
		-f ./build/compose.yaml \
		--env-file ./config/.env.develop \
		up simulation streaming --build -d
	$(END_LOG)