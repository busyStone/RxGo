package RxGo

import (
  "log"
  "testing"
)

type testObserver struct{}

func (t testObserver) OnNext(i interface{}) {
  log.Println(i)
}

func (t testObserver) OnComplete() {
  log.Println("complete.")
}

func (t testObserver) OnError(err error) {
  log.Println(err)
}

func TestObservableCreate(t *testing.T) {

  fo := func(observer Observer) {
    observer.OnNext(10)
    observer.OnNext(20)

    observer.OnComplete()
  }

  observable := Create(fo)
  observable.Subscribe(testObserver{})
}
