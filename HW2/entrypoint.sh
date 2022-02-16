#!/bin/sh

#Фактически это не старт а перезапуск который нужен чтобы ранее сделанная настройка на прослушку всех адресов заработала
service postgresql start &&\

#Обновим репозиторий с решением нужно для того чтобы подкачались свежайшие исходники с кодом
rm -R IoT_CloudAndDistributed &&\
git clone https://github.com/Farank338/IoT_CloudAndDistributed.git &&\
cd IoT_CloudAndDistributed &&\
cd HW2 &&\
chmod +x entrypoint.sh
function is_in_activation {
   activation=$(systemctl status postgresql | grep "Active: active" )  
   if [ -z "$activation" ]; then
      true;
   else
      false;
   fi

   return $?;
}

while is_in_activation network;       
    do true    
    done    


#Запускаем основной срипт
chmod +x run.sh 
./run.sh

