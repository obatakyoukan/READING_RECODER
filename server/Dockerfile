# 開発環境
FROM golang:1.11

WORKDIR /api
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh

RUN apt-get update && apt-get install -y mariadb-client
RUN go get -u  github.com/gin-gonic/gin \
  && go get  github.com/jinzhu/gorm \
  && go get  github.com/jinzhu/gorm/dialects/mysql \
  && go get  github.com/go-sql-driver/mysql \
  && go get  github.com/gin-contrib/cors \
  && go get  github.com/gin-gonic/contrib/static 



CMD ["sh", "./start_app.sh"] 
