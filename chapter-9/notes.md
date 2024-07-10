Modules sets clearn boundaries for your code, ensuring it remains cohesive and self-contained. Makes
code easier to share and reuse through logic isolation and modularity.

Key components of working with Go modules

- go.mod file: The go.mod file is the core of a Go module as it serves as the blueprint for modules and contains information
  such as the module's path, its version, and a list of its dependencies and their versions.
- go.sum file: The go.sum file works in tandem with the go.mod file. It contains the cryptographic checksums for the dependencies
  listed in the go.mod file. These checksums are security measures and validates that the downloaded dependencies produce the same values.
  If not, then it has been corrupted/tampered with.
- versioning of dependencies

The go.mod file
- module path: Path where the current module is expected to be found.
ex) for "module mymodule" specifies the module's path as mymodule
- dependencies: lists of dependencies with their module paths & specific versions or version ranges
- replace directives (optional): allows you to specify replacements for certain dependencies.
  useful for testing or resolving compatibility issues.
- exclude directives (optional): allows you to exclude specific versions of dependency that may have known issues
```
module mymodule
require (
  github.com/some/dependency v1.2.3
  github.com/another/dependency v2.0.0
)
replace (
  github.com/dependency/v3 => github.com/dependency/v4
)
exclude (
  github.com/some/dependency v2.0.0
)
```

In exercise 9.01, we created a module called bookutil that contained an "author" package. Within the "author" package, we have an author struct with methods. In our main package, we imported it for use.

Go module does not need to have the same name as your Go packages since you can have many packages to one module and project. In exercise 9.01's case, the primary purpose of the module is to manage and work with books and authors so the module's name should reflect the broader context.