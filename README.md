# Counzl_MUX

## **The purpose of our Counzl_MUX**
The purpose of 'Counzl_MUX' is to have a template for a server program in go that should be easy to develop further (by adding more modules) for both beginners and experienced programmers. Currently there is just one module: 'new_user'.

## Running the program
1. Set GOPATH to the library package:
```bash
export GOPATH=absolute/path/to/Counzl_MUX/library
```
Note: if you want this GOPATH to be permanent, add the command above to the .profile (or .bash_profile). This file should be located in the $HOME-folder. Either restart the terminal or execute the following command: 
```bash
source .profile
```

2. Go to the project folder ("Counzl_MUX") and type the following command: 
```bash
go run main.go
```
You should see the following text appear in the terminal: <br>
Launching server on: <br>
Port: 8081 

# Implemented functions Counzl_MUX

## Overview: 
To make the links work I had to write the headings in lower case. The correct function names are listed in this overview. The numbers represents go-files, the bullet point are functions.
1. [users.go](#users) 
* [findHighestID](#findhighestid)
* [PrintAllUsersByID](#printallusersbyid)
2. [user_listener.go](#user_listener)
* [Initialise_User_Listener](#initialise_user_listener)
* [handleConnection](#handleconnection)

## users
This file is almost identical to [user_management.go](https://github.com/BadNameException/Counzl_Client/blob/justanotherbranch/app/modules/users/user_management.go) except for a few variations: 
* A user is stored in the database with the user ID as the key (instead of username like in Counzl_Client), this makes it possible to store multiple users with the same username. 
* An additional db-file that contains just one element; the highest ID that exists in the user-db-file. This makes storing a user much quicker, because we do not have to iterate the user-db to find the next user ID. 
* A user is stored with a timestamp which indicates the time a user was stored. 

I am just going to explain the new functions. newUser, PrintAllUsers, serialize and deserialize are almost identical to CreateUser, PrintLocalUsers, serialize and deserialize in Counzl_Client. The only noticeable differences is that the key are ID (as mentioned earlier) and that newUser returns the new user's ID.

### findHighestID
* Opens the DB-file: u_highest_ID
* Fetches the ID from the db-file; which is the highest ID in the user-db
* If the value is "" (empty) the new ID has value 1. The key is **always** "only_value".
* If the value is 1 or more the new ID gets incremented by 1. Since The value is stored as a string, it has be to converted into an integer before it can be incremented by 1. This happens in converter.StringToInt_plus1 from the utilities-package. 
* The new ID is stored in the db-file (u_highest_ID).
* Returns the new ID.

### PrintAllUsersByID
* Declares a temporary variable 'i' that holds the current value of an ID. It is declared with the value 1.
* A for-loop that fetches the element with the ID = i (first iteration: i = 1, second iteration: i = 2, etc.). 'i' gets incremented by 1 for every iteration. 
* The loop stops when the function returns an error, which means that there are no elements left in the user-db.

## user_listener
Only listens to incoming connections at port 8081. 
### Initialise_User_Listener
* Loads a new keypair which consists of server.crt and server.key.
* Declares a pointer to a TLS configuration, in short this means a slice of the certificates (keypair). 
* Listens to incoming connections and accepts those with the right TLS configuration (correct keypair). When a connection is accepted the handleConnection will be called with the TLS configuration as input parameter.  

### handleConnection
* Consists of a foor-loop that retrieves the username from the connection, removes the line feed, calls newUser to get the new user ID and writes the ID to the connection (so the client gets the ID). 