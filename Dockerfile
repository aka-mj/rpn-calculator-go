FROM alpine:3.19.1 as build
RUN apk update && apk add --no-cache \
	build-base \
	go \
	git
WORKDIR /app
COPY go.mod .
COPY main.go .
COPY stack.go .
COPY eval_test.go .
RUN go mod tidy
RUN go build -o rpncalc && go test -v && go test -bench=.

FROM alpine:3.19.1
COPY --from=build ./app/rpncalc ./app/
ENTRYPOINT ["./app/rpncalc"]
