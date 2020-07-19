#!/bin/sh

echo "Waiting for MariaDB..."

while ! nc -z ${DB_HOST} ${DB_PORT}; do
    sleep 0.1
done

echo "MariaDB Started"

./monblog_api

echo "Server Started"