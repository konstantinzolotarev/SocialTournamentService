FROM golang:1.8

ADD ./ ./

EXPOSE 8000

ENV PORT 8000
ENV USER GameRole
ENV PASSWORD gamemaster
ENV DATABASE game
ENV ADDR 127.0.0.1:5432

CMD ["go", "run", "main.go"]