# butler_go

Clean your castle faster

## Prehistory
This CLI program provides functions for maintaining order on a server or home computer.
Have you ever heard of this ?

> How tired I am of looking at this trash in Downloads.  
> (c) Author

Golang fast enough to fix this problem.

Enjoy.


## How to use
The program provides three functions:  
1. **Cleaning**  
2. **Grouping**  
3. **Archiving**

In case the first and third functions are already familiar, grouping may seem like an interesting function.

**It is important to close slashes.**

### Examples 
**Cleaning**  
`butler.exe --clean /path/to/dir/`  
Clean target dir but not delete it.

**Grouping**  
`butler.exe --group /path/to/dir/`  
For example, you have a lot of *.exe* files, after this command you will have one directory called *ALL.EXE* with all your *.exe* files in it.

**Archiving**  
`butler_go.exe --archive /path/to/dir/`  
Create zip archive in current directory for target directory.


## Python prototype
[butler](https://github.com/CoolCoderCarl/butler)
