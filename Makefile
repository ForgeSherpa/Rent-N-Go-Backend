.DEFAULT_GOAL := install


define basic_install
	echo "Installing Go Dependecies"
	go mod download
endef

define unix_install
	echo "Generating queries"
	go run ./generator/*.go
	echo "Installation finish."
endef

define posix_install
	echo "Generating queries"
	go run "./generator/*.go"
	echo "Installation finish."
endef

generate:
	go run ./generator/*.go

migrate:
	go run *.go migrate

seed:
	go run *.go seed

windows:
	echo "Installing Project..."
	echo "Running pnpm install"
	pnpm install
	$(call basic_install)
	$(call posix_install)

windows_install_npm:
	echo "Installing Project"
	echo "Running 'npm install'"
	npm install
	$(call basic_install)
	$(call posix_install)

install:
	echo "Installing Project..."
	echo "Running pnpm install"
	pnpm install
	$(call basic_install)
	$(call unix_install)


install_npm:
	echo "Installing Project"
	echo "Running 'npm install'"
	npm install
	$(call basic_install)
	$(call unix_install)

run:
	go run *.go

build:
	go build .
