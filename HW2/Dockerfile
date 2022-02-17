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



#Порт который внутри докера слушается вашим решением
EXPOSE 8080

#Переменные окружения

#Слдущая строка нужна чтобы репозиторий кклонировался всегда при docker build даже если ранее данный шаг был закеширован
#ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache
#COPY init.sql /docker-entrypoint-initdb.d/
#Клоним свой репозиторий
COPY src src
COPY entrypoint.sh .
RUN cd src &&\
 go get -u &&\
 go mod tidy 
RUN cd src &&\
 go build 



#Делаем скрипт исполняемым
RUN chmod +x entrypoint.sh

#Задаём точку и условия входа
ENTRYPOINT ["/entrypoint.sh"]
CMD ["run"]