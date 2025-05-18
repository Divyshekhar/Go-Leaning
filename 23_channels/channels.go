package main

import (
	"fmt"
	"time"
)

// sending
func processNum(numChan chan int) {
	fmt.Println("processing number", <-numChan)
}

// recieve
func sum(result chan int, num1 int, num2 int) {
	numresult := num1 + num2
	result <- numresult

}

func task(done chan bool) {
	defer func() { done <- true }()

	fmt.Println("process....")
}

func emailSender(emailChan chan string, done chan bool){
	defer func(){done <- true}()

	for email := range emailChan{
		fmt.Println("sending email to ", email)
		time.Sleep(time.Second)
	}

}

func main() {

	

	//buffered channels (non blocking)

	emailChan := make(chan string, 100) 
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i:= 0; i < 100 ; i++{
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending")

	close(emailChan) //need to close the channel always (buffered)


	// emailChan <- "100xdevs.com"
	// emailChan <- "200xdevs.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)
	<-done
	





	// done := make(chan bool)

	// go task(done)

	// <-done //blocking till sending the data

	// result := make(chan int)
	// go sum(result, 4, 5)

	// res := <-result

	// fmt.Println(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// numChan <- 5

	// time.Sleep(time.Second * 2)

	// messageChan := make(chan string)

	// messageChan <- "ping" //blocking till other side is not ready

	// message := <-messageChan

	// fmt.Println(message)
}
