package ucenter

import (
	"context"
	"fmt"

	eid "github.com/guanguoyintao/kuafu/id"
)

var baseSalt = "user_center"

func HashUserID(ctx context.Context) (*eid.HashID, error) {
	salt := fmt.Sprintf("%s_%s", baseSalt, "user")
	h, err := eid.NewHashID(salt, 12)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func EncodeHashUserID(ctx context.Context, uid uint64) (string, error) {
	hid, err := HashUserID(ctx)
	if err != nil {
		return "", err
	}
	hashUid, err := hid.Encode(uid)
	if err != nil {
		return "", err
	}
	return hashUid, nil
}

func DecodeHashUserID(ctx context.Context, hashUid string) (uint64, error) {
	hid, err := HashUserID(ctx)
	if err != nil {
		return 0, err
	}
	uid, err := hid.Decode(hashUid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}
