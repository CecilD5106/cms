FROM golang:alpine as builder
# Create a build directory on builder
RUN mkdir build
# Make the build directory the work directory
WORKDIR /build
# Copy the application into the build directory
COPY . /build/
# Build the executable from the code in the directory
RUN CGO_ENABLED=0 go build -a -installsuffix cgo --ldflags "-s -w" -o /build/main
# Designate alpine as the base container
FROM alpine
# Create an app directory for the application
RUN mkdir app
# Make the app directory the work directory
WORKDIR /app
# Create a non root user to run the application
RUN adduser -S -D -H -h /app appuser
# Switch to the non root user
USER appuser
# Copy the application and other files to the app directory
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
# Expose port 8080 for the application to listen on
EXPOSE 8080
# Set the executable to run when the container starts
ENTRYPOINT ["./main"]