# HelloWorld

print statement is imported from fmt package.

# Values

A data can printed by concat with the string statement or using commas

# Variables

variables are declared using var keyword, the datatype is optional, given after the variable name.

var a int;

a:=1 is same as var a int = 1, := is for shorthand and this can't be used outside function.

# Constants

It is same as variable, but not updateble and const keyword is used.

# Loops

for is used, it is same for while.

there are some different versions of for loop.

# if else

if condition{

} else if condition{

} else { // the else statement must be near the curly bracket

}

# Switches

it can take a argument or it can use other variable while checking the case.

switch variable (optional){

}

# Arrays

var variable_name [size][size]data_type

Accessed by index

# slices

Dynamic arrays.


# Maps

Data Structure for keys and values.

# Functions

- args variable_name datatype

```go
func Add(a int, b int) int{
    return a + b
}
```

- multiple return value

```go
func vals() (int,int){
    return 3,7;
}

func main(){
    a, b := vals();
}
```
- variadic function, n number args

```go
func sum(nums ...int){
    total:=0
    for _,num :=range nums{
        total+=num
    }
    fmt.Println(total)
}

func main(){
    sum(1,2,3,4);
    sum(1,3);

    nums :=[]int{1,2,3,4};
    sum(nums...);
}
```

# Closures

Closures are anonymous functions.

So, we can create a function with in a function and also we can return the function.

```go
func intSeq() func(){
    i:=0;

    return func(){
        i+=1;
        return i;
    }
}

func main(){
    nextInt :=intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newNextInt :=intSeq()
    fmt.Println(newNextInt())
}
```

# Recursion

Function calling itself.

# Range function

To traverse through array, string, map.

# Pointers

Reference of addresses

# Strings, Bytes and Runes

Best reading on Characters, https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/

