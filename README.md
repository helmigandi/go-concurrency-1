# Go Currency

## Mutex

- Locking Mechanism that ensures that only one goroutine can access a specific section of your code.
- Prevent Race Condition among goroutine.
- Imagine these Mutex as somehow a traffic light which in the end allows only one car or 
  one specific goroutine in this case to pass through an intersection or a block of code at the time.
- Mutex will make goroutine FIFO (First In First Out) in queue.
