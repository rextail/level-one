package t21

type Headphones interface {
	ConnectWithMiniJack() string
}

type Amplifier interface {
	ConnectWithLargeJack() string
}

type SimpleHeadphones struct{}

func (h *SimpleHeadphones) ConnectWithMiniJack() string {
	return "Подключены через 3.5 мм разъем"
}

type SimpleAmplifier struct{}

func (a *SimpleAmplifier) ConnectWithLargeJack() string {
	return "Подключены через 6.3 мм разъем"
}

// Адаптер для наушников

type HeadphonesAdapter struct {
	headphones Headphones
}

func (adapter *HeadphonesAdapter) ConnectWithLargeJack() string {
	// Адаптер с 3.5 мм на 6.3 мм
	return adapter.headphones.ConnectWithMiniJack() + " через адаптер к 6.3 мм разъему"
}
