package cat6

import (
	"fmt"

	intcode "github.com/chr-ras/advent-of-code-2019/util/intcode2"
)

// RunNetwork starts a network of 50 intcode computers communicating via queues.
func RunNetwork(nicProgram []int64) {
	network := buildNetwork()

	for networkAddress := int64(0); networkAddress < 50; networkAddress++ {
		computer := network[networkAddress]

		go intcode.ExecuteProgram(fmt.Sprintf("%03d", networkAddress), nicProgram, computer.finalMemory, computer.inputQueue, computer.outputQueue, 4096)
	}

	handleNetworkTraffic(network)
}

func buildNetwork() map[int64]computer {
	network := make(map[int64]computer)

	for networkAddress := int64(0); networkAddress < 50; networkAddress++ {
		finalMemory := make(chan []int64)
		inputQueue := make(chan int64, 100)
		outputQueue := make(chan int64, 100)

		network[networkAddress] = computer{
			inputQueue:  inputQueue,
			outputQueue: outputQueue,
			finalMemory: finalMemory,
		}

		inputQueue <- networkAddress
		inputQueue <- -1
	}

	return network
}

func handleNetworkTraffic(network map[int64]computer) {
	for {
		for networkAddress := int64(0); networkAddress < 50; networkAddress++ {

			select {
			case targetAddress := <-network[networkAddress].outputQueue:
				x := <-network[networkAddress].outputQueue
				y := <-network[networkAddress].outputQueue

				network[targetAddress].inputQueue <- x
				network[targetAddress].inputQueue <- y

				fmt.Printf("Computer %d sends x = %d, y = %d to computer %d\n", networkAddress, x, y, targetAddress)
			default:
				continue
			}
		}
	}
}

type computer struct {
	inputQueue, outputQueue chan int64
	finalMemory             chan []int64
}
