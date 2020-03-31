package core

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MultipliedBy1e8 uint64

func (m *MultipliedBy1e8) UnmarshalJSON(bytes []byte) (err error) {
	data := strings.Trim(string(bytes), "\"")
	ss := strings.Split(data, ".")
	if len(ss) == 2 {
		var u1, u2 uint64
		if ss[0] != "" {
			u1, err = strconv.ParseUint(string(ss[0]), 10, 64)
			if err != nil {
				return
			}
		}
		if ss[1] != "" {
			u2, err = strconv.ParseUint(string(ss[1]), 10, 64)
			if err != nil {
				return err
			}
		}
		e := uint64(8 - len(ss[1]))
		*m = MultipliedBy1e8(u1*1e8 + u2*iPow(10, e))
		return nil
	}
	if len(ss) == 1 {
		u, err := strconv.ParseUint(string(bytes), 10, 64)
		if err != nil {
			return err
		}
		*m = MultipliedBy1e8(u * 1e8)
		return nil
	}
	return errors.New(fmt.Sprintf("can't parse %s to MultipliedByE8", string(bytes)))
}
