package main

import "sync"


type post struct{
	views int
	mu sync.Mutex
}

func (p *post) incrementViews(wg *sync.WaitGroup){
	defer func ()  {
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()//downside thoda sync type ka kr deta hai modification k around hi mutex lgao
	p.views++
	
}

func main(){
	myPost:=post{views: 0}
	var wg sync.WaitGroup
	
	for i:=0;i<1000;i++{
		wg.Add(1)
		go myPost.incrementViews(&wg)
	}
	wg.Wait()//race condition ko nhi dekh rha hai ye

	println("Views:", myPost.views)
}