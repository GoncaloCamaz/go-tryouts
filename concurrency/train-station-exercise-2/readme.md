Train station exercise 2
==============

# Problem statement

As we learned in the exercise number 1, the implemented solution only works fine with one passenger.
Thats because channels do not broadcast to every go routine listening.

Therefore, in this exercise I will try to implement some sort of broadcast server or implement the 
observer pattern.

In the folder broadcast-server-example, I have tested a solution that i found to broadcast to every go routine from:
https://betterprogramming.pub/how-to-broadcast-messages-in-go-using-channels-b68f42bdf32e

For the train example, i will implement a similar solution but with some small namming differences. To better understand how the 
broadcast server works, I will call the broadcast server the TrainPositionReport app where I will mock a scenario where the passengers
received a notification in their phone that the train has arrived to a certain station. They can choose to board or not if the train will
go to the station that the passenger wants to go.


# Conclusion:

The broadcast server worked as expected. However, another problem happened that requires attention.
I believe the train is not listening all of the passengers even though I changed the code a bit to try to listen
to every active passenger. It did not work. 

So in a third iteration that i will attack next, i will implement the fan-in pattern in the train side.