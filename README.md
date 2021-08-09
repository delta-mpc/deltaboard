# delta-board-server

## Run with provided mysql compose up
1.docker-compose up -d
2.visit http://localhost:8090
## Run with docker all in one 

1.docker build -t dashboard_in_all .

2.docker run -it -p 8090:8090 dashboard_in_all

3.visit http://localhost:8090
