# begin with the docker default image that uses the latest golang installed
FROM golang

# pull down the current api from github into the default location for the golang image
ADD . /go/src/github.com/velomatrix/surge/api

# install the revel framework and command line tool
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

# launch our application
RUN revel run github.com/velomatrix/surge/api &

# export port 9000 to the outside world
EXPOSE 9000
