package main

var user UserID

func init() {
	// TODO get username
	// TODO
}


func main() {
	// TODO set up everything (in INIT?)
	// TODO 

	c := make(chan Message)
	u := make(chan string)
	r := make(chan Message)
	q := make(chan bool)

	go func() {
		s, ok := GetUserInput()
		if !ok || s == "q" || s == "quit" {
			q <- true // Quit by sending true to quit channel
			return
		}
		c <- CreateMessage(s, user)
	}()

	// TODO start networking goroutines
	// TODO forwarding
	// TODO receiving


	<-q // Wait until we're done
}
