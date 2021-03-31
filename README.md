## How to run

```sh
$ docker-compose up

$ curl http://127.0.0.1:8000/v1/say_hello
$ curl http://127.0.0.1:/v1/ping_google
```

## Modes

Kong can run with two modes:

### Database mode

It stores the configuration as well as some state into a Postgres (or Cassandra) database.
Configuration is applied:
- through API calls
- via an administration UI
- through applying a `kong.yml` config file (more below)

This is the default option.

### Declarative configuration

Actived by defining 2 env vars:
```sh
KONG_DATABASE=off
KONG_DECLARATIVE_CONFIG=kong.yml
```

When this option is chosen, Kong is still mostly stateless, although you can still override the in-memory config through API calls (http, CLI, admin UI). Changes will be applied in memory only and be reset next time you restart it.  

Note that Kong only loads `kong.yaml` once at startup, any changes to the configuration require either a reboot,
or use the CLI tool to run `deck sync` (see [decK](https://github.com/Kong/deck))

## General thoughts

Pros:
- super easy to get started, to configure
- documentation is quite good (although it's sometimes not too easy to find things)\
- performance is good

Cons:
- Declarative configuration looks like the option we want to use (it works similarly as everything else we have and avoids state) but seems to be [lacking options](https://docs.konghq.com/gateway-oss/2.3.x/db-less-and-declarative-config/#partial-compatibility) the DB mode has. Maybe we can use DB and inject config at startup if that proves an issue?
- We don't use it in Beat yet
- It might need Postgres and/or Redis to be fully functional and optimal

## Latency

Latency added by Kong is < 5 ms (with very simple tests, running ±20 queries and looking at the info returned by Kong).

When booting Kong, it takes an initial 20 seconds for the first few queries to `hello-service` to return - this is probably due to a DNS timeout one way or another. This doesn't happen when querying services outside Docker-Compose (the `google` example works perfectly from the first try), and stops happening after the first few queries succeed. It then happens again after some time. This might have to do with DNS resolution (and a short TTL makes it happen again later). This might or might not happen in k8s.

If it does happen, I can see two ways (both are complementary) to help with that:
- Have a ready-check that pings the other services, and only return "ready" when the cache is built
- Use KONG_DATABASE mode to persist the cache

## Create an empty configuration

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