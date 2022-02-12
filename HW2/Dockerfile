# Как исходник будем использовать вариант уже собранного докера с Дебианом (Дебиан легок и максимально "голым поставляется")
FROM debian:stretch-slim


#Пусть это и скорее всего будет бессмысленно но обновим сведения о доступных приложения и обновим те что уже установленны
RUN apt-get update &&\
apt upgrade -y 

#Фишка чисто моего кода т.к. я пишу на компилируемом языке то будет хорошим тоном перекомпилить его, а это является подготовкой
#Вставитье здесь установку своего языка/бибилотек необходимых вашему решению
RUN  apt-get install build-essential -y
RUN apt -y install wget
RUN wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin 


#Установка чисто для моего удобства чтобы можно было дальше переключать пользователей максимально легко
RUN apt-get install sudo -y

#Ставлю БД замените на то которое используется у вас
RUN apt -y install postgresql postgresql-contrib 
RUN apt -y install postgresql-common

#Инициализирую БД и создаю в ней таблички для этого ставил судо собтвенно
RUN sudo service postgresql start &&\
echo "CREATE USER sem WITH PASSWORD 'sem';" | sudo -u postgres -i psql &&\
echo "CREATE DATABASE sem;" | sudo -u postgres -i psql &&\
#Меняю настройку через запрос к ранее запущенной бд чтобы она слушала на всех адресах
echo "ALTER SYSTEM SET listen_addresses = '*';" | sudo -u postgres -i psql

#Ставлю гит чтобы скачать своё решение
RUN apt install git -y

#Порт который внутри докера слушается вашим решением
EXPOSE 8080

#Переменные окружения
ENV USER_DB_NAME=sem 
ENV USER_DB_PASSWORD=sem 
ENV DB_NAME=sem
ENV DB_HOST=127.0.0.1

#Слдущая строка нужна чтобы репозиторий кклонировался всегда при docker build даже если ранее данный шаг был закеширован
#ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache

#Клоним свой репозиторий
RUN git clone https://github.com/Farank338/IoT_CloudAndDistributed.git

#Делаем скрипт исполняемым
RUN chmod +x IoT_CloudAndDistributed/HW2/entrypoint.sh

#Задаём точку и условия входа
ENTRYPOINT ["IoT_CloudAndDistributed/HW2/entrypoint.sh"]
CMD ["run"]