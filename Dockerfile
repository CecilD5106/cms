FROM golang:alpine as builder

RUN mkdir build

WORKDIR /build

COPY . /build/

RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /build/main

FROM alpine

RUN mkdir app

WORKDIR /app

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/
COPY ./form/Edit.html /app/form/Edit.html
COPY ./form/Footer.html /app/form/Footer.html
COPY ./form/Header.html /app/form/Header.html
COPY ./form/Index.html /app/form/Index.html
COPY ./form/Menu.html /app/form/Menu.html
COPY ./form/New.html /app/form/New.html
COPY ./form/NewUser.html /app/form/NewUser.html
COPY ./form/Show.html /app/form/Show.html
COPY ./form/UserList.html /app/form/UserList.html
COPY ./img/JB.jpg /app/img/JB.jpg
COPY ./img/JL.png /app/img/JL.png

EXPOSE 8080

ENTRYPOINT ["./main"]