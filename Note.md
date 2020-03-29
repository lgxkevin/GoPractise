### https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1s
- Shadowing variable  
Same name variables can be declared in **different** scopes.  

- All variable **must** be used  
- Type Conversions  
1.DestinationType(variable)  
2.Use strconv package for strings

---
--- 57:08 ---  
- Default value of boolean varaible is **false**
- Strings  
    1.Immutable  
    2.Can be concatenated with plus (+) operator  
    3.Can be converted to []byte
    
--- 
--- 1:26:29 ---
### Constants
- Constant type  
```
const a int = 42
const b string = 42
const c float32 = 42
const d bool = 42
```
- Constant can be shadowed
- iota  
1. Speical symbol **iota** allows related constants to be created easily
2. **Iota** starts at 0 in each **const** block and increments by one
3. Watch out of constant values that match zero values for variables 
![alt text](./Image/Note/Annotation%202020-03-30%20121930.jpg)
---
--- 1:47:57 ---
### Arrays and Slices
