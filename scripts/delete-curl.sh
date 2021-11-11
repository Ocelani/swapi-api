#!/bin/bash
#
PORT=$1
ID=$2

if [ $PORT -eq ]; then
  PORT="80"
fi


if [ $ID -eq ]; then
  ID="1"
fi

echo "# PLANET - DELETE #"
echo "delete id: $ID"
curl -i -X DELETE "localhost:$PORT/planet/$ID"
echo
echo ------------
