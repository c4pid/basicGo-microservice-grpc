# ASSIGNMENT BOOKING, BASIC MICROSERVICE, GPRC

## Proto Gen

```shell
protoc -I=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb
```

## Test with grpcui 
```
grpcui -plaintext 127.0.0.1:port
```

## API Customer

![create-customer.png](./create-customer.png)_
<br />
![update-customer.png](./update-customer.png)_
<br />
![changepass-customer.png](./changepass-customer.png)_
<br />

## API Flight
![flight-create.png](./booking-create.png)_
<br />
![flight-update.png](./flight-update.png)_
<br />
![flight-search.png](./flight-search.png)_
<br />


## API Booking
![booking-create.png](./booking-create.png)_
<br />
![booking-view.png](./booking-view.png)_
<br />
![booking-cancel.png](./booking-cancel.png)_
<br />
