# nPassword Distributed Password Manager

![Alt text](./demo.png?raw=true "Homepage")

The goal of this project is to create a password manager that can store a user's passwords effectively
with one-way encryption. Now with the whole idea of quantum computers, it might be soon possible to break those hashed passwords. Thus, I am attempting to split the password into chunks
to store them in different files on different machines in order to add an additional layer of protection. 

## Technology

The backend is built using Go. It handles the processing of the requests and passwords and then storing them into the db (currently local files). The frontend is built using React.

## Setup

Clone the repo. Install Go onto your machine https://go.dev/doc/install. Then in a terminal, type ```go version``` to ensure it is installed.
Install Node.js and npm using https://nodejs.org/en/download and then install axios using ```npm install axios```.

## Execution

To run the code, move to the ```nPasswordManager/server``` directory and run the service using ```go run main.go```. In another terminal, go to the
```nPasswordManager/client/npassword``` directory and run the client using ```npm start```

## Future

I am currently working on adding functionality to verify a password, remove and list the hashed passwords. I want to be able to then connect a database to store these different files.

If I had multiple systems, I would have liked to take advantage of Hadoop or Spark to streamline the distribution.

