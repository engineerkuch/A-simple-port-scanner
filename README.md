## A-simple-port-scanner

#### Description:
 - A basic port scanner which checks if a port is opened, closed, or filtered.
 
 This scanner can be improved in many ways, but nmap and several other tools are available. In short, I wrote this script some months ago for learning purpose. 
 Nevertheless, for anyone willing to extend it, think of making it use Rotating Proxies. And with that, the code can be re-wired for concurrent scanning and hence 
 scan large list of servers on multiple ports. It would be nice to also set some kind of delay before we consider the network closed. bla bla bla :-|. 
 Boring, right? Exactly.
 
 ## To run the code:
 - Install (Go)[https://go.dev/dl/]
 - Clone the project
 - Change the IP and port range in the code;
 - run:
   - $go run main.go
   
