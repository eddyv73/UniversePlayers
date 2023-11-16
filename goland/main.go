package main

// import control "universe-players/controllers"
import (
	"fmt"
	connector "universe-players/Connectors"
	control "universe-players/controllers"
)

func main() {
	control.Storm("https://ifconfig.me")
	res := connector.GetIp("https://ifconfig.me")
	fmt.Println("Result : ", res)

}
