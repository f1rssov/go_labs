package memo

type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type request struct {
	key      string
	response chan<- result
	done     <-chan struct{}
}

type Memo struct {
	requests chan request
}

var canceledKeys chan string

func New(f Func) *Memo {
	memo := &Memo{
		requests: make(chan request),
	}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response

	select {
	case <-done:
		canceledKeys <- key
	default:
	}
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for {
	CleanCancels:
		for {
			select {
			case key := <-canceledKeys:
				delete(cache, key)
			default:
				break CleanCancels
			}
		}

		select {
		case req := <-memo.requests:
			e := cache[req.key]
			if e == nil {

				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done) // call f(key)
			}
			go e.deliver(req.response)
		default:
		}
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	e.res.value, e.res.err = f(key, done)

	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
