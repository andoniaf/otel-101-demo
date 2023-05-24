SHELL := /bin/bash
COMPOSE_FILES=-f docker-compose.yml
OTEL_DEMO_COMPOSE_FILES=-f opentelemetry-demo/docker-compose.yml

.PHONY: build
## Build images
build:
	docker-compose ${COMPOSE_FILES} build

.PHONY: up
## Run demo containers
up:
	docker-compose ${COMPOSE_FILES} up

.PHONY: up-otel
## Run compose for otel-demo
up-otel-demo:
	docker-compose ${OTEL_DEMO_COMPOSE_FILES} up --no-build


.PHONY: down-otel
## Stop otel-demo
down-otel:
	docker-compose ${OTEL_DEMO_COMPOSE_FILES} down

.PHONY: down
## Stop demo
down:
	docker-compose ${COMPOSE_FILES} down

.PHONY: build-script
## Build go binary script
build-script:
	@-mv demo-script/otel-101-demo demo-script/otel-101-demo.old
	cd demo-script && go build -o otel-101-demo .


###################################################################

# Plonk the following at the end of your Makefile
.DEFAULT_GOAL := show-help
# From https://gist.github.com/klmr/575726c7e05d8780505a
.PHONY: show-help
show-help:
	@echo "$$(tput bold)Available rules:$$(tput sgr0)"
	@echo
	@sed -n -e "/^## / { \
		h; \
		s/.*//; \
		:doc" \
		-e "H; \
		n; \
		s/^## //; \
		t doc" \
		-e "s/:.*//; \
		G; \
		s/\\n## /---/; \
		s/\\n/ /g; \
		p; \
	}" ${MAKEFILE_LIST} \
	| LC_ALL='C' sort --ignore-case \
	| awk -F '---' \
		-v ncol=$$(tput cols) \
		-v indent=19 \
		-v col_on="$$(tput setaf 6)" \
		-v col_off="$$(tput sgr0)" \
	'{ \
		printf "%s%*s%s ", col_on, -indent, $$1, col_off; \
		n = split($$2, words, " "); \
		line_length = ncol - indent; \
		for (i = 1; i <= n; i++) { \
			line_length -= length(words[i]) + 1; \
			if (line_length <= 0) { \
				line_length = ncol - indent - length(words[i]) - 1; \
				printf "\n%*s ", -indent, " "; \
			} \
			printf "%s ", words[i]; \
		} \
		printf "\n"; \
	}' \
	| more $(shell test $(shell uname) == Darwin && echo '--no-init --raw-control-chars')
