Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Parallell programming aims to execute multiple processes at once using multiple cores or processors. Concurrent programming, on the other hand, runs overlapping processes, and makes it look like simultaneous excecution. 


What is the difference between a *race condition* and a *data race*? 
> Race condition is when a value gets a different outcome depending on timing of which thread finishes first. Data race is when a thread tries to access a mutable that is being written to by another thread. 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler aims for efficient utilazation of the CPU and concurence progrsmming so we dont waste resources. The programs that are ready to be executed are placed in a queue, and the scheduling determines which one of the threads will have access to the CPU for a given time. It can also switch between processes (contex switching) to run a thread with a higher priority first. 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
Using multiple threads makes so that we can run multiple programs that can communicate with each other at once, aims for efficient use of the CPU and runs faster.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> For threads the OS is responsible for managing the threads, while for green threads, the program itself is responsible. 

(https://aleksandr-pezikov.blog/threads-green-threads-fibers-coroutines-what-is-the-difference-cbc2c887c0c1)

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Gives much more flexibility to the developer when coding, but it can also lead to errors when its not properly implemented. 

What do you think is best - *shared variables* or *message passing*?
> shared variable is easier


