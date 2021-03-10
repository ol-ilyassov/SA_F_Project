
# Step 1 - Build go app
FROM golang as step1

RUN mkdir /app
RUN mkdir /app/article_hub
COPY . /app

ENV CGO_ENABLED=0
WORKDIR /app/cmd/web
RUN go build -o main .

# Step 2 - Optimization
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR ./
COPY --from=step1 /app /app/article_hub

CMD ["/app/article_hub/cmd/web/main"]
