package main

import (
	"github.com/bitly/go-simplejson"
	"fmt"
)

func main() {


	js, err := simplejson.NewJson([]byte(`{
		"test": {
			"string_array": ["asdf", "ghjk", "zxcv"],
			"string_array_null": ["abc", null, "efg"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"subkeyone": 1},
			{"subkeytwo": 2, "subkeythree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true,
			"sub_obj": {"a": 1}
		}
	}`))

	if err != nil{

	}

	js.Get("test").Set("float",1.11)

	fmt.Println(js.GetPath("test","float1").Float64())

}
