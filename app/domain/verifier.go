package domain

type verifier func() bool

func compose(fs ...verifier) verifier {
	return func() bool {
		for _, f := range fs {
			if !f() {
				return false
			}
		}

		return true
	}
}