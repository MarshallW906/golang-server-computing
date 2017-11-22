PKGNAME=github.com/MarshallW906/golang-server-computing/negroni-gbk/gbk
GOCMD=go
GOTOOL=$(GOCMD) tool
GOTEST=$(GOCMD) test
COVERFILE=cover.out
TAGSFILE=tags
TESTCOVER=$(GOTEST) -coverprofile $(COVERFILE)
GOCOVER=$(GOTOOL) cover -func=$(COVERFILE)

clean:
	rm -f $(COVERFILE) gbk/$(COVERFILE) gbk/$(TAGSFILE)

test:
	$(GOTEST) $(PKGNAME)

testcover:
	$(TESTCOVER) $(PKGNAME)
	$(GOCOVER)