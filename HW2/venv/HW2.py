import flask
from peewee import *
import json

user = 'sem'
password = 'sem'
db_name = 'hw2'

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

@app.route("/number",methods = ['POST'])
def addnumber():
    data=flask.request.data
    js = json.loads(data)
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
   