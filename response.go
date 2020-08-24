package gopenapi

type Responses map[string]*ResponseRef

// Response Object.
type Response struct {
	ExtensionProps
	Description *string               `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]*HeaderRef `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     Content               `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*LinkRef   `json:"links,omitempty" yaml:"links,omitempty"`
}

// func (value *ResponseRef) UnmarshalYAML(data []byte) error {
// 	return UnmarshalRef(data, &value.Ref, &value.Value)
// }

// func (response *Response) MarshalJSON() ([]byte, error) {
// 	return jsoninfo.MarshalStrictStruct(response)
// }

// func (response *Response) UnmarshalJSON(data []byte) error {
// 	return UnmarshalStrictStruct(data, response)
// }

// type ObjectEncoder struct {
// 	result map[string]json.RawMessage
// }

// type ObjectDecoder struct {
// 	Data            []byte
// 	remainingFields map[string]json.RawMessage
// }

// type StrictStruct interface {
// 	EncodeWith(encoder *ObjectEncoder, value interface{}) error
// 	DecodeWith(decoder *ObjectDecoder, value interface{}) error
// }

// func UnmarshalStrictStruct(data []byte, value StrictStruct) error {
// 	decoder, err := NewObjectDecoder(data)
// 	if err != nil {
// 		return err
// 	}
// 	return value.DecodeWith(decoder, value)
// }

// func NewObjectDecoder(data []byte) (*ObjectDecoder, error) {
// 	var remainingFields map[string]json.RawMessage
// 	if err := json.Unmarshal(data, &remainingFields); err != nil {
// 		return nil, fmt.Errorf("Failed to unmarshal extension properties: %v\nInput: %s", err, data)
// 	}
// 	return &ObjectDecoder{
// 		Data:            data,
// 		remainingFields: remainingFields,
// 	}, nil
// }
