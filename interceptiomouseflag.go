package interception

type MouseFlag int

const (
	MOUSE_MOVE_RELATIVE      MouseFlag = 0x000
	MOUSE_MOVE_ABSOLUTE      MouseFlag = 0x001
	MOUSE_VIRTUAL_DESKTOP    MouseFlag = 0x002
	MOUSE_ATTRIBUTES_CHANGED MouseFlag = 0x004
	MOUSE_MOVE_NOCOALESCE    MouseFlag = 0x008
	MOUSE_TERMSRV_SRC_SHADOW MouseFlag = 0x100
	MOUSE_CUSTOM             MouseFlag = 0x200
)
