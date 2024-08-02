package utils

func Assert(truthy bool, err error) {
	if !truthy {
		panic(err)
	}
}
