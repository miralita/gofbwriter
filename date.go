package gofbwriter

import (
	"fmt"
	"time"
)

/*  <xs:complexType name="dateType">
    <xs:annotation>
      <xs:documentation>A human readable date, maybe not exact, with an optional computer readable variant</xs:documentation>
    </xs:annotation>
    <xs:simpleContent>
      <xs:extension base="xs:string">
        <xs:attribute name="value" type="xs:date" use="optional"/>
        <xs:attribute ref="xml:lang"/>
      </xs:extension>
    </xs:simpleContent>
  </xs:complexType>*/
//A human readable date, maybe not exact, with an optional computer readable variant
type date time.Time

func (s *date) ToXML() (string, error) {
	dt := time.Time(*s).String()
	return fmt.Sprintf(`<date value="%s">%s</date>`, dt, dt), nil
}

func (s *date) tag() string {
	return "date"
}
