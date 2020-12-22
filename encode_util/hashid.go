package encode_util

import (
	"errors"
	"fmt"
	"github.com/speps/go-hashids"
	"strings"
)

var GHashIdHelper *HashIdHelper

type HashIdHelper struct {
	hashId *hashids.HashID
}


func NewHashIdHelper(salt string, minLength int) (*HashIdHelper, error)  {
	hd := hashids.NewData()
	hd.MinLength = minLength
	hd.Salt = salt
	h, err := hashids.NewWithData(hd)
	if err != nil{
		return nil, err
	}
	return &HashIdHelper{hashId: h}, nil
}

func (helper *HashIdHelper) EncodeIdInt64(id int64) (string, error){
	content := []int64{id}
	res, err := helper.hashId.EncodeInt64(content)
	if err != nil{
		return "", err
	}
	return res, nil
}

func (helper *HashIdHelper) DecodeIdInt64(eStr string) (int64, error)  {
	result, err := helper.hashId.DecodeInt64WithError(eStr)
	if err != nil{
		return 0, err
	}
	if len(result) < 1{
		return 0, errors.New("error when decoding string")
	}
	return result[0], nil
}

type HashId int64

func (v HashId) value() int64  {
	return int64(v)
}

func (v HashId) MarshalJson() ([]byte, error)  {
	result, err := GHashIdHelper.EncodeIdInt64(v.value())
	if err != nil{
		return nil, err
	}
	content := []byte(fmt.Sprintf(`\"%s\"`, result))
	return content, nil

}

func (v *HashId) UnmarshalJson(bytes []byte) error  {
	content := strings.Trim(string(bytes), `""`)
	result, err := GHashIdHelper.DecodeIdInt64(content)
	if err != nil{
		return err
	}
	*v = HashId(result)
	return nil
}

