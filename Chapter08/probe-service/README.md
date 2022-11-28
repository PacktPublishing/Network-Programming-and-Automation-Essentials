# To build
`docker build . –t probe-service`

# To run single service
`docker run -d -p 9900:9900 probe-service`

# To run multiple services
`docker compose up -d`

# To run docker swarm in a leader host
`docker swarm init`

# To add docker swarm nodes
`docker swarm join --token <token-number> <leader IP>:<leader PORT>`
