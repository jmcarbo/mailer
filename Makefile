build:
	docker run -ti --rm -v $$PWD:/go/src/mailer -w /go/src/mailer golang /bin/bash -c 'go get ./... && go build .'
