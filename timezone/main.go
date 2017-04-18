package main
import "time"
import "fmt"
func main() {
	location, _ := time.LoadLocation("Europe/Kiev")
	fmt.Println("Local time in Kiev: ", time.Now().In(location))

	location, _ = time.LoadLocation("Europe/Budapest")
	fmt.Println("Local time in Budapest: ", time.Now().In(location))

	location, _ = time.LoadLocation("America/Los_Angeles")
	fmt.Println("Local time in LosAngeles: ", time.Now().In(location))

	fmt.Println("GMT+2:", time.Now().UTC().Add(2*time.Hour))
}
