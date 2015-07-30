package RxGo

// Observer define an observer should implement methods.
type Observer interface {
  OnNext(interface{})
  OnComplete()
  OnError(err error)
}

// Observable xx
type Observable interface {
  Subscribe(observer Observer)
}
