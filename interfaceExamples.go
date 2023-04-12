package main

//Until recently, the Go spec said that an interface defines a method set,
//which is roughly the set of methods enumerated in the interface.
//Any type that implements all those methods implements that interface.

type etlMethods interface {
	start() error
	run(string) error
	shutDown()
}

type etlOne struct {
	name       string
	configPath string
}

func (eo *etlOne) start() error {
	return nil
}

func (eo *etlOne) run(env string) error {
	return nil
}

func (eo *etlOne) shutDown() {
}

type etlTwo struct {
	name       string
	configPath string
}

func (et *etlTwo) start() error {
	return nil
}

func (et *etlTwo) run(env string) error {
	return nil
}

func (et *etlTwo) shutDown() {
}

func alekMain() {
	//etl1 := new(etlOne)
	//etl2 := new(etlTwo)

	//err1 := etl1.start()
	//err2 := etl2.start()
}

//But another way of looking at this is to say that the interface defines a set of types,
//namely the types that implement those methods. From this perspective,
//any type that is an element of the interface’s type set implements the interface.

//THIS IS an interface used as type constraint
type etls interface {
	etlOne | etlTwo // << THIS CALLED 'TYPE SET'
}

// Another way of saying this is that Methods interface is satisfied by only etls type set
