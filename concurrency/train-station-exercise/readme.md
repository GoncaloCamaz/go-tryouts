Train station exercise
==============

# First steps

In this folder i will implement a problem that takes use of some of the go concurrency mechanisms.
The idea behind the execise is to have a train that follows a path indefinitly. The train has a capacity
for a certain number of passangers and should hop from station to another, letting passengers in and out.
A passenger has an origin and a destination and should leave when arrives at the destination.

In short:
- Implement a train that runs from station to station and lets passengers in and out
- Implement passengers that have an origin and destination and try to enter in the train

Kinda successfully, i have implemented the solution for the problem. However, with this solution, only 
one passenger can enter the train at a time.
This happens because even though I can have several passengers reading from the current station BoardChannel,
channels do not broadcast to every single go routine. In other words, only one passenger will hear
that the train has arrived at a certain station and that passenger might enter or not depending if the train
will take him until its destination.

To solve this issue we have two alternatives.
One solution would consider having as many channels as there are passengers. however, this 
solution does not seem that funny to me. 

Therefore, i will try to solve this issue with a broadcast server and implement a solution in which
each passenger will be an observer.

In short:
- Channels do not broadcast, only one go routine will receive its content
- we need wait groups to make sure all passengers complete their journey, however 
in this solution only one passenger will hehehe