package codec

// Codec defines basic marshal/unmarshal behavior.
type Codec interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}
