package main

import "crypto/tls"
import "crypto/x509"
import "encoding/pem"
import "os"
import "fmt"
import "time"
func readEntireFile(filename string) []byte {
    finfo,_ := os.Stat(filename)
    size := int(finfo.Size)
    contents := make([]byte, size)
    fh,_ := os.Open(filename,os.O_RDONLY,0666)
    total := 0
    for {
        n, err := fh.Read(contents[total:size])
        total += n
        if err != nil {
            return contents
        }
    }
    return contents
}

func main () {
    random,_ := os.Open("/dev/urandom", os.O_RDONLY, 0)
    pembytes := readEntireFile("/home/kris/SSL/gr0g.crt")
    cert,_ := pem.Decode(pembytes)
    keybytes := readEntireFile("/home/kris/SSL/gr0g.key")
    pk,_ := pem.Decode(keybytes)
    
    
    privatekey,_ := x509.ParsePKCS1PrivateKey(pk.Bytes)
   
    
    config := new(tls.Config)
    config.Certificates = make([]tls.Certificate, 1)
    config.Certificates[0].Certificate = [][]byte{cert.Bytes}
    config.Certificates[0].PrivateKey = privatekey
    config.Rand = random
    //config.RootCAs = caset  
    config.Time = time.Seconds 
    listener,err:=  tls.Listen("tcp","0.0.0.0:8443",config)
    
    
    
    fmt.Printf("%s\n", err)
    for {
        conn,_ := listener.Accept()
        go func () {
            for {
            buf := make([]byte, 1024)
            _, err := conn.Read(buf)
            if err != nil {return }
            fmt.Printf("%s", buf)
            }
        }()
    }
}
