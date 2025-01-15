package license

import (
	"bytes"
	"crypto/ed25519"
	"encoding/json"
	"encoding/pem"
	"errors"
	"os"
	"time"
)

var (
	ErrInvalidSignature = errors.New("许可证签名无效")

	ErrMalformedLicense = errors.New("许可证文件格式不正确")
)

type License struct {
	Iss string          `json:"iss,omitempty"` // 发证单位
	Cus string          `json:"cus,omitempty"` // 客户标识 Id
	Sub string          `json:"sub,omitempty"` // 订阅者 Id
	Typ string          `json:"typ,omitempty"` // 许可证类型
	Lim int             `json:"lim,omitempty"` // 许可证数量限制
	Iat time.Time       `json:"iat,omitempty"` // 签发时间
	Exp time.Time       `json:"exp,omitempty"` // 过期时间
	Dat json.RawMessage `json:"dat,omitempty"` // 附加信息
}

func (l *License) Expired() bool {
	return l.Exp.IsZero() == false && time.Now().Local().After(l.Exp)
}

func (l *License) Encode(privateKey ed25519.PrivateKey) ([]byte, error) {
	msg, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	sig := ed25519.Sign(privateKey, msg)
	buf := new(bytes.Buffer)
	buf.Write(sig)
	buf.Write(msg)

	block := &pem.Block{
		Type:  "LICENSE KEY",
		Bytes: buf.Bytes(),
	}
	return pem.EncodeToMemory(block), nil
}

func Decode(data []byte, publicKey ed25519.PublicKey) (*License, error) {
	block, _ := pem.Decode(data)
	if block == nil || len(block.Bytes) < ed25519.SignatureSize {
		return nil, ErrMalformedLicense
	}

	sig := block.Bytes[:ed25519.SignatureSize]
	msg := block.Bytes[ed25519.SignatureSize:]

	verified := ed25519.Verify(publicKey, msg, sig)
	if !verified {
		return nil, ErrInvalidSignature
	}
	out := new(License)
	err := json.Unmarshal(msg, out)
	return out, err
}

func DecodeFile(path string, publicKey ed25519.PublicKey) (*License, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Decode([]byte(data), publicKey)
}

func DecodeStr(str string, publicKey ed25519.PublicKey) (*License, error) {
	return Decode([]byte(str), publicKey)
}
