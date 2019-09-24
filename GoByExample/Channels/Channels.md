Channels are the pipes that connect concurrent goroutines. 
You can send values into channels from one goroutine and receive those values into another goroutine.

make(chan val-type).

### Channel Buffering
By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

### Cahnnel Directions
When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.

### Select
Go’s select lets you wait on multiple channel operations.