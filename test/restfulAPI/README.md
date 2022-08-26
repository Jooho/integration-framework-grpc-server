# RESTFUL API TEST
File extenstion `http` contains restful api url and sample data. Each test api can be called by rest client tool. In order to execute test api, you should install rest client in visual studio code.

## Tool
Visual Studio Code Extension: REST Client(humao.rest-client)

## Create Test Data
~~~
./sample-data/setup.sh
~~~
## How-to
Open a http file(storage.http) in this folder and click `Send Request` string above URL.
~~~
Send Request  <-- click
GET http://localhost:8000/api/v1/storage HTTP/1.1
~~~
