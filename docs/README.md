# Answers

## Task 1

### Buggy program 1

The program doesn't work because it arrives at a deadlock when it tries to send data in a channel without having any alive goroutines that can receive it. This is fixed by creating a goroutine and sending the data inside it and then receiving it outside.

### Buggy program 2

The issue with this program is that when the main function sends the last number, 11, it instantly closes down the program while the Print goroutine has to wait 10 ms before printing. This can be fixed by adding incrementing a WaitGroup before running Print and then waiting until it gets marked as done, decreased the WaitGroup, when the function has ended.

## Task 2

* What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function?

It closes the channel before all producers are able to send their strings and it the program then panics when they send on the closed channel.

* What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?

It closes the channel as soon as only one of the producers are finished and it then panics when the other producers try to send on the closed channel.

* What happens if you remove the statement `close(ch)` completely?

Nothing visible happens as the program can still close even though the consumers don't exit their for-loops as they don't receive a signal that the channel is closed. However, it is required when adding a WaitGroup for the consumers as they have to exit their for-loops to be marked as done.

* What happens if you increase the number of consumers from 2 to 4?

The program finishes faster as the more number of consumers are able to print the produced messages faster by being parallell.

* Can you be sure that all strings are printed before the program stops?

This is most likely the case because the program waits for all the producers to finish and they block until their strings are received by the consumers where they then are printed. However, it is not guaranteed that the strings actually have time to be printed as that specific task is not waited for.