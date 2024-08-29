package t14

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func DetectType(val interface{}) string {
	switch val.(type) {
	case int:
		return "int"
	case bool:
		return "bool"
	case string:
		return "string"
	case chan struct{}:
		return "chan struct{}"
	default:
		return "unsupported"
	}
}
