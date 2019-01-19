package main

import (
    "log"
    "strconv"
    "syscall"
    
)

func GetInfo() (string, string) {
    
    uid   := strconv.Itoa(syscall.Getuid())

    pid   := strconv.Itoa(syscall.Getpid())

    log.Println("UID of the user, that executed the program is", uid)
    log.Println("PID of the current process is", pid)

    return uid, pid
}

func main() {
    GetInfo()
//    var input string
//    fmt.Scanln(&input)
}
