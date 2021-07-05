FROM alpine:3.14.0
RUN apk update && apk upgrade
COPY todo_service /tmp/app
EXPOSE 90
CMD ["/tmp/./app"]