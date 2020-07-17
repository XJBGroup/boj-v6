package unittest

import "strings"

func deleteAbstractTestCases(ctx *caseContext, _ *SpecV1) (err error) {

	for i := len(ctx.Gd.TestCases) - 1; i >= 0; i-- {
		if ctx.Gd.TestCases[i].Abstract {
			for j := i + 1; j < len(ctx.Gd.TestCases); j++ {
				ctx.Gd.TestCases[j-1] = ctx.Gd.TestCases[j]
			}
			ctx.Gd.TestCases = ctx.Gd.TestCases[:len(ctx.Gd.TestCases)-1]
		}
	}
	return nil
}

func assignDefaultValuesToTestCases(ctx *caseContext, s *SpecV1) (err error) {

	for i := range s.Default {
		for k, v := range s.Default[i] {
			var matcher Matcher = TrueMatcher{}
			if len(k) != 0 && k[0] == '$' {
				// todo: selector
				li := strings.LastIndex(k[1:], ").")
				if li == -1 {
					panic("wrong selector dot")
				}
				li += 2
				//k, p = xs[0], xs[1]
				matcher = newTestCaseMatcher(ctx, k[1:li])
				k = k[li+1:]
			}

			k, v, err = ctx.parseMetaKV(k, v)
			if err != nil {
				return err
			}

			for _, t := range ctx.Gd.TestCases {
				if ok, err := matcher.Match(t); ok && err == nil {
					err = ctx.inheritPropertyKV(k, v, t.Meta)
					if err != nil {

						return err
					}
				} else if err != nil {

					return err
				}
			}
		}
	}
	return nil
}
