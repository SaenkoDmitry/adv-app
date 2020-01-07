FROM postgres:11

ENV LANG en_US.utf8

ARG PGDATA
ENV PGDATA "$PGDATA"

VOLUME "$PGDATA"

WORKDIR "$PGDATA"

USER postgres

EXPOSE 5432

RUN pg_ctl -D "$PGDATA" -l logfile initdb
