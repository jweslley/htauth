PROGRAM=htauth
VERSION=0.1.0
LDFLAGS="-X main.programVersion=$(VERSION)"

all:
	@for os in linux darwin; do \
		for arch in 386 amd64; do \
			target=$(PROGRAM)-$$os-$$arch-$(VERSION); \
			echo Building $$target; \
			mkdir $$target;         \
			cp ./README.md ./LICENSE $$target; \
			GOOS=$$os GOARCH=$$arch go build -ldflags $(LDFLAGS) -o $$target/$(PROGRAM) ; \
			tar -zcf $$target.tar.gz $$target; \
			rm -rf $$target;                   \
		done                                 \
	done

clean:
	rm -rf *.tar.gz
