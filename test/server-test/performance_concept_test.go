package server_test

//func createSet() int {
//	var a = map[string]interface{}{}
//	a["some_value"] = 1
//	return a["some_value"].(int)
//}
//
//func BenchmarkMappingContext(c *testing.B) {
//	for i := 0; i < c.N; i++ {
//		_ = createSet()
//	}
//}
//
//func createI() map[string]interface{} {
//	var a = map[string]interface{}{}
//	a["some_value"] = 1
//	return a
//}
//
//func getII(a map[string]interface{}) int {
//	return a["some_value"].(int)
//}
//
//func BenchmarkMappingContextII(c *testing.B) {
//	for i := 0; i < c.N; i++ {
//		a := createI()
//		_ = getII(a)
//	}
//}
//
//func createN(c int) map[string]interface{} {
//	var a = map[string]interface{}{}
//	for i := 0; i < c; i++ {
//		a[string([]byte{uint8(i)})] = i
//	}
//	return a
//}
//
//func getN(a map[string]interface{}, c int) (res int) {
//	for i := 0; i < c; i++ {
//		res += a[string([]byte{uint8(i)})].(int)
//	}
//	return
//}
//
//func BenchmarkMappingContextN(c *testing.B) {
//	for count := 7; count < 16; count++ {
//		c.Run("contextN", func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				a := createN(count)
//				_ = getN(a, count)
//			}
//		})
//	}
//}
//
//const (
//	APIShared = iota
//	APIShared2
//	APISharedLength
//)
//
//func createPreN(n int) (q []interface{}) {
//	q = make([]interface{}, n)
//	for i := 0; i < n; i++ {
//		q[i] = i
//	}
//	return
//}
//
//func getPreN(a []interface{}, n int) (res int) {
//	for i := 0; i < n; i++ {
//		res += a[i].(int)
//	}
//	return
//}
//
//func BenchmarkPreIndexedContextN(c *testing.B) {
//	for count := 7; count < 16; count++ {
//		c.Run("contextN", func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				a := createPreN(count)
//				_ = getPreN(a, count)
//			}
//		})
//	}
//}
//
//func get3N(a map[string]interface{}, c int) (res int) {
//	res += getN(a, c)
//	res += getN(a, c)
//	res += getN(a, c)
//	return
//}
//
//func BenchmarkMappingContext3N(c *testing.B) {
//	for count := 7; count < 16; count++ {
//		c.Run("contextN", func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				a := createN(count)
//				_ = get3N(a, count)
//			}
//		})
//	}
//}
//
//type sharedContext struct {
//	someValue int
//}
//
//func newSharedContext() *sharedContext {
//	return new(sharedContext)
//}
//
//func getIII(a *sharedContext) int {
//	return a.someValue
//}
//
//func BenchmarkMappingContextIII(c *testing.B) {
//	for i := 0; i < c.N; i++ {
//		var a = newSharedContext()
//		a.someValue = 1
//		_ = getIII(a)
//	}
//}
