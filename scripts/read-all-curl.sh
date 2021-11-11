#!/bin/bash
#
PORT=$1

if [ $PORT -eq ]; then
  PORT="80"
fi

echo "# PLANET - READ ALL #"
curl -i -X GET "localhost:$PORT/planet"
echo
echo ------------
