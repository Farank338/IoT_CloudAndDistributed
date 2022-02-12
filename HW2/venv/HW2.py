import flask
from peewee import *
import json
import os


user = os.environ.get('USER_DB_NAME')
password = os.environ.get('USER_DB_PASSWORD')
db_name = os.environ.get('DB_NAME')
host = os.environ.get('DB_HOST')

if user==None:
    print('user')
    exit(1)

if password==None:
    print('password')
    exit(1)

if db_name==None:
    print('db_name')
    exit(1)

if host==None:
    print('host')
    exit(1)


dbhandle = PostgresqlDatabase(
    db_name, user=user,
    password=password,
    host='localhost'
)

class BaseModel(Model):
    """A base model that will use our Postgresql database"""
    class Meta:
        database = dbhandle

class iot_cloud2(BaseModel):
    id = PrimaryKeyField(null=False)
    number = IntegerField(unique=True)    
    



if __name__ == '__main__':
    try:
        dbhandle.connect()
        iot_cloud2.create_table()
    except InternalError as px:
        print(str(px))
        
app = flask.Flask(__name__)
app.run(host='0.0.0.0')

@app.route("/number",methods = ['POST'])
def addnumber():
    #data=flask.request.data
    js =  flask.request.json
    data_number=js['number']
    data_number_less=data_number-1
    res=iot_cloud2.select().where(
        (iot_cloud2.number==data_number)|
        (iot_cloud2.number==data_number_less)
    )
    if len(res)==0:
        res=iot_cloud2.create(number=data_number)
        return {'number':data_number+1}
    else:      
        if len(res)==2:            
            return {'code':500,'message':'Both entries of the number and the number less by one are already in the database'},500        
        if res[0].number==data_number:
            return {'code':500,'message':'This number are already in the database'},500
        else:
            return {'code':500,'message':'Number less by one are already in the database'},500
   