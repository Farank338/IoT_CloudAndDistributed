#!/bin/sh

#Фактически это не старт а перезапуск который нужен чтобы ранее сделанная настройка на прослушку всех адресов заработала
service postgresql start &&\

#Обновим репозиторий с решением нужно для того чтобы подкачались свежайшие исходники с кодом
rm -R IoT_CloudAndDistributed &&\
git clone https://github.com/Farank338/IoT_CloudAndDistributed.git &&\
cd IoT_CloudAndDistributed &&\
cd HW2 &&\

#Запускаем основной срипт
chmod +x run.sh 
./run.sh

