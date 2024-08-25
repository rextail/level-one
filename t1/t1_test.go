package t1

import (
	"testing"
)

func TestAction_Work(t *testing.T) {
	h1 := Human{isWBWorker: true}
	h2 := Human{isWBWorker: false}

	a1 := Action{h1}
	a2 := Action{h2}

	if (a1.Work() == "Я работник Wildberries" && a2.Work() == `Я работник ООО "Рога и Копыта"`) == false {
		t.Errorf("a1: %s, a2: %s\n", a1.Work(), a2.Work())
	}
}
