**Golang RPC Example**

A simple distributed application for vegetable sales system using Go language and Remote Procedure Calls(RPC)

This repository contains a client and a server which communicates using rpc. Clients can use server functions
to do the following tasks.

1. Receive a list of all available vegetables and display
2. Get the price per kg of a given vegetable and display
3. Get the available amount of kg of a given vegetable and display
4. Send a new vegetable name to the server to be added to the server file
5. Send new price or available amount for a given vegetable to be updated in the server file

**Prerequisite**

You need to have Go installed in your machine.
Tested version - 1.13

**How to run**

1. Clone the project in to your local machine.
2. Open up a terminal and execute the command `cd <cloned-path>/golang-rpc-example/server` 
   by replacing the <cloned-path> placeholder.
3. Execute the command `go run server.go` , This will start up the server with port 9000.
4. Open up a new terminal and execute the command `cd <cloned-path>/golang-rpc-example/client`
   by replacing the <cloned-path> placeholder.
5. Execute the command `go run client.go` , This will start up the client.