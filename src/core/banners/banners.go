package banners

import(
    "fmt"
    "time"
    "bcolors"
    "math/rand"
)
var (
    version = "3.0.2"
)

type Banners struct{}

func Version() {
    fmt.Printf(`
%sFramework %s: %s
%sConsole   %s: %s
`, bcolors.GREEN, bcolors.ENDC, version, bcolors.GREEN, bcolors.ENDC, version)
}

func Graphics() {
    rand.Seed(time.Now().UnixNano())
    menu := rand.Intn(9) + 1

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
       ) // /   | |\ \        %s%s%s%s%s
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
            \ \';
             > |/          %s%s%s%s%s
            / //         %sJesus Christ%s
            |//       %sL❤️.VE'S U. He is%s
            \(\    %sThe Lamb that was slain%s
                         %sfor our sins.%s
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
                ('.         ,-,
                ' '.    ,;' /
                 '.  ,'/ .'
                  '. X/ .'
        .-;--''--.._' . (
      .'            /   '
     ,           ' '   Q '
     ,         ,   '._    \
  ,.|         '     '-.;_.'
  :  . '  ;    '  ' --,.._;
   ' '    ,   )   .'              %s%s%s%s%s
      '._ ,  '   /_            %sJesus Christ is%s
         ; ,''-,;' ''-     %sThe Lamb that was slain%s
          ''-..__''--'          %sfor our sins.%s
                                  %sJohn 3:16%s
`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 5:
        fmt.Printf(`%s
                                       ___
             _______                  /__/
            |.-----.|            ,---[___]*
            ||     ||           /    printer
            ||_____||    _____ /        ____
            |o_____*|   [o_+_+]--------[=i==]
             |     ________|           drive
             |  __|_
             '-/_==_\
              /_____\\         %s%s%s%s%s
                             %sJesus Christ is%s
                        %sThe Lamb that was slain%s
                             %sfor our sins.%s
                               %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 6:
        fmt.Printf(`%s
           __________   __________ __________
          |          |\|          |          |\
          |  *    *  |||  *  *  * |        * ||
          |  *    *  |||          |     *    ||
          |  *    *  |||  *  *  * |  *       ||
          |__________|||__________|__________||
          |          || '---------------------'
          |  *    *  ||
          |          ||
          |  *    *  ||        %s%s%s%s%s
          |__________||     %sJesus Christ is%s
           '----------' %sThe Lamb that was slain%s
                              %sfor our sins.%s
                               %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 7:
        fmt.Printf(`%s
                        ____
                 __,-~~/~   '---.
               _/_,---(      ,   )_
             _/        <    /   )  \_
- -----==;;;'==--------------------==';;;==----- -
             \/  ~"~"~"~"~"~\~"~)~"/
             (_ (   \  (     >    \) 
               \_( _ <         >_>'
                  ~ '-i' ::>|--"
                      I;|.|.|            %s%s%s%s%s
                     <|i::|i|'.        %sJesus Christ is%s
                    (' ^'"'-' ")   %sThe Lamb that was slain%s
                                       %sfor our sins.%s
                                         %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 8:
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
        |   |    ;|   |           %s%s%s%s%s
        |   \    //   |        %sJesus Christ is%s
        |    '._//'   |   %sThe Lamb that was slain%s
       .'             '.        %sfor our sins.%s
    _,'                 ',_       %sJohn 3:16%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.ENDC)

    case 9:
        fmt.Printf(`%s
                   _,.---.---.---.--.._
               _.-' '--.'---.'---'-. _,'--.._
              /'--._ .'.     '.     ','-.'-._\
             ||   \  '.'---.__'__..-'. ,''-._/
        _  ,'\ '-._\   \    '.    '_.-'-._,''-.
     ,'   '-_ \/ '-.'--.\    _\_.-'\__.-'-.'-._'.
    (_.o> ,--. '._/'--.-',--'  \_.-'       \'-._ \
     '---'    '._ '---._/__,----'           '-. '-\
               /_, ,  _..-'       %s%s%s%s%s    '-._\
               \_, \/ ._(      %sJesus Christ is%s
                \_, \/ ._\ %sThe Lamb that was slain%s
                 '._,\/ ._\     %sfor our sins.%s
                   '._// ./'-._   %sJohn 3:16%s
                     '-._-_-_.-'%s

`, bcolors.BLUE, bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC, bcolors.BLUE, bcolors.YELLOW, bcolors.BLUE,
bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.PURPLE, bcolors.BLUE, bcolors.ENDC)
    }
}

func AnonSurfBanner() {
    fmt.Printf(`%s
          \                                       _|
         _ \   __ \   _ \  __ \   __| |   |  __| |
        ___ \  |   | (   | |   |\__ \ |   | |    __|
      _/    _\_|  _|\___/ _|  _|____/\__,_|_|   _|  %s%s%s%s

`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)
}

func GraphicsMini() {
    rand.Seed(time.Now().UnixNano())
    menu := rand.Intn(5) + 1

    switch menu {
    case 1:
        fmt.Printf(`%s
             _.-;;-._
      '-..-'|   ||   |
      '-..-'|_.-;;-._|
      '-..-'|   ||   |
      '-..-'|_.-''-._|

                                                  %s%s%s%s
`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)

    case 2:
        fmt.Printf(`%s
         .'"".'/|\'.""'.
        :  .' / | \ '.  :
        '.'  /  |  \  '.'
         '. /   |   \ .'
           '-.__|__.-'
                                                  %s%s%s%s
`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)

    case 3:
        fmt.Printf(`%s
         _
       (( )_,     ,
.--.    \ '/     /.\
    )   / \=    /O o\     _
   (   / _/    /' o O| ,_( ))___     ('
    '-|   )_  /o_O_'_(  \'    _ '\    ) 
      '""""'            ='---<___/---'
                            "'
                                                  %s%s%s%s
`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)

    case 4:
        fmt.Printf(`%s
 ______________________________________
/ If it were not for you Lord God Jesus\
\ where would I have been?             /
 '------------------------------------'
  \
   \
     ,__,
     (oo)____
     (__)    )\
        ||--|| *
                                                  %s%s%s%s
`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)

    case 5:
        fmt.Printf(`%s
 ______________________________________
/ If it were not for you Lord God Jesus\
\ where would I have been?             /
 --------------------------------------
 \
  \
     __
    /  \
    |  |
    @  @
    |  |
    || |/
    || ||
    |\_/|
    \___/
                                                  %s%s%s%s
`, bcolors.Colors(), bcolors.RED, bcolors.ITALIC, version, bcolors.ENDC)
    }
}

func Banner() {
    Graphics()
}
