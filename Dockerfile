FROM centos:latest

RUN yum update

WORKDIR /app

COPY ./link /app/link

CMD ["/app/link"]
