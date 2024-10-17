package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var (
		services []string       = []string{"vodila", "Automobil", "Pick_Mobile", "Pavel_taxi"}
		wg       sync.WaitGroup = sync.WaitGroup{}
		res_ch                  = make(chan string)
		faster   string
		open     = make(chan struct{})
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, svc := range services {
		//svc := services[i]

		wg.Add(1)
		go func(svc string) {
			picked_ride(ctx, svc, res_ch, open)
			wg.Done()
		}(svc)
	}
	go func() {
		faster = <-res_ch
		close(open)
		cancel()
		close(res_ch)
	}()

	wg.Wait()

	fmt.Println("The car was found in", faster)
}

func picked_ride(ctx context.Context, serviсe_name string, resch chan string, open_ chan struct{}) {
	time.Sleep(3 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Was NOT picked: ", serviсe_name)
			return
		default:
			if rand.Float64() > 0.9999 && open(open_) {
				select {
				case resch <- serviсe_name:
					return
				default:
					return
				}
			}
			continue
		}
	}
}

func open(done chan struct{}) bool {
	select {
	case <-done:
		return false
	default:
		return true
	}
}

/*
	go func() {
 	for {
			select {
			case <-ctx.Done():
				fmt.Println("Контекст закончен")
				return
			default:
				fmt.Println("Working...")
			}
		}
	}()

	fmt.Print(services, len(services))
	time.Sleep(2 * time.Second)
*/
