package api

import "flag"

func FlagStr(name, alias, defaultValue, usage string) *string {
	value := flag.String(name, defaultValue, usage)
	aliasValue := flag.String(alias, defaultValue, `Alias for "--`+name+`"`)
	if *value == defaultValue {
		value = aliasValue
	}
	return value
}

