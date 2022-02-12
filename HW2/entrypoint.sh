#!/bin/sh
/etc/init.d/postgresql start &&\
cd IoT_CloudAndDistributed &&\
cd HW2 &&\
cd venv &&\
su - postgres
psql --command "CREATE USER sem WITH PASSWORD 'sem';" &&\
psql --command "CREATE DATABASE sem;" &&\
exit 
export USER_DB_NAME=sem &&\
export USER_DB_PASSWORD=sem &&\
export DB_NAME=sem &&\
export DB_HOST=sem &&\
export LC_ALL=C.UTF-8  &&\
export LANG=C.UTF-8 &&\
export FLASK_APP=HW2.py &&\
flask run &