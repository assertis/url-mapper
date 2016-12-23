# Go URL Mapper

Package mapping query string to struct using reflection and tags.
Support for optional validation of input values.

Example:

``
type Request struct {
    Origin string `query:"origin,regexp=^[A-Z]{3}$"`
    Destination string `query:"destination,regexp=^[A-Z]{3}$"`
    Adults int `query:"adults,default=1,max=9"`
    Children int `query:"children,optional,default=0,max=9"`
    Outward time.Time `query:"outward,dateFormat=RFC_3339"`
    Return *time.Time `query:"return,optional,dateFormat=RFC_3339"`
}``
