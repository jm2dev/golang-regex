##
# Regular expressions
#
# @file
# @version 0.1
.PHONY: $(MAKECMDGOLAS) build
.DEFAULT_GOAL := help

PROJECT_NAME=golang-regex

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help:		## Display this help.
	@printf -- "${FORMATTING_BEGIN_BLUE} ${PROJECT_NAME} ${FORMATTING_END}\n"
	@printf -- "\n"
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n make <target>\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "%-15s %s\n", $$1, $$2 } /^##@/ { printf "\n%s\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build: test

##@ Development
test:	## run all tests
	@go run github.com/onsi/ginkgo/v2/ginkgo

urls:			## run urls suite
	@go run github.com/onsi/ginkgo/v2/ginkgo --label-filter=urls

phonenumbers:	## run phonenumbers suite
	@go run github.com/onsi/ginkgo/v2/ginkgo --label-filter=phonenumbers

##@ Reporting
reports:		## generate reports
	@go run github.com/onsi/ginkgo/v2/ginkgo --json-report=report.json --junit-report=report.xml --output-dir=reports

allure: clean	## html allure report
	@mkdir allure-results && cp reports/report.xml allure-results && allure generate && echo "allure open"

clean:			## delete allure directories
	@rm -fr allure-results allure-report
# end
