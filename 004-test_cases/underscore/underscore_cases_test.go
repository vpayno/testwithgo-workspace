package underscore

type CamelTestCase struct {
	arg  string
	want string
}

var camelTestCases = []CamelTestCase{
	{
		arg:  "thisIsACamelCaseString",
		want: "this_is_a_camel_case_string",
	},
	{
		"with a space",
		"with a space",
	},
	{
		"endsWithA",
		"ends_with_a",
	},
}
