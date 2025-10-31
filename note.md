# Example: this declare is so messy, please clean it up

## NOTE: only struct with one field and field name is "Value"!!!! And type of "Value" ==:
- utils.INTEGER
- utils.OCTETSTRING
- utils.BOOLEAN
- is array/list: []....

## BUT: if type of "Value" == "util.ENUMERATED". Dont touch/edit these type.


```golang
type NrNsPmaxvalue struct {
	Additionalpmax             *PMax
	Additionalspectrumemission Additionalspectrumemission
}

// NOTE: only struct with one field and field name is "Value"!!!!
type Additionalspectrumemission struct {
	Value utils.INTEGER `lb:0,ub:7`
}
```

Lets combine them into one:
```golang
type NrNsPmaxvalue struct {
	Additionalpmax             *PMax
	Additionalspectrumemission utils.INTEGER `lb:0,ub:7`
}
```

# But after replace, if the field have type which replaced contain more than "*" - mean pointer (like "**" and "***", ....) -> do not change and do not delete this single-field-name-Value struct!