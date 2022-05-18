FROM golang:1.18.2-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
#RUN go mod tidy
RUN go mod download
COPY interface_adapters ./interface_adapters
COPY framework_drivers ./framework_drivers
COPY enterprise_business_rules ./enterprise_business_rules
COPY application_business_rules ./application_business_rules
COPY main.go ./

ARG MONGO_DATA_BASE_URL
ENV MONGO_DATA_BASE_URL=${MONGO_DATA_BASE_URL}

RUN go build -o /order-manager
##
## Deploy
##
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /order-manager /order-manager
EXPOSE 1323
USER nonroot:nonroot
ENTRYPOINT ["/order-manager"]
