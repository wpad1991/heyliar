# heyliar : logic parser
many expression and operator can processing
![Alt text](/doc/example1.png)

## Install
1. Go to the [releases][https://github.com/wpad1991/heyliar/releases], download the code. and import your code.
2. import my git. 
    - set GOPATH
    - import "github.com/wpad1991/heyliar"
    - enter into the terminal. "go get" at code path
![Alt text](/doc/example2.png)

## Usage
### Totalparsing
    Totalparsing("{3+2>5}or{33-2==31}")              // true
    Totalparsing("{33<22}or{(32+22*2-3*100)+2*6<0}") // true
    Totalparsing("({33<22}or{32+22>0})and{123>0}")   // true

### Calcliar
    Calcliar("3*5==15")                           // true
    Calcliar("3*5!=15")                           // false
    Calcliar("(3+2)*3+3+(-123)*2/(2+(-1)*7)<200") // true

### Calculate
    Calculate("1+2")             // 3
    Calculate("6*3")             // 18
    Calculate("(3*2)+3-2*(2/2)") // 7