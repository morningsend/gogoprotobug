package pb

import proto "github.com/gogo/protobuf/proto"

type MyCustomField struct {
	data []byte
}

var _ proto.Marshaler = (*MyCustomField)(nil)
var _ proto.Unmarshaler = (*MyCustomField)(nil)

var data = []byte("hello")

func NewMyCustomField(data []byte) *MyCustomField {
	return &MyCustomField{
		data: data,
	}
}

func (f MyCustomField) Size() int {
	return len(f.data)
}

func (f MyCustomField) Marshal() ([]byte, error) {
	return f.data, nil
}

func (f MyCustomField) Unmarshal(data []byte) error {
	data_copy := make([]byte, len(data))
	copy(data_copy, data)
	f.data = data_copy
	return nil
}

func (f *MyCustomField) MarshalTo(dAtA []byte) (int, error) {
	return copy(dAtA, f.data), nil
}
