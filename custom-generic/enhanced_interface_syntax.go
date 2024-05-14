package main

type L interface {
	Run() error
	Stop()
}

type M interface {
	L
	Step() error
}

// equivalent to
// type M interface {
//	Run() error
//	Stop()
//	Step() error
// }

type N interface {
	M
	interface{ Resume() }
	~map[int]bool    // tilde forms
	~[]byte | string // unions of terms
}

// equivalent to
// type N interface {
//	Run() error
//	Stop()
//	Step() error
//	Resume()
//	~map[int]bool
//	~[]byte | string
// }

type O interface {
	Pause()
	N
	string
	int64 | ~chan int | any
}

// equivalent to
// type O interface {
//	Run() error
//	Stop()
//	Step() error
//	Pause()
//	Resume()
//	~map[int]bool
//	~[]byte | string
//	string
//	int64 | ~chan int | any
// }
