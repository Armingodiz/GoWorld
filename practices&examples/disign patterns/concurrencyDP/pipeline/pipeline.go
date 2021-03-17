package main

// each executor in on the task which pipeline must do it
type Executor func(interface{}) (interface{}, error)

type Pipeline interface {
	Pipe(executor Executor) Pipeline
	Merge() <-chan interface{}
}

type pipeline struct {
	dataC     chan interface{} // input channel which will be used as output channel and then input channel for the next next executor .
	errC      chan error
	executors []Executor
}

func New(f func(chan interface{})) Pipeline {
	inC := make(chan interface{})

	go f(inC) // sending inputs to pipeline

	return &pipeline{
		dataC:     inC,
		errC:      make(chan error),
		executors: []Executor{},
	}
}
// attach new task to pipeline
func (p *pipeline) Pipe(executor Executor) Pipeline {
	p.executors = append(p.executors, executor)

	return p
}
// merge will run all executors and manage their inputs and outputs
func (p *pipeline) Merge() <-chan interface{} {
	for i := 0; i < len(p.executors); i++ {
		p.dataC, p.errC = run(p.dataC, p.executors[i])
	}

	return p.dataC
}

func run(inC <-chan interface{}, f Executor) (chan interface{}, chan error) {
	outC := make(chan interface{})
	errC := make(chan error)

	go func() {
		defer close(outC)
		for v := range inC {
			res, err := f(v)
			if err != nil {
				errC <- err
				continue
			}

			outC <- res
		}
	}()

	return outC, errC
}
