FROM alpine
WORKDIR /Initial
RUN mkdir -p /msproject/project-project/logs/debug
RUN mkdir -p /msproject/project-project/logs/info
RUN mkdir -p /msproject/project-project/logs/error
COPY ./target/project-project .
COPY ./config/config-docker.yaml .
RUN  mkdir config && mv config-docker.yaml config/config.yaml
EXPOSE 8081
ENTRYPOINT ["./project-project"]