# Api

### Database setup
- Create a keyspace.

```
CREATE KEYSPACE data
WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};
```
- Create a table. Enter database first by `use data;`

```sql
CREATE TABLE locations(
    loc_x float,
    loc_y float,
    loc_z int,
    user_hash text,
    createdAt timestamp,
    PRIMARY KEY(user_hash, createdAt)
);
```
- Insert testing data.

```sql
INSERT INTO locations (loc_x, loc_y, loc_z, user_hash, createdAt) VALUES (123, 321, 2, 'hell yeah', '2015-05-03 13:30:54.234');
```

### Project setup

1. After starting the server you should now be able to make requests to ```<ip>:8080/data```. For example ```curl --header "Limit: 2" 192.168.1.23:8080/data ```
