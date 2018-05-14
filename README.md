## Counzl
Counzl is a programs that consists of a client and a server that works like a template for interactive programming in golang. We have included buntDB for easy database management, TCP with TLS for safe communication between server and client, and authentication (sign up (online & local) and login). The client uses iShell (CLI) as the interface, because it is easy to implement new commands and function.  

### TCP, TLS
As mentioned above we use TCP (**t**ransmission **c**ontrol **p**rotocol)  with TLS (**t**ransport **l**ayer **s**ecurity) to communicate with the server application on UH-IaaS. We use TCP to make sure that the network packets is delivered to the server socket untampered. TLS makes the packages encrypted, and uses keys and certificates to verify the client.

We used Wireshark to analyze the traffic to make sure TLS worked properly and that the packets was encrypted; and they were (as expected).

## Tech/framework used
First of all, thanks to all the developers that contributed to the git repositories listed underneath. Your work are much appreciated as they have helped us to focus on other problems than to build the CLI itself, a database managing tool, password hashing, etc.

**Link to repositories:**<br>
* [Tidwall's BuntDB for database tool](https://github.com/tidwall/buntdb) <br>
* [Abiosoft's iShell for CLI](https://github.com/abiosoft/ishell)<br>
* [Golang's bcrypt package for password hashing](golang.org/x/crypto/bcrypt)<br>
* [Golang's terminal package for invisible password input]()

## System architecture
The system architecture for our program is quite simple. It can appear a little bit perplex, but this is because we use UH-IaaS to host the server.  
![System Architecture](https://github.com/BadNameException/Counzl/blob/master/sys_architecture.jpg)

## Our own 'Utilities' package 
This package consists of 'homemade' functions that we have realized were going to be applied in the project more than once.

### cmd
As you have probably guessed this package is related to the terminal/CLI. To make the design a bit more interesting we have implemented a function to change the color of a chosen string and also a loading indication. We also have an empty go file for executing bash script. We have not used this function yet, but the idea was to execute a script for generating keys and certificates. 

### converter 
This package is the most used. You use it to get the current time in our timezone (Europe/Oslo) in a human readable format, and convert string to boolean, int and string to hexadecimal and string to int. 
The most used function is removing and/or adding 0x0a from/to a string. For example an argument from stdin will always have a '\n' attached to it (finishes the argument with return/enter), and the same phenomenom occurs when you write a string to the server from the client and vice versa.

### crypto
This package has a function for hashing a password, and another function to compare the hashed password with a hashed argument (authentication). 

## Motivation
Our motivation behind this project is to learn more about golang development for interactive programs for the course IS-213 at the University of Agder. We wanted to create a program that should be versatile and easy to contribute to (or just use it as a template for your own program).

## Contribute
You can contribute to our project right here on GitHub! 
Follow our [contributing guideline](https://github.com/BadNameException/Counzl/blob/master/CONTRIBUTING.MD) to get started.

## Credits
Give proper credits. This could be a link to any repo which inspired you to build this project, any blogposts or links to people who contrbuted in this project. 

## License

MIT
