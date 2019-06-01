FROM busybox:latest

RUN mkdir /apps
ADD ./bin/linux_386 /apps
WORKDIR /apps
CMD ["/apps/app"]
EXPOSE 8080
