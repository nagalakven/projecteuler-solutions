# Number Spiral Diagonals

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?

# Formula

All n*n grids with odd value of n (3,5,7,9,...) hold similar pattern to get the diagonal numbers
Get the pattern for diagonal calculation

The dialgonals follows this pattern => k^2 + (k^2 - (k-1)) + (k^2 - 2(k-1)) + (k^2 - 3(k-1))
This is simplified to => (4*k^2) - 6(k-1)

# Answer

Answer: 669171001