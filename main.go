package main

import (
    "log"
    "os"
    "os/user"
)

func main() {

    user, err := user.Current()
    if err != nil {
        log.Println("Unable to get user uid", err.Error())
    }

    log.Println("UID of the user, that executed the program is", user.Uid , "and Username is", user.Name)
    log.Println("PID of the current process is", os.Getpid())
}
