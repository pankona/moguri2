#!/bin/bash -e

if [ -n "$1" ]; then
    curl -X POST http://localhost:3000/interact -d '{"action_num": '${1}'}'
fi
curl -X GET http://localhost:3000/current_interaction
