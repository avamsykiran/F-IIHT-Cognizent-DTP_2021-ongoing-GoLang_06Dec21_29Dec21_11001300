
IDE:        VS Code     https://code.visualstudio.com/download
Compile:    Go          https://golang.org/dl/
--------------------------------------------------------------

Go Lang
-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

        GPL

Features

        1. Simple Familiar Syntax.
        2. Concurrent
        3. Resource Intrinsic
        4. Distributed
        5. Compiler + Linker

Package         is a logical block of code.
                is a go approach to modularization.

Data Types:
-------------------------

    Primitive:
        Numeric:
            uint8   unsigned 8-bit integers (0 to 255)
            uint16  Unsigned 16-bit integers (0 to 65535)
            uint32  Unsigned 32-bit integers (0 to 4294967295)
            uint64  Unsigned 64-bit integers (0 to 18446744073709551615)
            int8    Signed 8-bit integers (-128 to 127)     
            int16   Signed 16-bit integers (-32768 to 32767)	
            int32   Signed 32-bit integers (-2147483648 to 2147483647)
            int64   Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
            
            byte    same as uint8
            rune    same as int32
            uint    32 or 64 bits
            int     same size as uint
            uintptr an unsigned integer 
                    to store the uninterpreted bits of a pointer value

            float32     IEEE-754 32-bit floating-point numbers
            float64     IEEE-754 64-bit floating-point numbers
            complex64   Complex numbers with float32 real and imaginary parts
            complex128  Complex numbers with float64 real and imaginary parts

        Non Numeric:
            bool
            string

Escape Sequences
------------------------------------------------

        \n              new line
        \t              tab space
        \r              carrage return
        \\              escape a \
        \"              escape a "

format specifiers
-------------------------------------------------

        %d
        %ld
        %f
        %c
        %s

Operators
-------------------------------------------------

        Binary
                Arithemtic
                        +
                        -
                        *
                        /
                        %
                Relational 
                        ==
                        <
                        >
                        <=
                        >=
                        !=
                Logical
                        &&
                        ||
                        !
                Assignment Operatos
                        =
                        +=              a+=b;           a = a + b;
                        -=
                        *=
                        /=
                        %=

                        &=
                        |=
                        ^=
                        ~=
        Bitwise
                &               0&0 1&0 0&1             0
                                1&1                     1
                
                |               0|0                     0
                                1|1 0|1 1|0             1

                ^               0^0 1^1                 1
                                0^1 1^0                 0

                ~               ~0                      1
                                ~1                      0

        Unary
                        ++
                        --

                        var x,y,z int;

                        x = 10;

                        y = ++x;    // x->11, y=11
                        z = x++;    // z->11, x->12

        Short hand Declaration Operator

                :=
                
                var a int;      //declaration
                a = 45;         //assignment

                var a int = 45; //initialization
                
                a := 45; //initialization, type is infered.

Decision Making Control Structures
---------------------------------------------

                if (bool_exp) {

                }

                if (bool_exp) {
                        //true blockm of code
                }else{
                        //false block of code.
                }

                Switch Statement
                ------------------------------
                                switch(exp){
                                        case val1:
                                                code-blcok-1;
                                        case val2:
                                                code-block-2;
                                        case val3:
                                                code-block-3;

                                        default:
                                                defalut-code-block;
                                }

                                switch {
                                        case exp:
                                                code-1;
                                        case exp2:
                                                code-2;
                                        case exp3:
                                                code-3...;

                                        default:
                                                defualt-code;
                                }

                                switch x.(type) {
                                        case type1:
                                         code...
                                        case type2:
                                         code..
                                        defualt:
                                         default-code..
                                }

                                select {
                                        case exp:
                                                code-1;
                                        case exp2:
                                                code-2;
                                        case exp3:
                                                code-3...;

                                        default:
                                                defualt-code;
                                }
                                

Looping Controls Strucutres
-------------------------------------

                loop goto

                for 

                        Indefinite loop
                                
                                for bool_exp {
                                        ...............
                                }

                        Definite loop

                                for i:=initalVal;bool_exp;inc/dec {
                                        ...............
                                }

