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

![create-customer.png](./create-customer.PNG)
<br />
![update-customer.png](./update-customer.PNG)
<br />
![changepass-customer.png](./changepass-customer.PNG)
<br />

## API Flight
![flight-create.png](./booking-create.PNG)
<br />
![flight-update.png](./flight-update.PNG)
<br />
![flight-search.png](./flight-search.PNG)
<br />


## API Booking
![booking-create.png](./booking-create.PNG)
<br />
![booking-view.png](./booking-view.PNG)
<br />
![booking-cancel.png](./booking-cancel.PNG)
<br />

