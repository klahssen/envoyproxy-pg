FROM scratch
ADD ./cmd/svc/svc /bin/server
ENTRYPOINT ["/bin/server"]