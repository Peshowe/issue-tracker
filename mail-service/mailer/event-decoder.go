package mailer

type EventDecoder interface {
	//DecodeEvent unmarshalls the given byte array data into event
	DecodeEvent(data []byte, event interface{}) error
}
