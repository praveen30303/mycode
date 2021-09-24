package main

import (
	"os"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

type bankOp struct{
	howMuch int
	confirm chan int
}

var accountBalance = 0
var bankRequests chan *bankOp

func increaseBalance(amt int) int {
   update := &bankOp{howMuch: amt, confirm: make(chan int)}
   bankRequests <- update
   newBalance := <-update.confirm
   return newBalance
}
func logBalance(current int) { }

func reportAndExit(msg string) {
   fmt.Println(msg)
   os.Exit(-1) // all 1s in binary
}

func main(){
	if len(os.Args) < 2 {
      reportAndExit("\nUsage: go run channels03.go <number of updates per thread>")
   }
   iterations, err := strconv.Atoi(os.Args[1])
   if err != nil {
      reportAndExit("Bad command-line argument: " + os.Args[1]);
   }
   bankRequests = make(chan *bankOp,8)

   var wg sync.WaitGroup

   go func(){
	   for {
		   select {
		   case request := <-bankRequests:
			   accountBalance += request.howMuch
			   request.confirm <- accountBalance
		   }
	   }
   }()

   wg.Add(1)           // increment WaitGroup counter
   go func() {
      defer wg.Done()  // invoke Done on the WaitGroup when finished
      for i := 0; i < iterations ; i++ {
         newBalance := increaseBalance(1)
         fmt.Println("Adding")
         logBalance(newBalance)
         runtime.Gosched()  // yield to another goroutine
      }
   }()

   // spendy decrements the balance
   wg.Add(1)           // increment WaitGroup counter
   go func() {
      defer wg.Done()
      for i := 0; i < iterations; i++ {
         newBalance := increaseBalance(-1)
         fmt.Println("Removing")
         logBalance(newBalance)
         runtime.Gosched()  // be nice--yield
      }
   }()

   wg.Wait()  // await completion of profiting and spendy
   fmt.Println("Final balance: ", accountBalance) // confirm the balance is zero
   }
   
