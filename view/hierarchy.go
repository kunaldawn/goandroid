// TODO : Documentation

package view

import (
	"encoding/xml"
)

type Hierarchy struct {
	XMLName  xml.Name `xml:"hierarchy"`     // Namespace of the hierarchy
	Rotation string   `xml:"rotation,attr"` // Rotation value of the hierarchy
	NodeList Nodes    `xml:"node"`          // Child nodes in the hierarchy
}

func (hierarchy Hierarchy) ConvertToViews() (Views, error) {
	return hierarchy.NodeList.ConvertToViews()
}
