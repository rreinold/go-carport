package main

import (
    "fmt"
    "github.com/tarm/goserial"
    "time"
)

func main() {
    c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 38400}
    s, err := serial.OpenPort(c)

    if err != nil {
            fmt.Println(err)
    }

    _, err = s.Write([]byte("01 01"))

    if err != nil {
            fmt.Println(err)
    }

    time.Sleep(time.Second/2)

    buf := make([]byte, 40)
    for true {

	    n, err := s.Read(buf)

	    if err != nil {
	            fmt.Println(err)
	    }
    	fmt.Println(string(buf[:n]))
    }


    s.Close()
}