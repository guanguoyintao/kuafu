package ucenter

import (
	"context"
	"fmt"
	"time"

	jwt5 "github.com/golang-jwt/jwt/v5"
	kxerrors "github.com/guanguoyintao/kuafu/kratos-x/kxerror"
	"github.com/guanguoyintao/kuafu/kratos-x/kxlogging"
	jwt "github.com/guanguoyintao/kuafu/kratos-x/kxmiddleware/auth"
)

// GenerateJwtToken 生成jwt token
func GenerateJwtToken(ctx context.Context, kvs map[string]any, duration time.Duration, secret string) (string, error) {
	claims := (jwt5.MapClaims)(kvs)
	claims["exp"] = time.Now().Add(duration).Unix()
	token, err := jwt5.NewWithClaims(jwt5.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// GetUidFromToken 获取登陆uid
func GetUidFromToken(ctx context.Context) (uint64, error) {
	claims, err := getClaim(ctx)
	if err != nil {
		return 0, err
	}
	if claims == nil {
		return 0, nil
	}
	uid, ok := claims["uid"]
	if !ok {
		return 0, kxerrors.SophonErrorTokenParseNoUID
	}
	hid, err := HashUserID(ctx)
	if err != nil {
		return 0, err
	}
	id, err := hid.Decode(uid.(string))
	if err != nil {
		return 0, kxerrors.SophonErrorUnauthorized
	}
	return id, nil
}

func getClaim(ctx context.Context) (jwt5.MapClaims, error) {
	// todo: jwt 白名单错误优化
	//ok := jwt.GetIsSkipAuthVerification(ctx)
	//if ok {
	//	return nil, nil
	//}
	token, _ := jwt.FromContext(ctx)
	kxlogging.GetGlobalLogger().WithContext(ctx).Info("token is ", token)
	if claims, ok := token.(jwt5.MapClaims); ok {
		return claims, nil
	}
	return nil, kxerrors.SophonErrorTokenClaimTypeInvalid
}

func GenContext(ctx context.Context, uid uint64, secret string) (context.Context, error) {
	hashid, err := HashUserID(ctx)
	if err != nil {
		return ctx, err
	}
	uidHash, err := hashid.Encode(uid)
	if err != nil {
		return nil, err
	}
	kvs := map[string]any{
		"uid": uidHash,
	}
	token, _ := GenerateJwtToken(ctx, kvs, 10*time.Second, secret)
	parsed, err := jwt5.Parse(token, func(token *jwt5.Token) (any, error) {
		if _, ok := token.Method.(*jwt5.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		panic(err)
	}
	if claims, ok := parsed.Claims.(jwt5.MapClaims); ok && parsed.Valid {
		ctx = jwt.NewContext(ctx, claims)
	}

	return ctx, nil
}
