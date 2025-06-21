//John 3:16

package agreements

import(
    "fmt"
    "bcolors"
)

func Covenant() {
    fmt.Printf(`%s
[ MIT License Copyright (c) 2024 r0jahsm0ntar1............]
[ Permission is hereby granted, free of charge, to any....]
[ person obtaining a copy of this software and associated.]
[ documentation files (the "Software"), to deal in the....]
[ Software without restriction, including without.........]
[ limitation the rights to use, copy, modify, merge,......]
[ publish, distribute, sublicense, and/or sell copies of..]
[ the Software, and to permit persons to whom the Software]
[ is furnished to do so, subject to the following.........]

[ Conditions🍄:...........................................]
[ The above copyright notice and this permission notice...]
[ shall be included in all copies or substantial portions.]
[ of the Software.........................................]

[ THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF...]
[ ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED.]
[ TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A.....]
[ PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT.....]
[ SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY]
[ CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION.]
[ OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR.]
[ IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER.....]
[ DEALINGS IN THE SOFTWARE................................]
%s
[ The 🦊Africana-Framework is designed purely for Good &..]
[ not evil. If you are planning on using this tool for..🐝]
[ malicious purposes that are not authorized by the.......]
[ 🥝company you are performing assessments for, you are...]
[ violating the terms of service and license of this......]
[ toolset. By hitting yes (only one time), you agree to...]
[ the terms of service and that you will only use this....]
[ tool for lawful purposes only.........................🎲]
%s`, bcolors.Green, bcolors.BrightRed, bcolors.Endc)
    fmt.Println(bcolors.Endc + `
¯\_(ツ)_/¯ ` + bcolors.Cyan + `🥥Do you agree to the terms of service... ` + bcolors.Yellow + `(y/n)
` + bcolors.Endc)
}
