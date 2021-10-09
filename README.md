# golang-architectures
Trying different types of architecture using Golang

## The Layered Architecture
The three independent layers are **delivery**, **use-case** and **datastore**.

### Important note!
The text that you're about to read comes from [Mithali R Shetty](https://medium.com/@r.mithalishetty) from [Medium](https://medium.com) checkout his profile!

### Delivery Layer
The delivery layer will recive the request and parse anything that is required from the request. It calls the use-case layer,
ensures that the response is the required format and writes it to the response writer.

### Use-Case Layer
The use-case layer does the business logic that is required for the application. This layer will communicate with the datastore
layer. Before and after calling the datastore layer, it applies the business logic that is required.

### Datastore Layer
The datastore stores the data. It can be any data storage. The use-case layer is the only layer that communicates with datastore.
This way each layer con be tested independently without depending on the other.

Since each layer is independent of each other, if the application grows to have gRPC support, only the delivery layer will change.
Datastore and use-case layer will remain the same. Even when there is a change in datastore, the entire application need not change.
Only the datastore layer will change. This way, it is easy to isolate any bugs, maintain the code and grow the application.