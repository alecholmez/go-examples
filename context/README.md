# Deadlines
A basic demonstration of using the context package for setting deadlines across subprocesses

## Running
```bash
dep ensure -v
go run main.go
```

## Editing
By default, we expect to hit the pre-defined context deadline. If you wish to change the timing, edit the `time.Sleep()` on line `69` to something below 3 seconds. This will let the catch condition (5 pings to google) to succeed and complete the computation defined in the `DoSomething()` function before the context sends the `ctx.Done()` signal.

## Questions
If you have any questions, feel free to create an issue with the tag `question` and I will try to respond as quickly as possible.