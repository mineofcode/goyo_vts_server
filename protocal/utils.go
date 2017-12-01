package protocalHandler

import (
	"encoding/binary"
	"fmt"
	"math"
)

//Batryper battry percent
func batryper(btr int) int {
	switch btr {
	case 6:
		return 100
	case 5:
		return 70
	case 4:
		return 50
	case 3:
		return 30
	case 2:
		return 10
	case 1:
		return 5
	case 0:
		return 1
	default:
		return 0
	}
}

//networkper network percentager
func networkper(gsm int) int {
	switch gsm {
	case 4:
		return 100
	case 3:
		return 75
	case 2:
		return 50
	case 1:
		return 25
	case 0:
		return 0
	default:
		return 0
	}
}

/// 016 16 bit, 08 8 bit
func hexToBin(s string, format string) (binString string) {

	binString = fmt.Sprintf(format, binary.BigEndian.Uint16([]byte(s)))
	return binString
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
