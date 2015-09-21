package inputvalidator

import (
    "testing"
    "net/url"
)


func TestFilterMandatoryPass(t *testing.T) {
    type User struct {
        Username string `inputvalidator:"mandatory"`
    }
    user := User{}
    
    inputs := make(url.Values)
    inputs.Set("Username", "jtruk")

    success, validationErrors, err := Filter(inputs, &user)
    if(err != nil || !success || len(validationErrors) != 0) {
        t.Errorf("Filter: Mandatory requirement should have passed but didn't")
    }
}


func TestFilterMandatoryFail(t *testing.T) {
    user := struct {
        Username string `inputvalidator:"mandatory"`
    } {}
    
    inputs := make(url.Values)

    success, validationErrors, err := Filter(inputs, &user)
    if(err != nil || success || len(validationErrors) != 1) {
        t.Errorf("Filter: Mandatory requirement should have failed but didn't")
    }
}


func TestFilterMaxLengthPass(t *testing.T) {
    user := struct {
        Username string `inputvalidator:"maxlength=10"`
    } {}
    
    inputs := make(url.Values)
    inputs.Set("Username", "jtruk")

    success, validationErrors, err := Filter(inputs, &user)
    if(err != nil || !success || len(validationErrors) != 0) {
        t.Errorf("Filter: Max Length should have passed but didn't")
    }
}


func TestFilterMaxLengthFail(t *testing.T) {
    user := struct {
        Username string `inputvalidator:"maxlength=10"`
    } {}
    
    inputs := make(url.Values)
    inputs.Set("Username", "mynameisjames")

    success, validationErrors, err := Filter(inputs, &user)
    if(err != nil || success || len(validationErrors) != 1) {
        t.Errorf("Filter: Max Length should have failed but didn't")
    }
}