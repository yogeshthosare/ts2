A] Problem statement

Simple Golang program to make a system call that displays:
i. UID of the user, that executed the program.
ii. PID of the current process.

iii. Expectations:
1. Test Case to validate the above.
2. The process cannot be forked, so please do not use ‘whoami’ and
‘hostname’ command.
3. Should be Implemented in Golang.


B] How to run main script - 

	go build -o [main] main.go

or use go run 

	yogesh@yogesh-VirtualBox:~/Go/src/ts2$ go run main.go 
	2019/01/19 12:39:34 UID of the user, that executed the program is 1000
	2019/01/19 12:39:34 PID of the current process is 30773 

C] How to run test script - 
just run following command from the same working directory
	go test -v

	yogesh@yogesh-VirtualBox:~/Go/src/ts2$ go test -v
	=== RUN   TestInfo
	2019/01/19 12:42:41 UID of the user, that executed the program is 1000
	2019/01/19 12:42:41 PID of the current process is 31268
	--- PASS: TestInfo (0.00s)
	PASS
	ok  	ts2	0.008s
