#Stage 1: Crating a Golang image:
FROM golang:1.22.5 as base

WORKDIR /app

# All the dependencies are stored in go.mod file
COPY go.mod .

# This command downloads the additional dependencies required to run application
RUN go mod download

COPY . .

RUN go build -o main .

# Now to build docker img with reduced size and secure in nature!
# Second stage -- Distroless Image
FROM gcr.io/distroless/base

COPY --from=base /app/main .

COPY --from=base /app/static ./static

EXPOSE 8080

CMD [ "./main" ]


