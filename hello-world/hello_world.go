/*
package greeting implements an answer to the user
It greets the user by its name if he type it or greets the world instead
*/
package greeting

const testVersion = 3

// HelloWorld greets the user name or the worlds instead if empty.
func HelloWorld(name string) (greet string) {
	if name == "" {
		greet = "Hello, World!"
	} else {
		greet = "Hello, " + name + "!"
	}
	return greet
}
