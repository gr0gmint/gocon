package main

import "crypto/tls"
import "bufio"
import "os"
import "fmt"

func main () {
    conn, _ := tls.Dial("tcp", "", "127.0.0.1:8443")
    reader := bufio.NewReader(os.Stdout)
    for {
        line,err := reader.ReadBytes('\n') 
        if err != nil { fmt.Printf("%s\n", err) }
        conn.Write(line)
    }

}
