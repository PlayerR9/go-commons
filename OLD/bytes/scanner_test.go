package bytes

import (
	"io"
	"testing"
)

// ReadRune implements the io.RuneReader interface.
//
// Errors:
//   - io.EOF: When the stream is exhausted.
//   - *ErrInvalidUTF8Encoding: When the stream has an invalid UTF-8 encoding.
//
// Do err == io.EOF to check if the stream is exhausted. As in Go specification, do not wrap this io.EOF error
// if you want to propagate it as callers should also be able to do err == io.EOF to check the error.
func TestReadRune(t *testing.T) {
	var s ByteStream

	s.Init([]byte("abc"))

	c, err := s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	} else if c != 'a' {
		t.Errorf("expected 'a' but got %c", c)
	}

	c, err = s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	} else if c != 'b' {
		t.Errorf("expected 'b' but got %c", c)
	}

	c, err = s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	} else if c != 'c' {
		t.Errorf("expected 'c' but got %c", c)
	}

	_, err = s.ReadByte()
	if err != io.EOF {
		t.Errorf("expected io.EOF, got %s instead", err.Error())
	}
}

// UnreadRune implements the io.RuneUnreader interface.
//
// Errors:
//   - *ErrUnreadRune: When no previous rune was read.
func TestUnreadRune(t *testing.T) {
	var s ByteStream

	s.Init([]byte("abc"))

	_, err := s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	}

	err = s.UnreadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	}

	c, err := s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	} else if c != 'a' {
		t.Errorf("expected 'a' but got %c", c)
	}
}

// Pos returns the current position in the stream.
//
// Returns:
//   - int: The current position in the stream.
func TestPos(t *testing.T) {
	var s ByteStream

	s.Init([]byte("abc"))

	_, err := s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	}

	_, err = s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	}

	pos := s.Pos()

	if pos != 2 {
		t.Errorf("expected 2 but got %d", pos)
	}

	_ = s.UnreadByte()

	pos = s.Pos()

	if pos != 1 {
		t.Errorf("expected 1 but got %d", pos)
	}
}

// IsExhausted checks if the stream is exhausted.
//
// Returns:
//   - bool: True if the stream is exhausted, false otherwise.
func TestIsExhausted(t *testing.T) {
	var s ByteStream

	s.Init([]byte("abc"))

	_, err := s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	}

	_, err = s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	}

	_, err = s.ReadByte()
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	}

	exhausted := s.IsExhausted()

	if !exhausted {
		t.Error("expected true but got false")
	}
}
