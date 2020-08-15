.PHONY:

build:
	docker build -t market-info --rm .

run:
	docker run -p 3001:80 -it --name='marketinfo' market-info
