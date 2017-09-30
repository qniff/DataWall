## DataWall 2017 Backend

### Project setup
1. Make sure you clone the project in your ```src``` folder in your ````GOPATH````. If you don't , the project won't be 
able to link the different packages properly. 
If you don't know where your ```GOPATH``` is, 
open a terminal and run ```go env``` and search for ```GOPATH```.

2. Now install all dependencies by going to the root of the repository and running ```go get .```.
3. Make sure you have a cassandra database running.
4. Create a keyspace in cassandra.
```
~ cqlsh
> CREATE KEYSPACE data
  WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};
```
5. Create a table to store the data in
```
> USE data;
> CREATE TABLE locations(
    loc_x float,
    loc_y float,
    loc_z int,
    user_hash text,
    createdAt timestamp,
    PRIMARY KEY(user_hash, createdAt)
  );
```
6. Now define your settings in the ```settings/settings.json``` file by adjusting the following values:
```js
{
  "IpAddress": "127.0.0.1",   // The IP address of your cassandra database
  "Keyspace": "data",         // The keyspace of your cassandra database
  "ApiPort": 8081,            // The port on which the API will run
  "Logging": true,            // Whether to show the logs or not
  "Token": "Your token from https://api.fhict.nl/Documentation/ShowToken"
}

```
7. Now add the following run configurations to your Gogland:
![Api settings](https://i.imgur.com/DhKzvWl.png)
![Data Gatherer settings](https://i.imgur.com/vgBt5JP.png)
