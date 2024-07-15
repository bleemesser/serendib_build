#!/bin/sh

# if --stop is passed, stop the service otherwise start it
if [ "$1" == "--stop" ]; then
#   sudo systemctl stop pocketbase.service
    echo "Stopping service..."
elif [ "$1" == "--start" ]; then
    echo "Starting service..."
#   sudo systemctl start pocketbase.service
fi

# if no arguments are passed, print the status of the service and a usage message
if [ "$1" == "" ]; then
  echo "Usage: ./run.sh [--start | --stop]"
#   sudo systemctl status pocketbase.service
fi