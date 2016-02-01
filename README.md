# CMPE273_Lab-3

## Simple RESTful key-value cache data store

The purpose of this lab is to demonstrate how consistent hashing works and to implement a simple RESTful key-value cache data store. Client implements consistent hashing using md5 hashing function and shards the provided sample data set among the available server nodes. RESTful endpoints are provided to store & retrieve key-value pairs.

### Requirements  
•	Golang latest stable version (I have used go1.5 on Windows)   
•	You can check the official golang releases here: https://golang.org/doc/devel/release.html  

### Installation

#### Installing Go (In case if you don't have it)
•	There are various ways to install go according to the operating system that you’re working on.   
•	All the required files and step-by-step instructions can be found here : https://golang.org/doc/install    

#### Installing Packages
After you have installed the Golang then run the following command      
```
go get github.com/onkarganjewar/CMPE273_Lab-3
```

You will also need to install the httprouter package which can be found here  
```
go get github.com/julienschmidt/httprouter
```

### Usage

Start 3 server instances namely:
* server1.go: http://localhost:3001
* server2.go: http://localhost:3002
* server3.go: http://localhost:3003

Also run the client. The program supports only 2 http calls namely PUT (to store the keys into any one of the server node using consistent hashing) & GET (to retrieve keys stored at any specific server or at all the nodes). Client performs consistent hashing over provided sample data set of 10 keys.

```
go run client.go
```

Now that the connection has been established, open any REST Console or using cURL commands perform following operations such as

  1. PUT http://localhost:3000/keys/{key_id}/{value}

    - Request:
  
      ```http
      PUT http://localhost:3000/keys/1/one
      ```
    - Response: 
  
      ```http
      HTTP 200
      ```
  2. GET http://localhost:3000/keys/{key_id} 

    - Request:
    
      ```http
      GET  http://localhost:3000/keys/1
      ```
    
    - Response:
      ```json
      {
      "key" : 1,
      "value" : "one"
      }
      ```
  3. GET http://localhost:3000/keys

    - Request:

      ```http
      GET  http://localhost:3000/keys
      ```
    
    - Response:
      ```json
        [
          {
            "key" : 1,
            "value" : "one"
          },
          {
            "key" : 2,
            "value" : "b"
          }
        ]
        ```

