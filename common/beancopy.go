package common

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func DeepCopyByJson(dst, src interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &dst); err != nil {
		return err
	}
	return nil
}
