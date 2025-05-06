ARG GO_VERSION=1.24

FROM golang:${GO_VERSION}-alpine as builder
RUN apk add make

WORKDIR /builder
COPY . .

ENV CGO_ENABLED=0
RUN make update && \
    make build
RUN ls  -a 

FROM alpine:latest as runner
RUN apk --no-cache add ca-certificates

WORKDIR /app
ARG APP=auth
COPY --from=builder /builder/resources ./resources
COPY --from=builder /builder/$APP ./

RUN touch entrypoint.sh
RUN echo -e "#!/bin/sh \n srv=\${1} \n ./$APP \$srv" >> ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
RUN chmod +x ./$APP
ENTRYPOINT ["./entrypoint.sh"]
CMD ["$APP"]
