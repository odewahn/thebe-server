#
# This will make a super tiny docker image
# See https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
#
# GOOS=linux go build -a -installsuffix cgo -o thebe-server .
# docker build -t thebe-server .
#
FROM scratch
ADD thebe-server /
ADD .env.prod /.env
CMD ["/thebe-server"]
