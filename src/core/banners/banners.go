package banners

import (
    "time"
    "fmt"
    "bcolors"
    "math/rand"
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
       ) // /   | |\ \
       \ \\'\   | |/ /      %sJesus Christ%s
        \ \\ \  | |\/  %sThe Lamb that was slain%s
         '" '"  '""         %sfor our sins.%s

`, bcolors.BLUE, bcolors.YELLOW + bcolors.YELLOW, 
   bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN,
                         bcolors.ENDC)

    case 2:
        fmt.Printf(`%s
         _      xxxx      _
        /_;-.__ / _\  _.-;_\
           '-._ ''_/''.-'
               '\   /'
                |  /
               /-.(
               \_._\
                \ \';
                 > |/        %sJesus Christ %s
                / //      %sL❤️.VE'S U. He is%s
                |//    %sThe Lamb that was slain%s
                \(\         %sfor our sins.%s

`, bcolors.BLUE, bcolors.YELLOW + bcolors.YELLOW, 
   bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN,
           bcolors.BLUE, bcolors.GREEN, bcolors.ENDC)

    case 3:
        fmt.Printf(`%s
                               .--,
                           ,.-( (o)\
                          /   .)/\ ')
                        .',./'/   )/
                    ()=///=))))==()
                      /
                      %swake up, Christian%s

              %sLord God Jesus Christ L❤️.VE'S you
                    follow the white Pigeon.
                     knock, knock, knock,
                          Man Of God.%s

`, bcolors.BLUE, bcolors.YELLOW + bcolors.YELLOW, 
   bcolors.GREEN, bcolors.GREEN, bcolors.ENDC)

    case 4:
        fmt.Printf(`%s
                                       ___
             _______                  /__/
            |.-----.|            ,---[___]*
            ||     ||           /    printer
            ||_____||    _____ /        ____
            |o_____*|   [o_+_+]--------[=i==]
             |     ________|           drive
             |  __|_
             '-/_==_\       %sJesus Christ is%s
              /_____\\  %sThe Lamb that was slain%s
                             %sfor our sins.%s
                               %sJohn 3:16%s

`, bcolors.BLUE, bcolors.YELLOW + bcolors.YELLOW, 
   bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN,
         bcolors.BLUE, bcolors.GREEN, bcolors.ENDC)

    case 5:
        fmt.Printf(`%s
           __________   __________ __________
          |          |\|          |          |\
          |  *    *  |||  *  *  * |        * ||
          |  *    *  |||          |     *    ||
          |  *    *  |||  *  *  * |  *       ||
          |__________|||__________|__________||
          |          || '---------------------'
          |  *    *  ||      %sJesus Christ is%s
          |          || %sThe Lamb that was slain%s
          |  *    *  ||      %sfor our sins.%s
          |__________||       %sJohn 3:16%s
           '----------'%s

`, bcolors.BLUE, bcolors.YELLOW + bcolors.YELLOW, 
   bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN,
        bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.ENDC)

    case 6:
        fmt.Printf(`%s
                         ____
                  __,-~~/~   '---.
                _/_,---(      ,   )_
              _/        <    /   )  \_
 - -----==;;;'==--------------------==';;;==----- -
              \/  ~"~"~"~"~"~\~"~)~"/
              (_ (   \  (     >    \)
                \_( _ <         >_>'
                   ~ '-i' ::>|--"    %sJesus Christ is%s
                       I;|.|.|   %sThe Lamb that was slain%s
                      <|i::|i|'.      %sfor our sins.%s
                     (' ^'"'-' ")       %sJohn 3:16%s 

`, bcolors.BLUE, bcolors.YELLOW + bcolors.YELLOW, 
   bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN,
         bcolors.BLUE, bcolors.GREEN, bcolors.ENDC)

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
               |   |    ;|   |
               |   \    //   |      %sJesus Christ is%s
               |    '._//'   |  %sThe Lamb that was slain%s
              .'             '.     %sfor our sins.%s
           _,'                 ',_     %sJohn 3:16%s

`, bcolors.BLUE, bcolors.YELLOW + bcolors.YELLOW, 
   bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN,
        bcolors.BLUE, bcolors.YELLOW, bcolors.ENDC)

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
               /_, ,  _..-'                    '-._\
               \_, \/ ._(      %sJesus Christ is%s
                \_, \/ ._\ %sThe Lamb that was slain%s
                 '._,\/ ._\     %sfor our sins.%s
                   '._// ./'-._
                     '-._-_-_.-'%s

`, bcolors.BLUE, bcolors.YELLOW + bcolors.YELLOW, 
  bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE,
                         bcolors.ENDC)
    }
}

func Banner() {
    Graphics()
}