Non-Conditional Control Statements

                        break
                        continue
                        goto 

        Functions
        ---------------------------------------------

                sub-programs - self contianed block of code-
                targeting a specific action/operation.

                func funName(param1 type,param2 type..)  {
                        
                }

                func funName(param1 type,param2 type..) returnType {
                        return soemthing;
                }

                func funName(param1 type,param2 type..) (returnType1,returnType2) {
                        return value1,value2;
                }

                Anonymous Functions
                        function with no name
                
                High Order Functions
                        function that return another function

                Closures
                        High order function that can
                        encapsulate a persisting state.
                

        Arrays
        -----------------------------------------------------------

                var arr [size]type

                var nums [10]int

                var nums = [5]int{1,2,3,4,5}

                len(array)              number of eles
                cap(array)              same as Len

                nums[4] ----------> 5

                range           this is keyword that returns 
                                        the index and ele from the 
                                        givne array until no more eles are left.

        Slice
        ------------------------------------------------------------
                 slice is a small part of an array.

                 s1 := nums[:7]; //s1 contains 0 to 7th eles of the nums array
                 s2 := nums[3:]; //s2 contains 3rd to last ele of the nums array
                 s3 := nums[4:8] //s3 contains 4th to 8th ele of the nums array


                len(slice)              number of eles in that slice
                cap(slice)              number of eles in the array of te slice, starting from the first
                                                ele in the slice

        Assignments
        ----------------------------------------
        Bubble Sort
                           10  5   7   19   8   4   6

                10 - 5      5   10  7   19   8   4   6
                10 - 7      5   7   10  19   8   4   6   
                10 - 19     5   7   10  19   8   4   6
                19 - 8      5   7   10   8  19   4   6
                19 - 4      5   7   10   8   4  19   6
                19 - 6      5   7   10   8   4   6  19

                repeat for 5   7   10   8   4   6

        Selection Sort

                            10  5   7   19   8   4   6

                        selected the ele '10'

                10 - 5          no swap needed
                10 - 7          no swap needed
                10 - 19     19  5   7   10   8   4   6

                        repeat the same with all other ele!

                repeat for 5   7   10   8   4   6

        Quick Sort              //recurisve funcitons
                                  0  1   2   3    4   5   6      
                                 10  5   7   19   8   4   6
                                 d                        u

                10 == @d (10)        d  
                10 > @d (5)              d 
                10 > @d (7)                   d 
                10 < @d (19)                  d

                10 > @u (6)                                u

                d (3) < u (6)             
                                  0  1   2   3   4   5   6      
                                 10  5   7   6   8   4    19
                                             d             u
                10 > @d (6)                      d
                10 > @d (8)                          d
                10 > @d (5)                                d
                10 < @d (19)                               d

                10 < @u (19)                         u
                10 > @u (5)                          u

                d (6) > u (5)
                                4 5 7 6 8 --- 10 --- 19

        RegExp
        ---------------------------------------------------
                Is a language sued for 
                constructing string patterns.

                \s      the matching char must be a space
                \S      the matching char must not be a space
                \w      the matchign char must be a word
                \W      the matchign char must not be a word
                \d      the matching char must be a digit
                \D      the matching char must not be a digit
                .       the matchign char cna be anything

                [A-Z]   the matching char cna be a capital alpha
                [a-z]   the matching char cna be a lower alpha
                [A-za-z]the matching char cna be a capital/lower alpha
                [0-9]   the matching char can be a digit

                [^A-Z]  the matching char can be anything other than A-Z

                [5-7]{3}        three digits ranging from 5 to 7

                [A-D]{0,5}      a string of 0 to 5 chars in length and must comprise of
                                        A to D only.

                +               1 or more occurences
                *               0 or more occurences
                ?               0 or 1 occurence

        User Defined Data Types
        -------------------------------------------------
        types

                structures
                ------------------------------------------------
                        is a user defiend data types
                        that can represent the properties
                        of an entity as members.

                        A Loan has
                                a principal
                                a timePeriod
                                a rateOfInterest
                                a simpleInterest

                        type loan struct {
                          p,t int
                          r float64
                        }

                        func (l loan) simpleInterest() float64 {
                                return (l.p * l.t * l.r) / 100
                        }
                        
                        var l1,l2 loan

                        l1.p = 789;

                        var l3 loan = loan{
                                p:890,
                                t:6,
                                r:0.10
                        }

                        var l4 loan = loan{1000,10,10.0}

                        Anonymous Fields / members
                        ----------------------------

                        type dummy struct{
                                string
                                bool
                                int
                        }

                        d := dummy{"Apple",true,450}

                        dummy.bool = false;

                        Anonymous Structures
                        ------------------------------

                                a structur with no name

                                st := struct {
                                        field1 string
                                        field2 int
                                        field3 bool
                                }{
                                        field1: "soem value",
                                        field2: 56,
                                        field3: false,
                                }

                        Nested Structures / Embeded Structres
                        ------------------------------------------------

                                as an a excersise for composition/inheretence

                                type str1 struct{
                                        f1 int
                                }

                                type str2 struct{
                                        f2 string
                                        s1 str1
                                }

                                v1 := str1{90};
                                v2 := str2{
                                        f2:"Apple",
                                        s1: str1{120},
                                };

                                v2.f2 = "mangoes";
                                v2.s1.f1 =990;

                                composition is excersised through nested structures
                                inheretence is excersised through nested structures 
                                                        + anonymous fields 
                                                        + field promotion 

                Assignment
                ----------------------------------

                        bankAccount     accnum,bal,branch,ifsc

                        address         doorNum,street,city

                        vendor          vid,title,account,contactAddress

                        customer        cid,title,deliveryAddress

                interfaces
                ------------------------------------------------
                        is a user defined data type that
                        has a set of method declarations that
                        can be implemented by structures.

                        interfaces provide abstraction and polymorphisim.

                maps
                --------------------------------------------------------

                        a map is a key-value pair collection.

        Pointers
        -------------------------------------------------
        
                a pointer is a variable that can
                store the address of another variable.

                v1 := 45

                fmt.Println(v1) ----------------> 45
                fmt.Println(&v1) ----------------> 'address of v1'

                vptr := &v1

                fmt.Println(vptr) ----------------> 'address of v1' 
                fmt.Println(&vptr) ----------------> 'address of vptr'
                fmt.Println(*vptr) ----------------> 45

                Call By Reference
                Dynamic Memory Allocation

        defer, panic and recover
        ---------------------------------------------------------------------

                defer is a keyword that postpones the execution of 
                a function call, until the caller function completes.

                panic is a built in function used to abruptly
                end the execution of the current function and raise
                an error....

                recover is a built in function used to receive, and error
                if one exits, when ever recover() is called, we get 'nil'
                if no error occured, ore else it retrive the value passed to
                the panic that occured.

        Concurrent Programming / asynchronous programming
        ---------------------------------------------------------------------


                Multi Tasking / Multi Processing                Multi Thread
                =========================================================================
                OS  manages multi-tasking                       OS to mange multi-thread

                executing individual apps parellalelly          executing sub-routines of 
                                                                the same app parellelly 
                
                Each process has a separate                     Each thread has separate
                and is isolated always.                         thread-context and
                                                                they share the same heap.
                context =                                              thread-context
                        programCount                                            programCount
                        stack                                                   stack
                        heap                                                    registries
                        registreis


                go functionCall()

                time package
                             sleep method is used to pause our current thread

                Channels

                        a channel is a stream of data 
                        into which one can send 
                        and other can receive 

                        both sender and receiver has to wait,
                        until the opponent is done.

                        <-      channel operator

                        ch := make(chan int)

                        ch<-90          //writing into the ch

                        x:= <-ch        //reading from the channel

                        x,ok:= <-ch    //reading from the channel,
                                        //ok is true if channel is still open.

                select-case statement

                        select {
                            case <-ch1:
                                dosoemthing..........
                            case <-ch2:
                                soemthing else.....    
                        }

                signal
                                is any channel that does not produce data,
                                but is used only to inform the compeltion
                                of a goroutine...     

                Synchronization
                -------------------------------------------------------

                        control the acess of multiple threads to shared memory/ data

                        
                Monitor and Lock Machanisim

                        resource is locked as and
                        when a thread starts usign it.

                        the resource will not available for any other thread.

                        a monitor ensures that the resource is unlocked
                        once the current thread accomplishes its job.

                Wait and notify machanisim / Producer-Consumer Pattern

                        a Bucket of water
                                emp1 should fill the bucket using a mug
                                emp2 should empty the bucket using a equal sized mug

                                1) If the bucket is empty:
                                        emp2 should wait until the bucket is refiled\
                                2) If the bucket is full:
                                        emp1 should wait until the bucket is a little emptied
                                3) If the bucket is full some arbitary level
                                        emp1 should fill the bucket using a mug
                                        emp2 should empty the bucket using a equal sized mug
                                
                        the current thread can choose to wait
                        even before the job is done, and 
                        the other waiting threads are notified...!

                        so that the locked resource cna be realeased
                        for the next witing thread,

                        The channels are by default synchroinized
                        and hence the wait / notify is always
                        available no extra cost.



        Assignemtn - ContactBook
        --------------------------------------------

        develop a go lang console basd application to manage the contact records.

                Contact
                        mobile          must be unique
                        first name
                        last name
                        email

                The app must be a menu driven app :
                        1. Creating a contact
                        2. List all contacts
                        3. Delete a contact
                        4. sort contacts by first name
                        5. sort contacts by last name
                        6. sort contacts by mobile
                        7. quit

                the contact must be maintained in a arrray / map / list (your choice)
                that lsit / map / array mus tbe writtne into a file before termianting the app
                and must be read in to the memory at the loading time of the app.


        Go Lang Custom Packages
        ---------------------------------------

                go env -w GO111MODULE=off

                this cmd is sued to switch the manditory module
                of go 'off' so that the packages can be used
                with out converting them into modules.

                Workspace
                -------------------------
                        src             *.go
                        ----------
                                pack1
                                        pack1.go
                                pack2
                                        pack2.go
                                app
                                        app.go


                        bin             *.exe
        
                
                go run          compile and execute the program
                                        but no output is written

                go build        compiles and write the exe file in the same folder

                go install      compile and write the exe into a path pointed by
                                                'gobin'

                set gobin=F:\IIHT\Cognizent\DTP_2021\GoLang_11May2021_28May2021_09001130\GoLang\09_Workspaces\bin

        Mulit-Layer
        ---------------------------------

                DAO                             any database connectivity
                Service                         hold any bussiness logic
                Controller and View             flow of the app and view
                Model                           hold the data and make it
                                                travel amongst the other layers
                                                        to share the data.
        
        Go Modules
        -------------------------------------------------------------------------

                A go module is a collection of
                go packages. (1.11)

                        go env -w GO111MODULE=on
                                        will enable the modules feature.

                        go.mod                  meta,dependency list

                        go mod init <modName>
                                        modName msut be mentioned if we
                                        have not created a folder by that name.

                Workspace structure
                ----------------------------------------------

                        ENV VARIABLES

                                GOROOT          points to the go instalaltion directory
                                GOPATH          point to the workspace, where
                                                        all our project lie.

                                                by default /users/vamsy/go/....
                                GOBIN           points to the directory where
                                                go intall tool whall create the .exe file.
                                                        if GOPATH is not set.

                        GOPATH\src              expected to contain all go packages and src code
                        GOPATH\bin              expected to contain all .exe....
                        GOPATH\pkg              expected hold thrid party lib...


                        a typical workspace hirarchy
                        -----------------------------------------------------

                        GOPATH=d:\MyGoWorkspace

                        d:\MyGoWorkspace
                        -----------------------
                                bin
                                pkg
                                src
                                 |-app1
                                 |   |-mypack
                                 |   |   |-packfile1.go  the mypack package
                                 |   |   |-packfile2.go  the mypack package
                                 |   |-app1.go            the main package
                                 |   |-go.mod
                                 |-app2                   the main package
                                 |   |-app2.go
                                 |   |-go.mod
                                 |-app3
                                 |   |-app3.go          the main package
                                 |   |-go.mod


        REST api
        ---------------------------------------------------------------
        
        /contacts

        Create          POST            /contacts               request body
        Retrive         GET             /contacts
        Retrive         GET             /contacts/101
        Update          PUT             /contacts/101           request body
        Deleting        DELETE          /contacts/101
       

                encoding/json            inbuilt package used for parsing json to golang objects and viceversa


                github.com/gorilla/mux   thrid party router package to host rest api.

                go get github.com/gorilla/mux

                PostMan		as windows app / as chrome extension
                Insomnia	as windows app

        Revel
        ---------------------------------------------------------------------------
        
        Setting the paths
                On Windows
                        set GOPATH=Your\WorkSpace\Path
                        set PATH=%PATH%;%GOPATH%\bin;
                
                On Linux
                        export GOPATH="Your/WorkSpace/Path"
                        export PATH="$PATH:$GOPATH/bin"

        Installing the Framework and Command Line Tool
                go get github.com/revel/revel
                go get github.com/revel/cmd/revel
                
                go get github.com/revel/modules/static
                go get github.com/garyburd/redigo/redis
                go get github.com/patrickmn/go-cache
                go get github.com/bradfitz/gomemcache/memcache

        Creating the project
                revel new app-name

        Execute the proejct
                revel run -a app-name


        SQL Package
        -------------------------------------------------

                RDBMS   ----- DRIVER ------   APP (golang)

                database/sql

                       con,err := sql.Open("drivename","database whereabouts and credits") //verify conn

                       defer con.clsoe();

                       con.Ping();      //open the connection for us...

                       res,err := con.Exec("sql statement",...params)

                       row,err := con.QueryRow("select staements",..params)
                       
                       rows,err := con.Query("select staements",..params)

                       row.Scan(...............)

                Postgre SQL driver
                ==============================================
                        
                        driver name "postgres"
                        go get github.com/lib/pq




        