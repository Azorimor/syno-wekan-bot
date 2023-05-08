get-wekan:
	curl https://raw.githubusercontent.com/wekan/wekan/master/docker-compose.yml -o wekan.docker-compose.yml

start-wekan:
	docker-compose -f wekan.docker-compose.yml up -d

pull-wekan:
	docker-compose -f wekan.docker-compose.yml pull

stop-wekan:
	docker-compose -f wekan.docker-compose.yml stop

remove-wekan:
	docker-compose -f wekan.docker-compose.yml down -v
