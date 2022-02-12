#!/bin/sh
#/etc/init.d/postgresql start &&\

service postgresql start &&\
pg_ctl listen_addresses '*' &&\
rm -R IoT_CloudAndDistributed &&\
git clone https://github.com/Farank338/IoT_CloudAndDistributed.git &&\
cd IoT_CloudAndDistributed &&\
cd HW2 &&\
cd src &&\
./hw2