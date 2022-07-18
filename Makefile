TEST_PATH:=./test/...
REPORT_PATH:=./test/allure-results

.PHONY: .install_allure
.install_allure:
ifeq ($(OS),Windows_NT)
	$(info Run Windows run pattern...)
	$(info Make sure scoop installed at your system. Check for more information: https://github.com/ScoopInstaller/Scoop#installation)
	scoop install allure
endif
ifeq ($(OS),Linux)
	$(info Run Linux run pattern...)
	$(info Make sure you have sudo rights for the system.)
	sudo apt-add-repository ppa:qameta/allure
	sudo apt-get update
	sudo apt-get install allure
endif
ifeq ($(OS),Darwin)
	$(info Run installation for Darwin OS)
	$(info Make sure brew installed at your system. Check for more information: https://docs.brew.sh/Installation)
	brew install allure
endif

.PHONY: install
install:
	go mod tidy

.PHONY: install-full
install-full: .install_allure install

.PHONY: demo
demo: test allure-serve

.PHONY: test
test:
	go test -v ${TEST_PATH}

.PHONY: allure-serve
allure-serve:
	allure serve ${REPORT_PATH}
