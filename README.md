## A-simple-port-scanner

#### Description:
 - A basic port scanner which checks if a port is opened, closed, or filtered.
 
 This scanner can be improved in many ways, but nmap and several other tools are available. In short, I wrote this script some months ago for learning purpose. 
 Nevertheless, for anyone willing to extend it, think of making it use Rotating Proxies. And with that, the code can be re-wired for concurrent scanning and hence 
 scan large list of servers on multiple ports. It would be nice to also set some kind of delay before we consider the network closed. bla bla bla :-|. 
 Boring, right? Exactly.
 
#### Understanding network handshake workflow:
 - For a port to be open, it follows necesarily that a three-way handshake has occuped. The Client send's a "syn" to the Server, the Server respond's with a "syn-ack", then the Client send's back an 'ack' to the Server, and a connection is established (in other words, the port is opened). Else if the Client received only a "rst" as response, then the port is closed. But if the client receive's no response at all, then the "syn" request sent by the client times-out (meaning, the connection to server is filtered). 
 - As illustated by the image below:
 - ![Untitled Diagram drawio](https://user-images.githubusercontent.com/59922186/146649090-5041f787-19bf-4c70-a6be-f5add0d4dafd.png)

    
 ## To run the code:
 - Install (Go)[https://go.dev/dl/]
 - Clone the project
 - Change the IP and port range in the code;
 - run:
   - $go run main.go
   
