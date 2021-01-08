FROM golang:1.15-alpine AS build

ENV APP=./app/cmd
ENV BIN=/bin/opus_spatium
ENV PATH_ROJECT=${GOPATH}/src/opus_spatium
ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOFLAGS=-mod=vendor
ARG VERSION
ENV VERSION ${VERSION:-0.1.0}
ARG BUILD_TIME
ENV BUILD_TIME ${BUILD_TIME:-unknown}
ARG COMMIT
ENV COMMIT ${COMMIT:-unknown}

WORKDIR ${PATH_ROJECT}
COPY . ${PATH_ROJECT}

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags -a -o ${BIN} ${APP}

FROM alpine:3.12 as production
COPY --from=build /bin/opus_spatium /bin/opus_spatium
ENTRYPOINT ["/bin/opus_spatium"]