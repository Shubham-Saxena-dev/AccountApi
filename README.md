# AccountApi

Basic CRUD operations using GO/GIN rest framework and mongodb. Also the app is dockerized.
Run:

docker-compose up

*******************************
#Create account
POST: /account

Body: {
    "name" : "anyName",
    "email": "anyEmail",
    "zip_code" : 123456
}

Status Code: 200
Response : {
    "_id": "60742bf2ccd99cfe6e4357f1"
}

_id is unique

********************************
#Get all accounts
GET: /account

Status Code : 200
Response:

{
    "_id": "60742bf2ccd99cfe6e4357f1",
    "name" : "anyName",
    "email": "anyEmail",
    "zip_code" : 123456
}

********************************
#Get particular account
GET: /account/:id

Status Code:200
Response:

{
    "_id": "60742bf2ccd99cfe6e4357f1",
    "name" : "anyName",
    "email": "anyEmail",
    "zip_code" : 123456
}

********************************
#Update Account

PATCH: /account/:id

Body: 
{
 "name" : "anyName",
 "zip_code" : 123456
}

Status Code:200
Response:

{
    "_id": "60742bf2ccd99cfe6e4357f1",
    "name" : "anyName",
    "email": "anyEmail",
    "zip_code" : 123456
}

********************************
#Delete Account

DELETE: /account/:id

Response Status Code : 204 No Content
