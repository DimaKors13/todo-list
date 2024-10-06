FROM ubuntu:latest

ENV TODO_PORT=7540 \
    TODO_DBFILE="./database/scheduler.db"

RUN apt-get update

WORKDIR /todo-list

EXPOSE $TODO_PORT

COPY web ./web
COPY ./builds/todo_app ./cmd/

RUN mkdir database

CMD ["cmd/todo_app"]