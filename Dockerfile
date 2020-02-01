from golang:1.12 as backend

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && \
    chmod +x /usr/local/bin/dep

WORKDIR /go/src/github.com/jfyne/accordopartners.com

COPY Gopkg.lock Gopkg.toml ./

RUN dep ensure -vendor-only

COPY ./back ./back
COPY ./views ./views
COPY server.go ./server.go

RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o accordo .

from node:6 as frontend

WORKDIR /build

COPY ./package.json .
COPY ./front ./front
COPY ./webpack.config.js .

RUN npm install && \
    npm run-script build

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/app/

COPY --from=backend /go/src/github.com/jfyne/accordopartners.com/accordo .
COPY --from=frontend /build/public ./public
COPY ./front ./front
COPY ./front/img ./public/img

EXPOSE 3000

CMD ["./accordo"]

