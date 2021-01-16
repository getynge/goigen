# goigen
Goigen is a tool for generating interfaces from method lists

For any given type, goigen can generate an interface to which that type complies,
along with mocks for that type (if mockgen is installed)

## Usage
`goigen [options] DIRECTORY RECEIVER INTERFACE`

goigen supports the following options:
* mockdirectory - specifies the directory mocks are written to relative to DIRECTORY
* fileprefix - specifies the prefix added to generated file names

goigen takes three arguments (options MUST come before arguments):
* DIRECTORY - the directory to search for methods in (e.g. `processor/testfiles`)
* RECEIVER - the receiver to match the interface to (e.g. `Example` or `*Example`)
* INTERFACE - the name of the interface to generate (e.g. `IExample`)
