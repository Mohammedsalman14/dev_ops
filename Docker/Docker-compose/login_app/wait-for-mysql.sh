#!/bin/sh

# Wait until MySQL is ready
until mysql -h mysql -u root -prootpass -e "SELECT 1" > /dev/null 2>&1; do
  echo "Waiting for MySQL to be ready..."
  sleep 2
done

# Start the app
npm start
