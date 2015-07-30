package RxGo

// TFuncSubscribe to wrap a subscribe as observable.
type TFuncSubscribe func(Observer)

// Subscribe implement observable Subscribe method.
func (f TFuncSubscribe) Subscribe(observer Observer) {
  f(observer)
}

// Create to create an observable data flow.
func Create(subscribe func(Observer)) Observable {
  return TFuncSubscribe(subscribe)
}
