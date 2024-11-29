package strcase_test

import (
	"fmt"

	"github.com/xo/ox/strcase"
)

func Example() {
	fmt.Println("Change CamelCase -> snake_case:", strcase.CamelToSnake("AnIdentifier"))
	fmt.Println("Change CamelCase -> snake_case (2):", strcase.CamelToSnake("XMLHTTPACL"))
	fmt.Println("Change snake_case -> CamelCase:", strcase.SnakeToCamel("an_identifier"))
	fmt.Println("Force CamelCase:", strcase.ForceCamelIdentifier("APoorly_named_httpMethod"))
	fmt.Println("Force lower camelCase:", strcase.ForceLowerCamelIdentifier("APoorly_named_httpMethod"))
	fmt.Println("Force lower camelCase (2):", strcase.ForceLowerCamelIdentifier("XmlHttpACL"))
	fmt.Println("Change snake_case identifier -> CamelCase:", strcase.SnakeToCamelIdentifier("__2__xml___thing---"))
	// Output:
	// Change CamelCase -> snake_case: an_identifier
	// Change CamelCase -> snake_case (2): xml_http_acl
	// Change snake_case -> CamelCase: AnIdentifier
	// Force CamelCase: APoorlyNamedHTTPMethod
	// Force lower camelCase: aPoorlyNamedHTTPMethod
	// Force lower camelCase (2): xmlHTTPACL
	// Change snake_case identifier -> CamelCase: XMLThing
}
