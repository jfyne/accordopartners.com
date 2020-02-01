from golang as backend

WORKDIR /go/src/github.com/jfyne/accordopartners.com

COPY go.mod go.sum ./
COPY ./back ./back
COPY ./views ./views
COPY server.go ./server.go

RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o accordo .

from node as frontend

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

