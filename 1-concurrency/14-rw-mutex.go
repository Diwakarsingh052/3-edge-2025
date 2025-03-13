package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Theater represents a theater with a specific number of seats
type Theater struct {
	seats   int          // Total seats available in the theater
	invoice chan string  // Channel to store the name of the person whose seat is booked
	rw      sync.RWMutex // Read-write Mutex for synchronizing shared access to Seats
}

var wgBook = new(sync.WaitGroup)

func (t *Theater) bookSeat(name string) {
	defer wgBook.Done()
	// when Write lock is acquired, no other read or writes are allowed
	t.rw.Lock()
	// Releases the write lock when func completes
	defer t.rw.Unlock()

	if t.seats > 0 {
		// Simulate a seat booking-making process
		fmt.Println("Seat is available for", name)
		//time.Sleep(2 * time.Second)
		fmt.Println("Booking confirmed", name)

		t.seats-- // Decrement available seats
		t.invoice <- name
	} else {
		fmt.Println("No seats available for", name) // Inform that no seats are available
	}
}

func (t *Theater) printInvoice(wg *sync.WaitGroup) {
	defer wg.Done()
	for name := range t.invoice {
		fmt.Printf("Invoice is sent to %s\n", name)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	t := Theater{
		seats:   2,
		invoice: make(chan string),
	}
	for i := 0; i < 5; i++ {
		wgBook.Add(1)
		go t.bookSeat("User " + strconv.Itoa(i))
	}
	wg.Add(1)
	go t.printInvoice(wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		//waiting for all the bookings to finish, so we can close the channel
		// to stop the for range which is receiving values from t.invoice
		wgBook.Wait()
		close(t.invoice)
	}()

	wg.Wait()
}
