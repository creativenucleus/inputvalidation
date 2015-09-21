package inputvalidator

import (
	"reflect"
	"strings"
)

// DecodeTag returns a map[string]string of tags for a given field on a struct
// Examples of struct tagstrings (from structTag.Get("tagName")):
//	omitempty
//	string,required
//	maxlength
func DecodeTag(structTag reflect.StructTag, tagName string) map[string]string {
	tags := make(map[string]string)
	
	tagString := structTag.Get(tagName)
	if(tagString == "") {
		return tags
	}
		
	// NB This will break any individual tags with comma in them
	for _, kvpairGlued := range strings.Split(tagString, ",") {
		kvPair := strings.Split(kvpairGlued, "=")
		switch len(kvPair) {
		case 1:
			tags[kvPair[0]] = kvPair[0];			
		case 2:
			tags[kvPair[0]] = kvPair[1];
		default:
			key, value := kvPair[0], kvPair[1:]
			tags[key] = strings.Join(value, "=")
		}
	}

	return tags
}
