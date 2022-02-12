FROM debian:stretch-slim



RUN apt-get update &&\
apt update &&\
apt upgrade -y 

 
RUN apt -y install wget
RUN wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz



RUN  apt-get install build-essential -y

RUN apt-get install sudo -y
RUN apt -y install postgresql postgresql-contrib 
RUN apt -y install postgresql-common

RUN apt-get install systemd -y ;
RUN sudo service postgresql start &&\
echo "CREATE USER sem WITH PASSWORD 'sem';" | sudo -u postgres -i psql &&\
echo "CREATE DATABASE sem;" | sudo -u postgres -i psql &&\
echo "ALTER SYSTEM SET listen_addresses = '*';" | sudo -u postgres -i psql

RUN tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin 

RUN apt install git -y 

ADD "https://www.random.org/cgi-bin/randbyte?nbytes=10&format=h" skipcache
RUN git clone https://github.com/Farank338/IoT_CloudAndDistributed.git
RUN cd IoT_CloudAndDistributed/HW2/src


RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/postgres

EXPOSE 8080
ENV USER_DB_NAME=sem 
ENV USER_DB_PASSWORD=sem 
ENV DB_NAME=sem 

ENV DB_HOST=127.0.0.1 
RUN chmod +x IoT_CloudAndDistributed/HW2/entrypoint.sh
ENTRYPOINT ["IoT_CloudAndDistributed/HW2/entrypoint.sh"]
CMD ["run"]