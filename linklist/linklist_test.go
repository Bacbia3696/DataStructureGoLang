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

func TestDeleteFirst(t *testing.T) {
	l := LinkedList{}
	l.InsertBegin(0)
	Convey("Remove 1 element linklist should have length 0", t, func() {
		i, err := l.DeleteFirst()
		So(err, ShouldEqual, nil)
		v := i.(int)
		So(v, ShouldEqual, 1)
	})
}

func TestDeleteLast(t *testing.T) {
	l := LinkedList{}
	l.InsertBegin(0)
	l.InsertBegin(1)
	Convey("Remove last element of 1, 0 should return 0", t, func() {
		l.Display()
		i, err := l.DeleteLast()
		So(err, ShouldEqual, nil)
		v := i.(int)
		So(v, ShouldEqual, 0)
		l.Display()
	})
}

func TestDelete(t *testing.T) {
	l := LinkedList{}
	l.InsertBegin(0)
	l.InsertBegin(1)
	l.InsertBegin(2)
	Convey("Remove middle element", t, func() {
		l.Display()
		i, err := l.Delete(1)
		So(err, ShouldEqual, nil)
		v := i.(int)
		So(v, ShouldEqual, 1)
		l.Display()
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
	l.Insert(-4, 10000)
	l.Display()
}
