FROM golang:1.14-alpine AS build
RUN apk add build-base tzdata
# Configure Timezone
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && echo "Asia/Jakarta" > /etc/timezone
RUN mkdir -p /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go mod verify
RUN go build -o backend-api /app/cmd

FROM alpine
RUN mkdir -p /app
WORKDIR /app
COPY --from=build /app/backend-api .
EXPOSE 8080
CMD ["./backend-api"]
