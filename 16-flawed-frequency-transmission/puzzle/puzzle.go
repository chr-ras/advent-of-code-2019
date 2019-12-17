package main

import (
	"fmt"

	"github.com/chr-ras/advent-of-code-2019/16-flawed-frequency-transmission/fft"
)

func main() {
	inputSignal := "59768092839927758565191298625215106371890118051426250855924764194411528004718709886402903435569627982485301921649240820059827161024631612290005106304724846680415690183371469037418126383450370741078684974598662642956794012825271487329243583117537873565332166744128845006806878717955946534158837370451935919790469815143341599820016469368684893122766857261426799636559525003877090579845725676481276977781270627558901433501565337409716858949203430181103278194428546385063911239478804717744977998841434061688000383456176494210691861957243370245170223862304663932874454624234226361642678259020094801774825694423060700312504286475305674864442250709029812379"

	outputSignal := fft.ApplyFft(inputSignal, 100)

	fmt.Printf("Ouput signal (first 8 digits): %v\n", outputSignal[0:8])
}
