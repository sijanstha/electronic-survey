#!/bin/sh

set -e
COMMAND=$@

echo 'Waiting for database to be available...'
maxTries=10
while [ "$maxTries" -gt 0 ] && ! mysql -h "$DATA_SOURCE_HOST" -P "$DATA_SOURCE_PORT" -u"$DATA_SOURCE_USER" -p"$DATA_SOURCE_PASSWORD" "$DB_NAME" -e 'SHOW TABLES'; do
    maxTries=$(($maxTries - 1))
    sleep 3
done
echo
if [ "$maxTries" -le 0 ]; then
    echo >&2 'error: unable to contact mysql after 10 tries'
    exit 1
fi

exec $COMMAND