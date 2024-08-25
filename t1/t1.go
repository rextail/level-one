package t1

type Action struct {
	//встроили
	Human
}

type Human struct {
	isWBWorker bool
}

func (h Human) Work() string {
	if h.isWBWorker {
		return "Я работник Wildberries"
	} else {
		return `Я работник ООО "Рога и Копыта"`
	}
}
