package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"sync"
)

var wg1 sync.WaitGroup
var wg2 sync.WaitGroup
var wg3 sync.WaitGroup
var wg4 sync.WaitGroup
var wg5 sync.WaitGroup

/********** Define Objects **********/

// Cryptographer Object
type Cryptographer struct {
	secret bool // Result of the flip coin
	paying bool // Is the cryptographer paying?
	pos int // Position at table
	comparison bool // Result of comparison with neighbour
}

// Cryptographers are gonna flipping coins
func (c *Cryptographer) Flip (channel chan bool) {
	c.secret = flip_coin(get_Random())
	fmt.Println("Cryptographer " + strconv.Itoa(c.pos) + "'s coin flip: " + strconv.FormatBool(c.secret))
	channel <- c.secret
	close(channel)
	wg1.Done()
}

// Cryptographers are gonna comparing coins
func (c *Cryptographer) Compare (channel chan bool) {
	if c.paying {
		c.comparison = Un_xor(c.secret, <-channel)
	} else {
		c.comparison = xor(c.secret, <-channel)
	}
	wg2.Done()
}

// Cryptographers are (sometimes) gonna paying
func (c *Cryptographer) Paying (payer [4]bool) {
	if (payer[c.pos] == true) {
	c.paying = true
	}
}

/********** Define Basic Functions **********/

// Get random number
func get_Random() int {
	return rand.Int()
}

// XOR operation on two booleans
func xor(a bool, b bool) bool {
	return a != b
}

// Let certain cryptographer "manually" turn coin if she payed
func Un_xor(a bool, b bool) bool {
	x := xor(a, b)
	return xor(x, true)	//Xor by 1 gives the opposite of the original number
}

// Use randomness to flip a coin, odd for true, even for false
func flip_coin(randomness int) bool {
	// calculate modulo 2, if no rest exists, randomness is even, if rest exists, randomness is odd
	if (randomness % 2 == 0){
		return false
	} else {
		return true
	}
}

// Randomly determine payer (One of the three Cryptographers or NSA)
func determine_Payer(randomness int) [4]bool {
	payer := [4]bool{false, false, false, false}
	payer[randomness % 4] = true
	return payer
}

func Observer (c Cryptographer) {
	fmt.Println("Cryptographer " + strconv.Itoa(c.pos) + "'s comparison: " + strconv.FormatBool(c.comparison))
	wg3.Done()
}

func RestaurantOwner (comparison1 bool, comparison2 bool, comparison3 bool) {
	var temp bool
	temp = xor(comparison1, comparison2)
	temp = xor(temp, comparison3)
	fmt.Printf("\nThe Restaurant Owner concludes:\n")
	if temp {
		fmt.Println("A Cryptographer paid!")
	} else {
		fmt.Println("The NSA paid!")
	}

	wg4.Done()
}

// God-like cryptographer0 aka bruce_schneier
func bruce_schneier (secret1 bool, secret2 bool, secret3 bool, comparison1 bool, comparison2 bool, comparison3 bool) {
	fmt.Printf("\nBruce Schneier aka Cryptographer0 knows:\n")
	if Un_xor(secret1, secret3) == comparison1 {
		fmt.Println("Cryptographer 1 paid!")
	} else if Un_xor(secret2, secret1) == comparison2 {
		fmt.Println("Cryptographer 2 paid!")
	} else if Un_xor(secret3, secret2) == comparison3 {
		fmt.Println("Cryptographer 3 paid!")
	} else {
		fmt.Println("The NSA paid!")
	}

	wg5.Done()
}


/********** Let's do it **********/
func main() {

	// Channel for Coin Communication
	coin_channel1 := make(chan bool, 1)
	coin_channel2 := make(chan bool, 1)
	coin_channel3 := make(chan bool, 1)

 	// Get initial seed to ensure randomness
	rand.Seed(time.Now().UTC().UnixNano())

	// Create objects
	c1 := Cryptographer{
		paying: false,
		pos: 1,
    }

	c2 := Cryptographer{
		paying: false,
		pos: 2,
    }
	c3 := Cryptographer{
		paying: false,
		pos: 3,
    }

	// Determine Payer either randomly or manually (index 0 = NSA)
	var payer [4]bool = determine_Payer(get_Random())
	//payer := [4]bool{true, false, false, false}

	// Assign Payer Role
	c1.Paying(payer)
	c2.Paying(payer)
	c3.Paying(payer)


	// Flip the Coins in their own Goroutines
	
	fmt.Println("\n########################################################")
	fmt.Println("Show coin flipping to check results:")
	fmt.Println("(this is not visible to Cryptographers and Restaurant Owner)")
	wg1.Add(3)
	go c1.Flip(coin_channel1)
	go c2.Flip(coin_channel2)
	go c3.Flip(coin_channel3)
	wg1.Wait()
	fmt.Println("########################################################")

	// Sleep to avoid race conditions
	//time.Sleep(time.Millisecond * 100)

	// Compare the neighboured coins
	wg2.Add(3)
	go c1.Compare(coin_channel3)
	go c2.Compare(coin_channel1)
	go c3.Compare(coin_channel2)
	wg2.Wait()

	// Sleep to avoid race conditions
	//time.Sleep(time.Millisecond * 100)

	// Let the Observer observe
	fmt.Println("\nThe Observer observed:")
	fmt.Println("(difference = true, same results = false; odd number of differences = cryptographer paid, even number = NSA paid)")
	wg3.Add(3)
	go Observer(c1)
	go Observer(c2)
	go Observer(c3)
	wg3.Wait()
	
	// Let the Restaurant Owner conclude
	wg4.Add(1)
	go RestaurantOwner(c1.comparison, c2.comparison, c3.comparison)
	wg4.Wait()

	// Sleep to avoid race conditions
	//time.Sleep(time.Millisecond * 100)

	// Go Bruce Schneier aka Cryptographer0, Go!
	wg5.Add(1)
	go bruce_schneier(c1.secret, c2.secret, c3.secret, c1.comparison, c2.comparison, c3.comparison)
	wg5.Wait()

	// Sleep to avoid race conditions
	//time.Sleep(time.Millisecond * 100)

}
