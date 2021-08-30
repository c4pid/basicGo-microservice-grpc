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

![create-customer.png](./create-customer.PNG)_
<br />
![update-customer.png](./update-customer.PNG)_
<br />
![changepass-customer.png](./changepass-customer.PNG)_
<br />

## API Flight
![flight-create.png](./booking-create.PNG)_
<br />
![flight-update.png](./flight-update.PNG)_
<br />
![flight-search.png](./flight-search.PNG)_
<br />


## API Booking
![booking-create.png](./booking-create.PNG)_
<br />
![booking-view.png](./booking-view.PNG)_
<br />
![booking-cancel.png](./booking-cancel.PNG)_
<br />

