docker network create mynetwork
docker run --name postgres --network mynetwork
