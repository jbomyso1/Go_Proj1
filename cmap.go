package main

type Reduction struct {
	functor ReduceFunc
	accum_int int
	accum_str string
}

type ReductionAnswer struct {
	word string
	count int
}

type ChannelMap struct {
    askChan chan string
    addChan chan string
    countChan chan int

	killChan chan int
    returnChan chan ReductionAnswer

    reduceChan chan Reduction
    store map[string]int
}

func NewChannelMap() *ChannelMap {
	m := new(ChannelMap)
    m.store = make(map[string]int)
    m.addChan = make(chan string)
    m.askChan = make(chan string)
    m.countChan = make(chan int)
    m.killChan = make(chan int)
    m.reduceChan = make(chan Reduction)
    m.returnChan = make(chan ReductionAnswer)
    return m
}

func NewLockingMap() *ChannelMap {
	m := new(ChannelMap)
    m.store = make(map[string]int)
    return m
}

func (self * ChannelMap) Listen() {
	for {
		select {
			case message := <-self.addChan:
				if val, ok := self.store[message]; ok {
					self.store[message] = val + 1
				} else{
					self.store[message] = 1
				}
			case word := <-self.askChan:
				if val, ok := self.store[word]; ok {
					self.countChan <- val
				} else {
					self.countChan <- 0
				}

			case req := <-self.reduceChan:
				println("received a reduction")
				ret := ReductionAnswer{req.accum_str, req.accum_int}

				for k, v := range self.store {
					ret.word, ret.count = req.functor(ret.word, ret.count, k, v)
				}
				self.returnChan <- ret
			case <- self.killChan:
				return
		}
	}

}
func (self * ChannelMap) Stop() {
	self.killChan <- 1
}

func (self * ChannelMap) Reduce(functor ReduceFunc, accum_str string, accum_int int) (string, int) {
	self.reduceChan <- Reduction{functor, accum_int, accum_str}
	ret := <-self.returnChan
	return ret.word, ret.count
}

func (self * ChannelMap) AddWord(word string) {
	self.addChan <- word
}

func (self *ChannelMap) GetCount(word string) int {
	self.askChan <- word
	return <-self.countChan
}

