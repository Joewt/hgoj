FROM golang
MAINTAINER Joewt <joewttx@gmail.com>
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee
RUN go get github.com/shirou/gopsutil/mem
RUN go get github.com/yinrenxin/hgoj
WORKDIR $GOPATH/src/github.com/yinrenxin/hgoj
EXPOSE 8080
CMD ["bee", "run"]