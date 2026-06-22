MAKEFLAGS += --always-make

all: test serve

clean:
	rm -rf allure-results allure-report

serve:
	allure serve

test:
	-go test . -v -tags example -count=1 -race -timeout=4s
