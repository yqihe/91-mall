#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/user_admin"
exec "$CURDIR/bin/user_admin"
