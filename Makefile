OUT="go-nl"
INST="$(GOPATH)/bin/$(OUT)"

build: $(OUT)

$(OUT):
	go build go-nl.go

$(INST):
	cp $(OUT) $(INST)

install: $(OUT) $(INST)

clean:
	rm $(OUT)
