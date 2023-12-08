package queue

type Queue interface {
	Put(interface{}) bool
	Poll() (interface{}, bool)
}
