package unittest

func mapString(f func(string) string, x []string) []string {
	for i := range x {
		x[i] = f(x[i])
	}
	return x
}
