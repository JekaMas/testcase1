package inject

import "fmt"

const zero = '0'

//toUint32 straight forward byte slice convert to uint32
func toUint32(b []byte) (uint32, error) {
	var (
		res     uint32
		current int64
	)

	for _, c := range b {
		current = int64(c - zero)
		if current < 0 || current > 9 {
			return 0, fmt.Errorf("incorrect int")
		}

		res = res*10 + uint32(current)
	}
	return res, nil
}

//toUint64 straight forward byte slice convert to uint64
func toUint64(b []byte) (uint64, error) {
	var (
		res     uint64
		current int64
	)

	for _, c := range b {
		current = int64(c - zero)
		if current < 0 || current > 9 {
			return 0, fmt.Errorf("incorrect int")
		}

		res = res*10 + uint64(current)
	}
	return res, nil
}
