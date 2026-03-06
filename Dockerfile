FROM ubuntu:latest
LABEL authors="egorgusakov"

ENTRYPOINT ["top", "-b"]