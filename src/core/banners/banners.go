package banners

import (
    "fmt"
    "time"
    "bcolors"
    "math/rand"
)
var (
    version = "v2.0.6"
)

type Banners struct{}

func Graphics() {
    rand.Seed(time.Now().UnixNano())
    menu := rand.Intn(8) + 1

    switch menu {
    case 1:
        fmt.Printf(`%s
                     _,._
                 __.'   _)
                <_,)'.-"a\
                  /' (    \
      _.-----..,-'   ('"--^
     //              |
    (|   ';      ,   :
      \   ;.----/  , /
       ) // /   | |\ \         %s%s%s%s%s
       \ \\'\   | |/ /      %sJesus Christ%s
        \ \\ \  | |\/  %sThe Lamb that was slain%s
         '" '"  '""         %sfor our sins.%s
                             %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 2:
        fmt.Printf(`%s
     _      xxxx      _
    /_;-.__ / _\  _.-;_\
       '-._ ''_/''.-'
           '\   /'
            |  /
           /-.(
           \_._\
            \ \';           %s%s%s%s%s
             > |/        %sJesus Christ%s
            / //       %sL❤️.VE'S U. He is%s
            |//    %sThe Lamb that was slain%s
            \(\          %sfor our sins.%s
                           %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.PURPLE, bcolors.PURPLE, bcolors.ENDC)

    case 3:
        fmt.Printf(`%s
                               .--,
                           ,.-( (o)\
                          /   .)/\ ')
                        .',./'/   )/
                    ()=///=))))==()
                      /
                            %s%s%s%s%s
                      %swake up, Christian%s
              %sLord God Jesus Christ L❤️.VE'S you%s
                    %sfollow the white Pigeon.%s
                     %sknock, knock, knock,%s
                          %sMan Of God.%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.DARKCYAN, bcolors.GREEN, bcolors.PURPLE, bcolors.ENDC)

    case 4:
        fmt.Printf(`%s
                                       ___
             _______                  /__/
            |.-----.|            ,---[___]*
            ||     ||           /    printer
            ||_____||    _____ /        ____
            |o_____*|   [o_+_+]--------[=i==]
             |     ________|           drive
             |  __|_            %s%s%s%s%s
             '-/_==_\       %sJesus Christ is%s
              /_____\\  %sThe Lamb that was slain%s
                             %sfor our sins.%s
                               %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 5:
        fmt.Printf(`%s
           __________   __________ __________
          |          |\|          |          |\
          |  *    *  |||  *  *  * |        * ||
          |  *    *  |||          |     *    ||
          |  *    *  |||  *  *  * |  *       ||
          |__________|||__________|__________||
          |          || '---------------------'
          |  *    *  ||         %s%s%s%s%s
          |          ||     %sJesus Christ is%s
          |  *    *  || %sThe Lamb that was slain%s
          |__________||      %sfor our sins.%s
           '----------'        %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 6:
        fmt.Printf(`%s
                          ____
                   __,-~~/~   '---.
                 _/_,---(      ,   )_
               _/        <    /   )  \_
  - -----==;;;'==--------------------==';;;==----- -
               \/  ~"~"~"~"~"~\~"~)~"/
               (_ (   \  (     >    \)
                 \_( _ <         >_>'     %s%s%s%s%s
                    ~ '-i' ::>|--"    %sJesus Christ is%s
                        I;|.|.|   %sThe Lamb that was slain%s
                       <|i::|i|'.      %sfor our sins.%s
                      (' ^'"'-' ")       %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 7:
        fmt.Printf(`%s
                .--.       .--.
            _  '    \     /    '  _
             '\.===. \.^./ .===./'
                    \/'"'\/
                 ,  |     |  ,
                / '\|'-.-'|/' \
               /    |  \  |    \
            .-' ,-''|   ; |''-, '-.
                |   |    \|   |
                |   |    ;|   |          %s%s%s%s%s
                |   \    //   |      %sJesus Christ is%s
                |    '._//'   |  %sThe Lamb that was slain%s
               .'             '.      %sfor our sins.%s
            _,'                 ',_     %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 8:
        fmt.Printf(`%s
                   _,.---.---.---.--.._ 
               _.-' '--.'---.'---'-. _,'--.._
              /'--._ .'.     '.     ','-.'-._\
             ||   \  '.'---.__'__..-'. ,''-._/
        _  ,'\ '-._\   \    '.    '_.-'-._,''-.
     ,'   '-_ \/ '-.'--.\    _\_.-'\__.-'-.'-._'.
    (_.o> ,--. '._/'--.-',--'  \_.-'       \'-._ \
     '---'    '._ '---._/__,----'           '-. '-\
               /_, ,  _..-'        %s%s%s%s%s    '-._\
               \_, \/ ._(      %sJesus Christ is%s
                \_, \/ ._\ %sThe Lamb that was slain%s
                 '._,\/ ._\     %sfor our sins.%s
                   '._// ./'-._   %sJohn 3:16%s
                     '-._-_-_.-'%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.BLUE, bcolors.ENDC)
    }
}

func anonymas() {
     banner := `
                    .;1tfLCL1,
                   .,,..;i;f0G;
                         ,:,tCC.  ...
                         ;i:fCL,1LLtf1i;,
                       .,::tCL1LC1::;, .,,
                       ;1:tCL,tLt,1:
                      ,::tLf, 1Lf;::.
                    .ii:tLt.  .1Lf;i1.
                    ,:;tf1      1ft;::
                 .1;:tf1  ,i1t1,  ift;;1,
                ,i:t;f. ,LLffLL:  tft;i:
                .;:;fff  .LCLLLf, 1ffi:;.
                :fi;Lff1.  ,;;:  ifffi;f;
                 .:::tCLLfi:,,:ifLfLt::;.
                  ,11:1CCCCCLLLLLLf1;1t:
                  .it;:;1fLLLLfft1;:;ti.
                     ,:;::;;;;;;;;;;,
                       .,::::::::,.`
    fmt.Printf("%s%s%s\n", bcolors.DARKGREEN, banner, bcolors.ENDC)
}

func Graphicx() {
    rand.Seed(time.Now().UnixNano())
    menu := rand.Intn(5) + 1

    switch menu {
    case 1:
        anonymas()
        fmt.Printf(`%s
                                                    _____
  _____    ____   ____   ____   ________ __________/ ____\
  \__  \  /    \ /  _ \ /    \ /  ___/  |  \_  __ \   __\
   / __ \|   |  (  <_> )   |  \\___ \|  |  /|  | \/|  |
  (____  /___|  /\____/|___|  /____  >____/ |__|   |__|
       \/     \/            \/     \/               %s%s%s%s

`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)

    case 2:
        anonymas()
        fmt.Printf(`%s
     _______  ________  ________   ________  ________  ________
   ╱╱       ╲╱    ╱   ╲╱        ╲ ╱        ╲╱        ╲╱        ╲
  ╱╱        ╱         ╱         ╱_╱       ╱╱        _╱        _╱
 ╱       --╱         ╱        _╱╱         ╱-        ╱╱       ╱
 ╲________╱╲___╱____╱╲____╱___╱ ╲________╱╲________╱ ╲______╱
                                                    %s%s%s%s

`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)

    case 3:
        anonymas()
        fmt.Printf(`%s

    ██      ▄   ████▄    ▄      ▄▄▄▄▄   ▄   █▄▄▄▄ ▄████
    █ █      █  █   █     █    █     ▀▄  █  █  ▄▀ █▀   ▀
    █▄▄█ ██   █ █   █ ██   █ ▄  ▀▀▀▀▄ █   █ █▀▀▌  █▀▀
    █  █ █ █  █ ▀████ █ █  █  ▀▄▄▄▄▀  █   █ █  █  █
       █ █  █ █       █  █ █          █▄ ▄█   █    █
      █  █   ██       █   ██           ▀▀▀   ▀      ▀
     ▀                                              %s%s%s%s

`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)

    case 4:
        anonymas()
        fmt.Printf(`%s
  ▐▄▄▄▄▄▄ ..▄▄ · ▄• ▄▌.▄▄ ·   ▄▄·  ▄ .▄▄▄▄  ▪  .▄▄ · ▄▄▄▄▄
   ·██▀▄.▀·▐█ ▀. █▪██▌▐█ ▀.  ▐█ ▌▪██▪▐█▀▄ █·██ ▐█ ▀. •██
 ▪▄ ██▐▀▀▪▄▄▀▀▀█▄█▌▐█▌▄▀▀▀█▄ ██ ▄▄██▀▐█▐▀▀▄ ▐█·▄▀▀▀█▄ ▐█.▪
 ▐▌▐█▌▐█▄▄▌▐█▄▪▐█▐█▄█▌▐█▄▪▐█ ▐███▌██▌▐▀▐█•█▌▐█▌▐█▄▪▐█ ▐█▌·
  ▀▀▀• ▀▀▀  ▀▀▀▀  ▀▀▀  ▀▀▀▀  ·▀▀▀ ▀▀▀ ·.▀  ▀▀▀▀ ▀▀▀▀  ▀▀▀ 
                                                    %s%s%s%s

`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)

    case 5:
        anonymas()
        fmt.Printf(`%s
        ______  ________  ________  ________  ________ 
       /      \/        \/        \/    /   \/        \
      /       /         /        _/         /        _/
    _/      //        _/-        /         /-        /
    \______//\________/\________/\________/\________/
                                            %s%s%s%s

`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)

    }
}

func Banner() {
    Graphics()
}

func AnonSurfBanner() {
    Graphicx()
}
