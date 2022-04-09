FROM golang:1.13
WORKDIR /home/app
COPY . .
RUN apt-get update
RUN apt-get install -y sqlite3
CMD ["tail", "-f", "/dev/null"]