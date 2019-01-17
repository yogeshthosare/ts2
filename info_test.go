package main

import (
    "testing"
    "os"
    "log"
    "os/user"
    "strconv"
    "regexp"
    "fmt"
    "io/ioutil"
)

func TestInfo(t *testing.T){
	// get username pid from the main script
	uid, pid := GetInfo()

	//get expectedUid and expectedPid from the test script
	user, err := user.Current()
    	if err != nil {
        	log.Println("Unable to get user uid", err.Error())
    	}
	var expectedUid = user.Uid
	var expectedPid = strconv.Itoa(os.Getpid())

	filename := "/proc/"+uid+"/status"
	file, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Print(err)
        }

        filecontent := string(file)

	re := regexp.MustCompile(`1000`)
	matches := re.FindStringSubmatch(filecontent)
	fmt.Println(matches[1])

	
	re = regexp.MustCompile(`17183`)
	matches = re.FindStringSubmatch(filecontent)
	fmt.Println(matches[1])

	if uid != expectedUid {
                t.Fatalf("Expected %s but got %s", uid, expectedUid)
        }
	if pid != expectedPid {
                t.Fatalf("Expected %s but got %s", pid, expectedPid)
        }
}

