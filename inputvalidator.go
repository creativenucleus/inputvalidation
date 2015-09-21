package inputvalidator

import(
    "net/url"
    "reflect"
    "fmt"
    "strconv"
)


func Filter(in url.Values, out interface{}) (bool, []string) {
    validationErrors := make([]string, 0)
    
    outValues := reflect.ValueOf(out).Elem()
    
    for i := 0; i < outValues.NumField(); i++ {
        var err error
        mandatory := false
        maxlengthIsSet := false
        var maxlength int
        
        fieldValue := outValues.Field(i)
        fieldType := outValues.Type().Field(i)
    
        tags := DecodeTag(fieldType.Tag, "inputvalidator")

        _, ok := tags["mandatory"]
        if ok {
            mandatory = true
        }
        
        v, ok := tags["maxlength"]
        if ok {
            maxlength, err = strconv.Atoi(v)
            if err == nil {
                maxlengthIsSet = true
            }
        }

        valIn, ok := in[fieldType.Name]
        if !ok {
            if(mandatory) {
                validationErrors = append(validationErrors, fmt.Sprintf("Could not find mandatory value on %s", fieldType.Name))
            }
            continue
        }

        strValIn := valIn[0]
        if(maxlengthIsSet && len(strValIn) > maxlength) {
            validationErrors = append(validationErrors, fmt.Sprintf("Field size was greater than %d characters for %s", maxlength, fieldType.Name))
            continue
        }

        fieldValue.SetString(strValIn)
    }
        
    return len(validationErrors) == 0, validationErrors
}
