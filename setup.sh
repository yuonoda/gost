#!/bin/bash

set -eu

readonly DBFILE_NAME="my.db"

# Create DB file
if [ ! -e ${DBFILE_NAME} ];then
  echo ".open ${DBFILE_NAME}" | sqlite3
fi

# Create DB Tables
echo "creating tables..."
sqlite3 ${DBFILE_NAME} "
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS users(\
	id TEXT PRIMARY KEY NOT NULL,\
	name TEXT NOT NULL\
);
"
