package cat6

import (
	"fmt"
	"time"

	intcode "github.com/chr-ras/advent-of-code-2019/util/intcode2"
)

// RunNetwork starts a network of 50 intcode computers communicating via queues.
func RunNetwork(nicProgram []int64) {
	network := buildNetwork()
	natQueue := make(chan networkPackage, 20000)
	networkIdle := make(chan struct{})

	for networkAddress := int64(0); networkAddress < 50; networkAddress++ {
		computer := network[networkAddress]

		go intcode.ExecuteProgram(fmt.Sprintf("%03d", networkAddress), nicProgram, computer.finalMemory, computer.inputQueue, computer.outputQueue, 4096)
	}

	go runNAT(natQueue, networkIdle, network)
	handleNetworkTraffic(network, natQueue, networkIdle)
}

func buildNetwork() map[int64]computer {
	network := make(map[int64]computer)

	for networkAddress := int64(0); networkAddress < 50; networkAddress++ {
		finalMemory := make(chan []int64)
		inputQueue := make(chan int64, 20000)
		outputQueue := make(chan int64, 20000)

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

func handleNetworkTraffic(network map[int64]computer, natQueue chan networkPackage, networkIdle chan struct{}) {
	for {
		// The network works with time slices, as otherwise the network would be considered idle as the qeues may be empty
		// while the computers do some work before pushing the next output into a queue.
		time.Sleep(10 * time.Millisecond)

		netWorkIsIdle := true

		for networkAddress := int64(0); networkAddress < 50; networkAddress++ {

			select {
			case targetAddress := <-network[networkAddress].outputQueue:
				netWorkIsIdle = false

				x := <-network[networkAddress].outputQueue
				y := <-network[networkAddress].outputQueue

				if targetAddress == 255 {
					natQueue <- networkPackage{x: x, y: y}
					fmt.Printf("Computer %d sends x = %d, y = %d to NAT\n", networkAddress, x, y)
					continue
				}

				network[targetAddress].inputQueue <- x
				network[targetAddress].inputQueue <- y

				fmt.Printf("Computer %d sends x = %d, y = %d to computer %d\n", networkAddress, x, y, targetAddress)
			default:
				continue
			}
		}

		if netWorkIsIdle {
			networkIdle <- struct{}{}
		}
	}
}

func runNAT(natInputQueue chan networkPackage, networkIdle chan struct{}, network map[int64]computer) {
	currentPackage := networkPackage{x: -1, y: -1}
	lastSentY := int64(-1)

	for {
		select {
		case currentPackage = <-natInputQueue:
			fmt.Printf("NAT received package x = %d, y = %d\n", currentPackage.x, currentPackage.y)
		default:
		}

		select {
		case <-networkIdle:
			if lastSentY != -1 && lastSentY == currentPackage.y {
				fmt.Printf("Identical NAT package Y sent to computer 0 with Y = %d\n", currentPackage.y)
			}

			enqueue(network[0].inputQueue, currentPackage.x)
			enqueue(network[0].inputQueue, currentPackage.y)

			lastSentY = currentPackage.y
		default:
		}
	}
}

func enqueue(queue chan int64, value int64) {
	select {
	case queue <- value:
		return
	default:
		panic(fmt.Errorf("enqueue failed"))
	}
}

type computer struct {
	inputQueue, outputQueue chan int64
	finalMemory             chan []int64
}

type networkPackage struct {
	x, y int64
}
