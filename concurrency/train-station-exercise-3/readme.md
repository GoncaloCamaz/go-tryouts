Train station exercise 3
==============

*** Final steps ***

# Conclusions

In the previous train station implementations, I came accross some concurrency challanges.
On the first iteration of the exercise, only one passenger would board the train because channels
do not broadcast to all go routines. To address this issue, on the second implementation, I implemented 
a broadcast service like the observer pattern for all the go routines representing the passengers to Subscribe
to the train current station report channel. It worked and every go routine received the station.
However! for some reason only one passenger would enter the train and if we have more than one, we will have 
deadlocks...
I changed the code a bit for the train to have an array of active passengers (onboard or waiting for the train). And, 
in order to have all the passenger responses, I added an additional channel to ignore the station report if the passenger
does not want to board or disembark. However, the deadlock problem still persisted. So, after digging a bit, I came 
accross the fan-in pattern where multiple go routines write to a single go routine via a channel. 
To implement this, I will try to implement this pattern but with some refactors, for exemple, merging 
board, disembark and ignore channel into only one channel and add a struct to represent the passenger response.

So I tried to apply a fan in pattern by listening to the channel and appending the results to an array and later read by the train.
However, this solution was making dead locks to happen after one passenger got out of the train.

To fix the issue, i removed the fan in pattern and replaced by an passenger active array and then added a wait group to wait 
for all active passengers to answer. This fixed the issue and now the train exercise is working just fine.