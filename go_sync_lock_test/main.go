package main

import "sync"



func counter_without_lock(){
	var wg sync.WaitGroup
	var counter int
	counter = 0
	for i:=0;i<1000;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	println(counter)
}

func counter_with_lock(){
	var wg sync.WaitGroup
	var mx sync.Mutex
	var counter int
	counter = 0
	for i:=0;i<1000;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			mx.Lock()
			counter++
			mx.Unlock()
		}()
	}
	wg.Wait()
	println(counter)
}

func main(){
	counter_without_lock()
	counter_with_lock()
}

