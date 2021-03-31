## How to run

```sh
$ docker-compose up

$ curl http://127.0.0.1:8000/v1/say_hello
$ curl http://127.0.0.1:/v1/ping_google
```


## Useful stuff

```sh
docker run -d --rm --name kong \
    -e "KONG_DATABASE=off" \
    -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
    -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
    -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
    -p 9000:8000 \
    -p 9443:8443 \
    -p 9001:8001 \
    -p 9444:8444 \
    kong
docker exec -it kong kong config init /home/kong/kong.yml
docker exec -it kong cat /home/kong/kong.yml >> kong.yml
```