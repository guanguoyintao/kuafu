package eid

import (
	"errors"

	"github.com/speps/go-hashids/v2"
)

type HashID struct {
	hashID *hashids.HashID
}

// NewHashID 哈希函数的碰撞概率可以用生日问题的原理来估算。假设我们有
// 𝑁
// N 个可能的哈希值，并且我们生成
// 𝑛
// n 个哈希 ID，那么发生碰撞的概率可以近似计算如下：
// p=1-exp(e, -n**2/2N)
func NewHashID(salt string, minLength int) (*HashID, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}
	return &HashID{
		hashID: h,
	}, nil
}

func (h *HashID) Encode(id uint64) (string, error) {
	e, err := h.hashID.EncodeInt64([]int64{int64(id)})
	if err != nil {
		return "", err
	}
	return e, nil
}

func (h *HashID) Decode(hashid string) (uint64, error) {
	d, err := h.hashID.DecodeWithError(hashid)
	if err != nil {
		return 0, err
	}
	if len(d) != 1 {
		return 0, errors.New("hash id decode result should have length 1")
	}
	return uint64(d[0]), nil
}
