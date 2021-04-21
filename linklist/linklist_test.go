package linklist

import (
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	// code to run before each test
	fmt.Println("This run before Each test case")
	exitVal := m.Run()
	// code to run after each test
	os.Exit(exitVal)
}

func TestLength(t *testing.T) {
	l := LinkedList{}
	Convey("Empty list should have length 0", t, func() {
		So(l.Length(), ShouldEqual, 0)
	})
	Convey("1 element list should have length 1", t, func() {
		l.InsertBegin(200)
		l.Display()
		So(l.Length(), ShouldEqual, 1)
	})
	Convey("2 element list should have length 2", t, func() {
		l.InsertEnd(400)
		_ = l.Display()
		So(l.Length(), ShouldEqual, 2)
	})
}

func TestDisplay(t *testing.T) {
	l := LinkedList{}
	Convey("Display empty list should throw ErrListEmpty", t, func() {
		err := l.Display()
		So(err, ShouldEqual, ErrListEmpty)
	})
	l.InsertBegin(1000)
	Convey("Display none empty list should not throw err", t, func() {
		err := l.Display()
		So(err, ShouldEqual, nil)
	})
}

func TestInsert(t *testing.T) {
	l := LinkedList{}
	l.Insert(0, 10)
	l.Display()
	l.Insert(0, 100)
	l.Display()
	l.Insert(1, 1000)
	l.Display()
	l.Insert(3, 10000)
	l.Display()
}
