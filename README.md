# Go Data Structures

[![GoDoc](https://godoc.org/github.com/kevinpollet/go-datastructures?status.svg)](https://godoc.org/github.com/kevinpollet/go-datastructures) [![Go Report Card](https://goreportcard.com/badge/github.com/kevinpollet/go-datastructures)](https://goreportcard.com/report/github.com/kevinpollet/go-datastructures)

Project to keep track of my learnings and experiments while learning the Go programming language and reviewing the common Abstract Data Types (ADT), data structures, and algorithms.

## Installation

```shell
$ go get github.com/kevinpollet/go-datastructures
```

## Abstract Data Types & Data Structures

### List

```go
type List interface {
	Add(value interface{})
	Clear()
	Get(index int) (interface{}, error)
	IndexOf(value interface{}) int
	Insert(index int, value interface{}) error
	IsEmpty() bool
	Remove(index int) (interface{}, error)
	Size() int
}
```

### Stack

```go
type Stack interface {
	Clear()
	IsEmpty() bool
	Peek() (interface{}, error)
	Pop() (interface{}, error)
	Push(value interface{})
	Size() int
}
```

### Queue

```go
type Queue interface {
	Clear()
	IsEmpty() bool
	Offer(value interface{})
	Peek() (interface{}, error)
	Poll() (interface{}, error)
	Size() int
}
```

### Dequeue / Double-ended Queue

```go
type Dequeue interface {
	Clear()
	IsEmpty() bool
	OfferFirst(value interface{})
	OfferLast(value interface{})
	PeekFirst() (interface{}, error)
	PeekLast() (interface{}, error)
	PollFirst() (interface{}, error)
	PollLast() (interface{}, error)
	Size() int
}
```

### PriorityQueue

```go
type PriorityQueue interface {
	Clear()
	IsEmpty() bool
	Offer(value interface{}, priority int)
	Peek() (interface{}, error)
	Poll() (interface{}, error)
	Size() int
}
```

## Readings & Lectures

- [Abstract Data Types](https://brilliant.org/wiki/abstract-data-types/)
- [Easy to Advanced Data Structures](https://www.udemy.com/introduction-to-data-structures/)

## License

[MIT](./LICENSE)
