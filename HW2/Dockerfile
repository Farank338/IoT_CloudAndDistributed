FROM debian:stretch-slim



RUN apt-get update &&\
apt update &&\
apt upgrade -y &&\
apt install build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libsqlite3-dev libreadline-dev libffi-dev curl libbz2-dev -y &&\
apt install build-essential libssl-dev libffi-dev python3-dev -y &&\
apt-get install wget -y &&\
apt install python3.9 -y &&\
apt install -y python3-pip &&\
apt install -y python3-venv &&\
pip3 install flask

RUN pip3 install peewee &&\
apt install python3-dev libpq-dev -y &&\
pip3 install psycopg2 &&\
apt update 


RUN apt -y install postgresql &&\
/etc/init.d/postgresql start 



USER root
RUN apt install git -y &&\
git clone https://github.com/Farank338/IoT_CloudAndDistributed.git &&\



ENTRYPOINT ["/home/student/IoT_CloudAndDistributed/HW2/entrypoint.sh"]
CMD ["run"]