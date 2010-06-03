package main

import "./wanderer"
func main() {
    //Listen on port
    //Whatever
    server := new(WandererServer)
    server.Init()
    server.Main()
}
