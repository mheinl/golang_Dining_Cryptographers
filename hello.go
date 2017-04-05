package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)



/********** Define Objects **********/

// Cryptographer Object
type Cryptographer struct {
	Next *Cryptographer
	secret bool // Result of the flip coin
	paying bool // Is the cryptographer paying?
	pos int // Position at table
	comparison bool // Result of comparison with neighbour
}

// Cryptographers are gonna flipping coins
func (c *Cryptographer) Flip (channel chan bool) {
	c.secret = flip_coin(get_Random())
	fmt.Println("Cryptographer " + strconv.Itoa(c.pos) + "'s coin flip:" + strconv.FormatBool(c.secret))
	channel <- c.secret
}

// Cryptographers are gonna comparing coins and yell the result
func (c *Cryptographer) Compare (channel chan bool) {
	if c.paying {
		c.comparison = Un_xor(c.secret, <-channel)
	} else {
		c.comparison = xor(c.secret, <-channel)
	}
	fmt.Println("Cryptographer " + strconv.Itoa(c.pos) + "'s comparison:" + strconv.FormatBool(c.comparison))
}

// Cryptographers are (sometimes) gonna paying
func (c *Cryptographer) Paying (payer [4]bool) {
	if (payer[c.pos] == true) {
	c.paying = true
	}
}

// Restaurant Owner Object
type Owner struct {
	Next *Owner
	secret float64
	pos string
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

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 2, 64)
}

func Observer (result1 bool, result2 bool, result3 bool) {
	var temp bool
	temp = xor(result1, result2)
	temp = xor(temp, result3)
	fmt.Printf("The Observer says: ")
	if temp {
		fmt.Println("A Cryptographer paid!")
	} else {
		fmt.Println("The NSA paid!")
	}
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
	
	// Assign Payer Role (Couldn't find of a more elegant way, sorry :-D )
	c1.Paying(payer)
	c2.Paying(payer)
	c3.Paying(payer)

	// Flip the Coins in their own Goroutines
	go c1.Flip(coin_channel1)
	go c2.Flip(coin_channel2)
	go c3.Flip(coin_channel3)
	
	// Sleep to avoid race conditions
	time.Sleep(time.Millisecond * 100)
	
	// Compare the neighboured coins and let the Observer observe the Cryptographer's yelling their comparisons
	go c3.Compare(coin_channel2)
	go c2.Compare(coin_channel1)
	go c1.Compare(coin_channel3)
	
	// Sleep to avoid race conditions
	time.Sleep(time.Millisecond * 100)

	// Let the Observer ovserve
	go Observer(c1.comparison, c2.comparison, c3.comparison)
	
	// Sleep to avoid race conditions
	time.Sleep(time.Millisecond * 100)
	
/********** DEBUG **********/
	
	/*
	// check payer array
	for i:=1; i<4; i++{
		fmt.Println(strconv.FormatBool(payer[i]))
	}
	
	// Sleep in order to see result of Goroutines before main exits
	time.Sleep(time.Millisecond * 100)

	fmt.Println(strconv.FormatBool(c1.paying))
	fmt.Println(strconv.FormatBool(c2.paying))
	fmt.Println(strconv.FormatBool(c3.paying))
	

	if flip_coin(get_Random()) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
	*/	

}
