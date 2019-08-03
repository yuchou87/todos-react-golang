FROM golang:latest

LABEL maintainer="yuchou87 <yuchou87@gmail.com>"

WORKDIR /app

COPY ./app/release/linux/amd64/ /app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]