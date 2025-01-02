FROM golang:alpine AS builder

WORKDIR /src

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go


FROM alpine:latest

# Set the Current Working Directory inside the container
#WORKDIR /app
# Install tzdata and set the timezone
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" > /etc/timezone

# Set the environment variable for timezone
ENV TZ=Asia/Bangkok

WORKDIR /

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /src/main .

# Copy the config file to the desired location
# COPY --from=builder /src/.development ./.development

# Expose port 8090 to the outside world
EXPOSE 3000

# Command to run the executable
ENTRYPOINT ["./main","buildImage"]

# Command to run the executable
#ENTRYPOINT ["./main"]

#CMD ["go","run","main.go"]
