#!/bin/bash
#
PORT=$1

if [ $PORT -eq ]; then
  PORT="80"
fi

echo "# PLANET - CREATE #"
curl \
  -H 'Content-Type: application/json' \
  -i \
  -X POST \
  -d '{
  "id": "618890b936ed15c0d9e4745d",
  "name": "PlanetX",
  "climate": "hot",
  "terrain": "desert"
}' \
  "localhost:$PORT/planet"
echo
echo ------------
