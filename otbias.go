// OTBIAS -- calculate anode current and dissipation from centre tap and
//           anode voltages and output transformer resistance
//
// usage: otbias Vcentretap Vanode Rtransformer [Panode]
//
// svm 18-JUN-2023
//
// I = V/R         R = V/I         V = R*I
// P = I*V         I = P/V         V = P/I
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 4 || len(os.Args) > 5 {
		println("bad arg")
		os.Exit(1)
	}

	vcentre, _ := strconv.ParseFloat(os.Args[1], 64)
	vanode, _ := strconv.ParseFloat(os.Args[2], 64)
	rtrans, _ := strconv.ParseFloat(os.Args[3], 64)

	var pdesired float64
	if len(os.Args) == 5 {
		pdesired, _ = strconv.ParseFloat(os.Args[4], 64)
	}

	vdrop := vcentre - vanode
	ianode := vdrop / rtrans
	panode := vanode * ianode

	fmt.Printf("centre tap voltage            = %1.4g V\n", vcentre)
	fmt.Printf("anode voltage                 = %1.4g V\n", vanode)
	fmt.Printf("output transformer resistance = %1.4g Ω\n", rtrans)
	if pdesired > 0 {
		fmt.Printf("desired anode dissipation     = %1.4g W\n", pdesired)
	}
	fmt.Println()
	fmt.Printf("voltage drop                  = %1.4g V\n", vdrop)
	fmt.Printf("anode current                 = %1.4g mA\n", 1000*ianode)
	fmt.Printf("anode dissipation             = %1.4g W\n", panode)
	if pdesired > 0 {
		idesired := pdesired / vanode
		vdrop = idesired * rtrans
		fmt.Println()
		fmt.Printf("desired voltage drop          = %1.4g V\n", vdrop)
	}
}
