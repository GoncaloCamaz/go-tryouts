Udemy course first example
==============

In this example, we will be looking into wait groups and go routines.

The objective is to learn about how to create go routines and avoid getting dead locks.

To do that, we must use wait groups. Every routing we will be creating, we will add it to the wait group.

After executing all go routines, we will wait for all of them to finish. Important to nothe that we must call 
wg.Done at the end of every routine.