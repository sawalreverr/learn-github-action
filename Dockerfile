FROM golang:1.22 AS build-stage

WORKDIR /projects/learn-github-action

COPY . . 

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /learn-github-action

FROM gcr.io/distroless/base-debian11 AS build-release-stage

COPY --from=build-stage /learn-github-action /learn-github-action

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./learn-github-action"]