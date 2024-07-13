Package naming guidelines
- concise names
- lowercased
- avoid generic package names
- non-plural names

It's standard in go to name your packages something short and concise.
ex) 
- strconv
- sync
- time

Also in Go, it's standard to name your package with only lowercase with no underscores. Camel case and snake case are frowned upon.

Abbreviations are encouraged if name is common and familiar in programming community.
ex)
- strconv (string conversion)
- regex (regular express search)
- sync (synchronization)
- os (operating system)

The above are guidelines, but hard rules. However, avoid package names that do not illuminate the purpose of the package.

Package declarations
Every Go file starts with a package declaration. The package declaration is the name of the package. The first line of each file must be the package declaration. 
ex)
```
package <packageName>
```

If a file is within the same package as a different file. It will have access to code from the other files. 

Exported and unexported code
- if you want functions, types, variable, interfaces, etc to be exportable to other packages then you will have to name them Uppercased. This will tell Go that that entity is exported and can be used outside of package. If any variable is lowercased, then it can only be used within the package it's declared in, even if it's within a different file.

Aliasing packages
- package name might not making it easy to understand its purpose. For clarity, it might be better to alias it and use a different name
- package name could be too long and aliasing it will make it more concise
- a scenario can happen where 2 packages have different paths, but the same name. In this case, you need to differentiate them by aliasing.