package parser

import (
	. "github.com/phodal/design/languages/design"
)

func NewDesignAppListener() *DesignAppListener {
	return &DesignAppListener{}
}

type DesignAppListener struct {
	BaseDesignListener
}
