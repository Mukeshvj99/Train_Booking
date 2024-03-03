# Train_Booking
It is an application that demonstrates ticket booking using golang,Grpc.

Commands to Build & Run:

1) command to build client executable :
  
   step 1: move inside the client directory (/Train_Booking/client)
   step 2: go build .
   step 3: ./client.exe

2) Command to build Server executable :
 
   step 1: move inside the client directory (/Train_Booking/client)
   step 2: go build .
   step 3: ./server.exe


Functionalities :
 
 1) BookTicket - > Take user details , travelling details as a input and Book a ticket.

 2) GetReceipt - > Take Seatno and Coachno  and Give the Complete Ticket Details. 

 3) GetAllUsers - > Take coachno as input and provide all the user details in that Coachno.

 4) ModifySeat  -> Take seatno and Coacho as input and book a new seat in the train. 

 5) DeleteTicket -> Take seatno and Coachno as input and delete the user ticket from the train.