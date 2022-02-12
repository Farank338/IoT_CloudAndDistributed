FROM debian:stretch-slim



RUN apt-get update &&\
apt update &&\
apt upgrade -y 

RUN apt -y install postgresql &&\
/etc/init.d/postgresql start 

RUN apt install git -y 
RUN git clone https://github.com/Farank338/IoT_CloudAndDistributed.git
EXPOSE 8080
ENV USER_DB_NAME=sem 
ENV USER_DB_PASSWORD=sem 
ENV DB_NAME=sem 
ENV DB_HOST=sem 
RUN chmod +x IoT_CloudAndDistributed/HW2/entrypoint.sh
ENTRYPOINT ["IoT_CloudAndDistributed/HW2/entrypoint.sh"]
CMD ["run"]