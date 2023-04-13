from golang:1.20-alpine as backend

WORKDIR /go/src/github.com/jfyne/accordopartners.com

COPY go.mod go.sum ./
COPY ./back ./back
COPY ./views ./views
COPY ./cmd ./cmd

RUN go install ./...

from node:14 as frontend

WORKDIR /build

COPY ./package.json .
COPY ./front ./front
COPY ./webpack.config.js .

RUN npm install && \
    npm run-script build

FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/app/

COPY --from=backend /go/bin/accordo .
COPY --from=frontend /build/public ./public
COPY ./front ./front
COPY ./front/img ./public/img

EXPOSE 3000

CMD ["./accordo"]

