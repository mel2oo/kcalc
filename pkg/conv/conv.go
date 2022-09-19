package conv

import (
	"errors"
	"strconv"
	"strings"
)

func Int2Bin(i int64, length int) string {
	return align(strconv.FormatInt(i, 2), length)
}

func align(str string, length int) (s string) {
	l1 := length - len(str)
	for l1 > 0 {
		s += "0"
		l1--
	}

	return s + str
}

type Address struct {
	H string
	L string
}

func Addr64ToBin(addr string) (*Address, error) {
	alist := strings.Split(addr, "`")
	if len(alist) != 2 {
		return nil, errors.New("address format error")
	}

	hint, err := strconv.Atoi(alist[0])
	if err != nil {
		return nil, errors.New("address not invalid")
	}

	lint, err := strconv.Atoi(alist[1])
	if err != nil {
		return nil, errors.New("address not invalid")
	}

	return &Address{
		H: Int2Bin(int64(hint), 32),
		L: Int2Bin(int64(lint), 32),
	}, nil
}
