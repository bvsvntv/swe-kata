# Infra

## PostgreSQL

### Create a Docker Network

Create a dedicated Docker network for connection and communication between containers.

```sh
docker network create swe-kata-net
```

### Connect Existing Containers

Attach the PostgreSQL primary and replica db containers to the newly created network.

```sh
docker network connect swe-kata-net pg-swe-kata
docker network connect swe-kata-net pg-swe-kata-replica
```

### Verify the Network

Inspect the network to confirm that both containers are attached successfully.

```sh
docker network inspect swe-kata-net
```

## Redis

Connect using `redis-cli`
```sh
redis-cli -h localhost -p 16379 -a ro0T
```
