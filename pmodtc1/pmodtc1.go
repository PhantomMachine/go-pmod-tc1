package pmodtc1

import "golang.org/x/exp/io/spi"

func (p *PmodTC1) Open() (func() error, error) {
	spidev, err := spi.Open(&spi.Devfs{
		Dev:      p.device,
		Mode:     spi.Mode3,
		MaxSpeed: 500000,
	})
	if err != nil {
		return nil, err
	}

	p.tx = spidev

	return spidev.Close, nil
}

func New(device string) *PmodTC1 {
	return &PmodTC1{
		device,
		nil,
	}
}

type PmodTC1 struct {
	device string
	tx     interface {
		Tx(w, r []byte) error
	}
}

func (p *PmodTC1) ReadTemp() float32 {
	rd := make([]byte, 4)
	if err := p.tx.Tx([]byte{
		0, 0, 0, 0,
	}, rd); err != nil {
		panic(err)
	}

	// convert 4 bytes into a raw 32 bit integer
	var raw uint32 = 0
	for i, word := range rd {
		raw |= uint32(word)
		if i != len(rd)-1 {
			raw <<= 8
		}
	}

	// only top 14 bits of raw are used, cut the rest
	buf := raw & 0xFFFC0000
	buf >>= 18

	// convert uint32 to a float32
	mod := buf % 4
	temp := float32(buf) / 4.0

	temp = temp + (float32(mod) * .25)

	return temp
}
