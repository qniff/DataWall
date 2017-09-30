## DataWall 2017 Backend

### Project setup
1. Make sure you clone the project in your ```src``` folder in your ````GOPATH````. If you don't , the project won't be 
able to link the different packages properly. 
If you don't know where your ```GOPATH``` is, 
open a terminal and run ```go env``` and search for ```GOPATH```.

1. Now install all dependencies by going to the root of the repository and running ```go get .```.
1. Make sure you have a cassandra database running and that you know the **IP** and the **Keyspace**
1. Now define your settings in the ```settings/settings.json``` file by adjusting the following values:
```js
{
  "IpAddress": "127.0.0.1",   // The IP address of your cassandra database
  "Keyspace": "data",         // The keyspace of your cassandra database
  "ApiPort": 8081,            // The port on which the API will run
  "Logging": true,            // Whether to show the logs or not
  "Token": "Your token from https://api.fhict.nl/Documentation/ShowToken"
}

```
