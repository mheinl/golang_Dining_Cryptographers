# golang_Dining_Cryptographers

Implement the network of dining cryptographers problem with 3 cryptographers in the GO Programming language. You must adhere to the following requirements: 

1. Every cryptographer is its own process.

2. Every coin is its own process

3. The restaurant owner is its own process

4. A table observer (yet another process) that sees the external events.

# Deliverables:

1. Your code

2. Take the events that the table observer sees, and show that they satisfy the anonymity requirement.

3. Allow cryptographer 0 to see all 3 coins. Write a routine for him to determine who pays, and show that your code works by asking one cryptographer to pay.

# Note:

1. http://cpansearch.perl.org/src/SHEVEK/Crypt-Dining-1.01/lib/Crypt/Dining.pm has a Pearl implementation, but it does not answer the turn-in parts of 2 and 3.

2. Dining Philosophers have been implemented in GO at http://f.souza.cc/2011/10/go-solution-for-dining-philosophers.html.

3. The GO programming language has an excellent online tutorial at http://golang.org, and lots of references.
