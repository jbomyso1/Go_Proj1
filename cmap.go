package main

type ChannelMap struct {
    askChan chan int
    addChan chan int
    reduceChan chan int
    store map[string]int
}

func NewChannelMap() *ChannelMap{
    return &ChannelMap()
}

func (cmap * ChannelMap) Listen() {

}
func (cmap * ChannelMap) Stop() {

}

func (cmap * ChannelMap) Reduce(functor ReduceFunc, accum_str string, accum_int int) (string, int) {

}

func (cmap * ChannelMap) AddWord(word string) {

}

func (cmap * ChannelMap) GetCount(word string) int {

}

