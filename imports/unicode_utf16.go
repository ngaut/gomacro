// this file was generated by gomacro command: import _b "unicode/utf16"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"unicode/utf16"
)

// reflection: allow interpreted code to import "unicode/utf16"
func init() {
	Packages["unicode/utf16"] = Package{
	Binds: map[string]Value{
		"Decode":	ValueOf(utf16.Decode),
		"DecodeRune":	ValueOf(utf16.DecodeRune),
		"Encode":	ValueOf(utf16.Encode),
		"EncodeRune":	ValueOf(utf16.EncodeRune),
		"IsSurrogate":	ValueOf(utf16.IsSurrogate),
	}, 
	}
}
