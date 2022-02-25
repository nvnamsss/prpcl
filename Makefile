run:
	@cd src/cmd && go run ./*.go

build:
	@cd $(PWD)/src/cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prpcl .

migration:
	@cd $(PWD)/src/migrations && read -p "Enter migration name: " migration_name; \
	goose create $${migration_name} sql

migrate:
	@bash $(PWD)/scripts/migration.sh

ADAPTERS_DIR := $(patsubst %,%,$(notdir $(wildcard $(PWD)/src/adapters/*)))
gen_mock_adapters:
	@for dir in $(ADAPTERS_DIR); do \
		cd $(PWD)/src && mockery --case=underscore --dir=$(PWD)/src/adapters/$$dir --output $(PWD)/src/mocks/adapters/$$dir --all ; \
	done

gen_mock_repo:
	cd $(PWD)/src && mockery --case=underscore --dir=$(PWD)/src/repositories --output $(PWD)/src/mocks/repositories --all ; \

gen_docs:
	@cp src/cmd/main.go src/
	@cd src && swag init
	@rm -rf src/main.go
	@rm -rf src/cmd/docs
	@mv src/docs src/cmd
test: ## Run unit tests
	@./scripts/unittest.sh $(SUB_PROJECTS)
