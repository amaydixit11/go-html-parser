package dom

type WhatToShowType uint64

const (
	SHOW_ALL                    WhatToShowType = 0xFFFFFFFF
	SHOW_ELEMENT                WhatToShowType = 0x1
	SHOW_ATTRIBUTE              WhatToShowType = 0x2
	SHOW_TEXT                   WhatToShowType = 0x4
	SHOW_CDATA_SECTION          WhatToShowType = 0x8
	SHOW_ENTITY_REFERENCE       WhatToShowType = 0x10 // legacy
	SHOW_ENTITY                 WhatToShowType = 0x20 // legacy
	SHOW_PROCESSING_INSTRUCTION WhatToShowType = 0x40
	SHOW_COMMENT                WhatToShowType = 0x80
	SHOW_DOCUMENT               WhatToShowType = 0x100
	SHOW_DOCUMENT_TYPE          WhatToShowType = 0x200
	SHOW_DOCUMENT_FRAGMENT      WhatToShowType = 0x400
	SHOW_NOTATION               WhatToShowType = 0x800 // legacy
)

type AcceptNodeType uint8

const (
	FILTER_ACCEPT AcceptNodeType = 1
	FILTER_REJECT AcceptNodeType = 2
	FILTER_SKIP   AcceptNodeType = 3
)

type NodeFilter interface {
	acceptNode(node *Node) AcceptNodeType
}
