#/bin/bash
echo "START RUNNING CONTAINER"
docker-compose build
docker-compose down
docker-compose up -d --remove-orphans
echo "DONE RUNNING"