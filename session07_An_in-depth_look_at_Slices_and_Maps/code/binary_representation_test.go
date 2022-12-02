package code

import "fmt"

func Example_BinaryRepresentation() {
	var signed int8 = -126
	fmt.Printf("Decimal Value: %+d ; Binary form: %b\n", signed, signed)

	// the Most Significant Bit (MSB) == 1 in value 0b11111110
	// if the number is positive, this has a value of 0, but if the number is negative, the bit is 1.
	var unsigned uint8 = 0b11111110
	fmt.Printf("Decimal Value: %+d ; Binary form: %b\n", unsigned, unsigned)

	// output:
	// Decimal Value: -126 ; Binary form: -1111110
	// Decimal Value: +254 ; Binary form: 11111110
}
