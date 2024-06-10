package types

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Payload interface {
	io.WriterTo
	io.ReaderFrom
	Bytes() []byte
}

type Binary []byte

func (m Binary) Bytes() []byte {
	return m
}

func (m Binary) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, uint32(len(m)))

	if err != nil {
		fmt.Println("write to error", err)
		return 0, err
	}

	n, err := w.Write(m)

	// ini karena ada overhead jadi +5 bytes
	return int64(n + 4), err
}

func (m *Binary) ReadFrom(r io.Reader) (int64, error) {

	var size int32

	err := binary.Read(r, binary.BigEndian, &size)

	if err != nil {
		fmt.Println("read from error", err)
		return 0, err
	}

	*m = make([]byte, size)

	s, err := r.Read(*m)

	return int64(s + 4), err
}

func Decode(r io.Reader) (Payload, error) {

	payload := new(Binary)

	_, err := payload.ReadFrom(r)

	if err != nil {
		fmt.Println("decode error", err)
		return nil, err
	}

	return payload, nil
}
