package heyliar

import (
	"testing"
)

func TestTotalparsing(t *testing.T) {

	println("Test Start!")

	println(Totalparsing("{3+2>5}or{33-2==31}"))              // true
	println(Totalparsing("{33<22}or{(32+22*2-3*100)+2*6<0}")) // true
	println(Totalparsing("({33<22}or{32+22>0})and{123>0}"))   // true

	println(Calcliar("3*5==15"))                           // true
	println(Calcliar("3*5!=15"))                           // false
	println(Calcliar("(3+2)*3+3+(-123)*2/(2+(-1)*7)<200")) // true

	println(Calculate("1+2"))             // 3
	println(Calculate("6*3"))             // 18
	println(Calculate("(3*2)+3-2*(2/2)")) // 7

}
