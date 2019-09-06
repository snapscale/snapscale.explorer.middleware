# step 0
# compile
FROM golang
COPY / /file
WORKDIR /file
RUN /file/auto

# step 1
# server
FROM golang
LABEL Author="Miguel Cao(miguel.cao@xeniro.io)"
COPY --from=0 /file /file

EXPOSE 8089
EXPOSE 8090

CMD ["/file/output/target"]