Algorithms
==============

This file contains a resume about all the algorithms added in this ./algorithms folder.

# Binary Search

Binary search algorithm tries to find an element in the array. For this algorithm to work,
the array must be sorted.

This are the steps to implement the algorithm:
1 - create a function that receives the array and the target number.
	1.1 - If the implementation is iterative, we can proceed
	1.2 - If its recursive, add an low and high number to the function
2 - while the lowest number is smaller or equal then the highest number calculate the middle element
3 - midElement = low + (high - low) / 2
4 - if the mid element is equal to the target, we found the number
5 - if the mid element is smaller than the target, we can discard the right part of the array
6 - if the mid element is larger than the target, we can discard the left part of the array
7 - dont forget to add -1 to return a value if the element is not found, when low > high

# Depth first algorithm

Depth first algorithm works by visiting the children of a node before visiting the siblings or neighbours
In the case of a binary tree, we will try to go left until we cant proceed any further. 

To implement this with graphs, we need to consider a graph as a map from int to int array where each 
vertex, or key int the map, has a certain connection to another vertexts, the values of the map for that key
graph map[int][]int. We will also use a map to mark the vertexes as visited. map[int]bool

1 - mark the start node as visited and assign its neighbours to a temporary array
2 - for every element in that array, if we have not marked it as equal yet, visit it 
	by calling the recurvie call in which, the start node will be the neighbour

When using trees, we dont need an temporary array because we can either go left or right in the tree

1 - if the tree is null, return because we reached an leaf node
2 - mark the node as visited and call recursevly the function, first on the left and then on the right

# Breadth first algorithm

This algorithm works by visiting all the siblings on a node before visiting the children
Therefore, for this to work, we need to implement a queue system to store the references to each sibling 
and only then go to the children
We can implement this both recursevely or iteratively.

For graphs, we will have the same structure as in the depth first algorithm were we need to use a map
to represent the nodes and its connections

For the iterative version:

1 - start a queue with the first vertex in it
2 - create a for loop which will stop when we dont have any more elements in the queue
3 - inside the for loop, assign the currentVertex to the head of the queue
4 - mark the vertex as visited
5 - remove the head of the queue: queue = queue[1:]
6 - for every neighbour iniside the map key, if the neighbour has not been visited append the neighbour to the queue
	and mark the neighbour as visited

For the recursive version:
1 - mark the vertex as visited
2 - appen the neighbour nodes to the queue
3 - for each member inside the queue, if the member is not visited, remove the first member from the queue
4 - call recursive function starting on the member

# Insertion sort

This is a simple sorting algorithm and not very efficient to sort an array
The way it works is by playing a bit with indexes

1 - create a for loop to go through the array but start the iteration of variable i at 1
2 - store that element in a temporary variable
3 - create a new variable to start at the position before element i ( j = i - 1)
4 - iterate over array using index j, while j >= 0 and arr[j] is bigger than the element in temporary variable
5 - move the previous element to the next position and decrease j
6 - outside the j for loop, assign the arr[j+1] = temp element

for example: 

2 4 5 6 1 3
	||
2 4 5 6 6 3
	||
2 4 5 5 6 3
	||
2 4 4 5 6 3
	||
2 2 4 5 6 3
	||
1 2 4 5 6 3


# Merge sort

This algorithm sorts an array by using the divide and conquer technique. We split the array in smaller
pieces and then merge every piece in the correct order
To implement this algorithm we need three functions

1 - mergeSort to receive an array and call the recursion function 
2 - mergeSortRecursion that receives the array, the start position and the last position
3 - merge to merge the arrays which receives the array, left middle and right position

2.1 - if left < right find the middle position of the array
2.2 - middle = left + ( right - left ) / 2
2.3 - call the function using recursion on the left part of the array until the middle
2.4 - call the function using recursion from the middle + 1 until the right part of the array
2.5 - call the merge function after the recursion

3.1 - receiving the arr, left, middle and right positions, we need to calculate the size of the arrays
3.2 - leftSize = midle - left + 1 and right side, right - middle
3.3 - create a temporary right and left array
3.4 - create a for loop to assign the content of arr from left + i until the size of left
3.5 - create a for loop to assign the content of the arr from the middle + 1 + j to the right temp array
3.6 - create another for loop with i, j and k, where k starts at left position and run until k <= right position
3.7 - if the i < leftSize and tempLeft i <= tempRight j, we assign the arr[k] the content of tempLeft i, else the content of tempRight John
3.8 - we increase the variable of the temp array content we reached

# Quick sort

This algorithm works by selecting a pivot from an array and placing every smaller items on the right
We need four functions for this

1 - to call the recursion by passing the array, start and end position
2 - a recursion function to partition the array, and call recursion on the left and on the right side
3 - a partition function to move the elements
4 - a swap function to swap the elements inside the array

2.1 - while low < high, get the pivot by calling the partition function
2.2 - call recursion on the left side, from low to pivot - 1
2.3 - call recursion on the right side, from pivot + 1, to high

3.1 - assign the pivot to the last position (or random if we want the optimized version)
3.2 - create an i variable and assign it to the lowest position (low)
3.3 - start a for loop with another variable j at low, and while its lower than the high position
	if the content arr[j] is smaller than the pivot value, swap positions and increment the i variable
3.4 - after ending the cycle, swap i position and high and return the index i which is the pivor

4.1 - assign arr[i] to a temp variable
4.2 - assign arr[i] to arr[j]
4.3 - assign arr[j] to temp
