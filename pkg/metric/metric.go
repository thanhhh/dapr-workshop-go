package metric

type Metrics interface {
	IncHits(status int, method, path string)
	ObserverResponseTime(status int, method, path string, observerTime float64)
}

// TODO: Implemetation it later
