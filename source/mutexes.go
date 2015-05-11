package source
import (
	"fmt"
	"sync/atomic"
	"runtime"
	"time"
	"sync"
	"math/rand"
)

func atomics() {
	var ops uint64 = 0
	for i := 0; i<50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

func mutexs() {

	var status = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64 = 0

	for r := 0; r<100; r++ {
		go func() {
			for {
				key := rand.Intn(5)
				mutex.Lock()
				key=status[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for w := 0; w<10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				status[key]=val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops", opsFinal)

	mutex.Lock()
	fmt.Println("status", status)
	mutex.Unlock()

}

func Mutexes() {
	atomics()
	mutexs()
}

