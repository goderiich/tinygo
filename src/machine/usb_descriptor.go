//go:build sam || nrf52840
// +build sam nrf52840

package machine

type USBDescriptor struct {
	Device        []byte
	Configuration []byte
	HID           map[uint16][]byte
}

func (d *USBDescriptor) Configure(idVendor, idProduct uint16) {
	d.Device[8] = byte(idVendor)
	d.Device[9] = byte(idVendor >> 8)
	d.Device[10] = byte(idProduct)
	d.Device[11] = byte(idProduct >> 8)
}

var descriptorCDC = USBDescriptor{
	Device: []byte{
		0x12, 0x01, 0x00, 0x02, 0xef, 0x02, 0x01, 0x40, 0x86, 0x28, 0x2d, 0x80, 0x00, 0x01, 0x01, 0x02, 0x03, 0x01,
	},
	Configuration: []byte{
		0x09, 0x02, 0x4b, 0x00, 0x02, 0x01, 0x00, 0xa0, 0x32,
		0x08, 0x0b, 0x00, 0x02, 0x02, 0x02, 0x00, 0x00,
		0x09, 0x04, 0x00, 0x00, 0x01, 0x02, 0x02, 0x00, 0x00,
		0x05, 0x24, 0x00, 0x10, 0x01,
		0x04, 0x24, 0x02, 0x06,
		0x05, 0x24, 0x06, 0x00, 0x01,
		0x05, 0x24, 0x01, 0x01, 0x01,
		0x07, 0x05, 0x81, 0x03, 0x10, 0x00, 0x10,
		0x09, 0x04, 0x01, 0x00, 0x02, 0x0a, 0x00, 0x00, 0x00,
		0x07, 0x05, 0x02, 0x02, 0x40, 0x00, 0x00,
		0x07, 0x05, 0x83, 0x02, 0x40, 0x00, 0x00,
	},
}

var descriptorCDCHID = USBDescriptor{
	Device: []byte{
		0x12, 0x01, 0x00, 0x02, 0xef, 0x02, 0x01, 0x40, 0x86, 0x28, 0x2d, 0x80, 0x00, 0x01, 0x01, 0x02, 0x03, 0x01,
	},
	Configuration: []byte{
		0x09, 0x02, 0x64, 0x00, 0x03, 0x01, 0x00, 0xa0, 0x32,
		0x08, 0x0b, 0x00, 0x02, 0x02, 0x02, 0x00, 0x00,
		0x09, 0x04, 0x00, 0x00, 0x01, 0x02, 0x02, 0x00, 0x00,
		0x05, 0x24, 0x00, 0x10, 0x01,
		0x04, 0x24, 0x02, 0x06,
		0x05, 0x24, 0x06, 0x00, 0x01,
		0x05, 0x24, 0x01, 0x01, 0x01,
		0x07, 0x05, 0x81, 0x03, 0x10, 0x00, 0x10,
		0x09, 0x04, 0x01, 0x00, 0x02, 0x0a, 0x00, 0x00, 0x00,
		0x07, 0x05, 0x02, 0x02, 0x40, 0x00, 0x00,
		0x07, 0x05, 0x83, 0x02, 0x40, 0x00, 0x00,
		0x09, 0x04, 0x02, 0x00, 0x01, 0x03, 0x00, 0x00, 0x00,
		0x09, 0x21, 0x01, 0x01, 0x00, 0x01, 0x22, 0x65, 0x00,
		0x07, 0x05, 0x84, 0x03, 0x40, 0x00, 0x01,
	},
	HID: map[uint16][]byte{
		2: []byte{
			// keyboard and mouse
			0x05, 0x01, 0x09, 0x06, 0xa1, 0x01, 0x85, 0x02, 0x05, 0x07, 0x19, 0xe0, 0x29, 0xe7, 0x15, 0x00,
			0x25, 0x01, 0x75, 0x01, 0x95, 0x08, 0x81, 0x02, 0x95, 0x01, 0x75, 0x08, 0x81, 0x03, 0x95, 0x06,
			0x75, 0x08, 0x15, 0x00, 0x25, 0x73, 0x05, 0x07, 0x19, 0x00, 0x29, 0x73, 0x81, 0x00, 0xc0, 0x05,
			0x01, 0x09, 0x02, 0xa1, 0x01, 0x09, 0x01, 0xa1, 0x00, 0x85, 0x01, 0x05, 0x09, 0x19, 0x01, 0x29,
			0x03, 0x15, 0x00, 0x25, 0x01, 0x95, 0x03, 0x75, 0x01, 0x81, 0x02, 0x95, 0x01, 0x75, 0x05, 0x81,
			0x03, 0x05, 0x01, 0x09, 0x30, 0x09, 0x31, 0x09, 0x38, 0x15, 0x81, 0x25, 0x7f, 0x75, 0x08, 0x95,
			0x03, 0x81, 0x06, 0xc0, 0xc0,
		},
	},
}
