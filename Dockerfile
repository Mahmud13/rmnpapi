FROM golang:1.10

RUN go get github.com/astaxie/beego 
RUN go get github.com/beego/bee
RUN go get firebase.google.com/go
RUN go get github.com/jinzhu/gorm/dialects/mysql  
RUN go get github.com/jinzhu/gorm 


EXPOSE 8080

CMD ["bee", "run"]
