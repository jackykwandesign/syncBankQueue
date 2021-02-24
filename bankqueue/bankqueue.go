package bankqueue

const (
	lockStart = iota
	lockEnd
)

type Queue struct {
	wait    int
	run     bool
	txQueue chan int
	msg     chan int
}

var bankQ *Queue

func init() {
	bankQ = new()
	go bankQ.work()
}

func GetInstance() *Queue {
	return bankQ
}
func new() *Queue {
	q := Queue{
		msg:     make(chan int),
		txQueue: make(chan int),
	}
	return &q
}

func (q *Queue) work() {
	for {
		select {
		case n := <-q.msg:
			switch n {
			case lockStart:
				q.wait++
			case lockEnd:
				q.run = false
			}
			if q.wait > 0 && !q.run {
				q.run = true
				q.wait--
				q.txQueue <- 1
			}
		}
	}
}

// func New() *Queue {
// 	q := Queue{
// 		msg: make(chan int),
// 		// txQueue: make(chan int),
// 	}
// 	go func() {
// 		for {
// 			select {
// 			case n := <-q.msg:
// 				switch n {
// 				case lockStart:
// 					q.wait++
// 				case lockEnd:
// 					q.run = false
// 				}
// 				if q.wait > 0 && !q.run {
// 					q.run = true
// 					q.wait--
// 					// q.txQueue <- 1
// 				}
// 			}
// 		}
// 	}()
// 	return &q
// }

func (q *Queue) Start() {
	q.msg <- lockStart
	<-q.txQueue
}

func (q *Queue) End() {
	q.msg <- lockEnd
}
