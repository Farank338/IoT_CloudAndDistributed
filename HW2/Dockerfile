FROM debian:stretch-slim



RUN apt-get update &&\
apt update &&\
apt upgrade -y 

RUN apt-get install sudo -y
RUN apt -y install postgresql postgresql-contrib 
RUN apt -y install postgresql-common

RUN apt-get install systemd -y ;
RUN sudo service postgresql start &&\
echo "CREATE USER sem WITH PASSWORD 'sem';" | sudo -u postgres -i psql &&\
echo "CREATE DATABASE sem;" | sudo -u postgres -i psql


RUN apt install git -y 
RUN git clone https://github.com/Farank338/IoT_CloudAndDistributed.git
EXPOSE 8080
ENV USER_DB_NAME=sem 
ENV USER_DB_PASSWORD=sem 
ENV DB_NAME=sem 
ENV DB_HOST=127.0.0.1 
RUN chmod +x IoT_CloudAndDistributed/HW2/entrypoint.sh
ENTRYPOINT ["IoT_CloudAndDistributed/HW2/entrypoint.sh"]
CMD ["run"]