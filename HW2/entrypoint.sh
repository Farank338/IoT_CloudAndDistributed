#!/bin/sh
#/etc/init.d/postgresql start &&\
#pg_ctl start
rm -R IoT_CloudAndDistributed &&\
git clone https://github.com/Farank338/IoT_CloudAndDistributed.git &&\
cd IoT_CloudAndDistributed &&\
cd HW2 &&\
cd src &&\
./hw2