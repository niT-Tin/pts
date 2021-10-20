.PHONY:
build:
	rm -f base/main user/main paste/main gateway/main
	cd base && go build main.go && cd ../paste && go build main.go && cd ../user && go build main.go && cd ../gateway && go build main.go
