package unittest

import (
	"fmt"
	"reflect"
	"runtime"
)

func generateCaseV1(s *SpecV1, opt *Option) (*GoDynamicTestData, error) {

	ctx, err := newContext(opt)
	if err != nil {
		return nil, err
	}

	for _, parserFunc := range []func(ctx *caseContext, s *SpecV1) (err error){
		interpretPackageDef,
		interpretSelectors,
		interpretTestCases,
		deleteAbstractTestCases,
		assignDefaultValuesToTestCases,
	} {
		err = parserFunc(ctx, s)
		if err != nil {

			f := runtime.FuncForPC(reflect.ValueOf(parserFunc).UnsafeAddr())
			if f != nil {
				return nil, fmt.Errorf("%v: %v", f.Name(), err)
			} else {
				return nil, err
			}
		}
	}
	//s.Selector
	return ctx.Gd, nil
}
