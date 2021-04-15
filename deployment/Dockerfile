FROM golang:1.16

RUN	DEBIAN_FRONTEND=noninteractive apt-get update -y && \
	apt-get install -yq --no-install-recommends \
	   git python3 python3-pip curl

RUN  	cd /tmp && \
	git clone https://github.com/biljith/MarkMyLink.git && \
	mv MarkMyLink/ /app && \
	cd /app 

ENV	FLASK_PORT 8080

WORKDIR /app
# CMD	[ "/bin/sh", "-x", "/app/app.sh" ]
ENV DATABASE_URI "mongodb+srv://diddlz:terimarzi69@fse.l0c9p.mongodb.net/BookmarkDB?retryWrites=true&w=majority"

CMD ["go", "mod", "tidy"]
CMD ["go", "mod", "vendor"]
CMD ["go", "run", "main.go"]
