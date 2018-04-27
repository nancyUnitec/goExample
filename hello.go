//hello.go   
package main
// import "fmt"//实现格式化的I/O  

import ( 
    "encoding/json" 
    "fmt" 
    "io" 
    "log" 
    "strings"
    "os" 
) 
 

func testDecoder(){
	const jsonStream = ` 
        { "Name" : "Ed" , "Text" : "Knock knock." } 
        { "Name" : "Sam" , "Text" : "Who's there?" } 
        { "Name" : "Ed" , "Text" : "Go fmt." } 
        { "Name" : "Sam" , "Text" : "Go fmt who?" } 
        { "Name" : "Ed" , "Text" : "Go fmt yourself!" } 
        `

        type Message struct { 
        	Name , Text string 
    	} 

		dec := json. NewDecoder ( strings. NewReader ( jsonStream ) ) 
		    for { 
		        var m Message 
		        if err := dec. Decode ( & m ) ; err == io. EOF { 
		            break 
		        } else if err != nil { 
		            log . Fatal ( err ) 
		        } 
		        fmt. Printf ( "%s: %s \n " , m. Name , m. Text ) 
		    } 

       fmt.Printf("Hello,GO!\n")   

}

//struct to json
func testMarshal(){
    fmt.Printf ( "mmmmmmmmmmmmmmmmmmmarshal \n" ) 
    type ColorGroup struct { 
        ID     int 
        Name   string 
        Colors [ ] string 
    } 
    group := ColorGroup { 
        ID :     1 , 
        Name :   "Reds" , 
        Colors : [ ] string { "Crimson" , "Red" , "Ruby" , "Maroon" } , 
    } 
    b , err := json. Marshal ( group ) 
    if err != nil { 
        fmt. Println ( "error:" , err ) 
    } 
    os. Stdout . Write ( b ) 
    //output is a json format: {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
}

//json to struct
func testUnMarshal(){
    fmt.Printf ( "\n uuuuuuuuuuuuunmarshal \n" ) 
    var jsonBlob = [ ] byte ( ` [ 
        { "Name" : "Platypus" , "Order" : "Monotremata" } , 
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" } 
    ] ` ) 
    type Animal struct { 
        Name  string 
        Order string 
    } 
    var animals [ ] Animal 
    err := json. Unmarshal ( jsonBlob , & animals ) 
    if err != nil { 
        fmt. Println ( "error:" , err ) 
    } 
    fmt.Printf ( "%+v" , animals ) 
    //output: [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorphia}]
}
/*Printf someting*/  
func main(){   
	testDecoder()
	testMarshal()
	testUnMarshal()
}