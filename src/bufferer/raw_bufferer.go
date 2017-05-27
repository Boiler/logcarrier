package bufferer

import (
	"fileio"
	"logio"
)

// RawBufferer ...
type RawBufferer struct {
	l *logio.Writer
	d *fileio.File
}

// NewRawBufferer constructor
func NewRawBufferer(l *logio.Writer, d *fileio.File) *RawBufferer {
	return &RawBufferer{
		l: l,
		d: d,
	}
}

// Write implementation
func (b *RawBufferer) Write(p []byte) (n int, err error) {
	return b.l.Write(p)
}

// Close implementation
func (b *RawBufferer) Close() error {
	if err := b.l.Flush(); err != nil {
		return err
	}
	if err := b.d.Close(); err != nil {
		return err
	}
	return nil
}

// Flush implementation
func (b *RawBufferer) Flush() error {
	if b.l.WorthFlushing() {
		if err := b.l.Flush(); err != nil {
			return err
		}
	}
	return nil
}
