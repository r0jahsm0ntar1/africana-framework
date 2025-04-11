package banners

import(
    "fmt"
    "time"
    "bcolors"
    "math/rand"
)
var (
    version = "3.0.3-dev"
)

type Banners struct{}

func Version() {
    fmt.Printf(`
     %sFramework%s: %s
       %sConsole%s: %s

`, bcolors.Green, bcolors.Endc, version, bcolors.Green, bcolors.Endc, version)
}

func GraphicsLarge() {
    rand.Seed(time.Now().UnixNano())
    gbanner := rand.Intn(7) + 1

    switch gbanner {
    case 1:
        fmt.Printf(`%s
     _      xxxx      _
    /_;-.__ / _\  _.-;_\
       '-._ ''_/''.-'
           '\   /'
            |  /
           /-.(
           \_._\
            \ \';
             > |/
            / //
            |//
            \(\%s
`, bcolors.Colors(), bcolors.Endc)

    case 2:
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
       \ \\'\   | |/ /
        \ \\ \  | |\/
         '" '"  '""%s
`, bcolors.Colors(), bcolors.Endc)

    case 3:
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
        |   \    //   |
        |    '._//'   |
       .'             '.
    _,'                 ',_%s
`, bcolors.Colors(), bcolors.Endc)

    case 4:
        fmt.Printf(`%s
                ('.         ,-,
                ' '.    ,;'  /
                 '.  ,'/  . '
                  '. X/ .'
        .-;--''--.._' . (
      .'            /   '
     ,           ' '   Q '
     ,         ,   '._    \
  ,.|         '     '-.;_.'
  :  . '  ;    '  ' --,.._;
   ' '    ,   )   .'
      '._ ,  '   /_
         ; ,''-,;' ''-
          ''-..__''--'%s
`, bcolors.Colors(), bcolors.Endc)

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
      /_____\\%s
`, bcolors.Colors(), bcolors.Endc)

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
    |  *    *  ||
    |__________||
     '----------'%s
`, bcolors.Colors(), bcolors.Endc)

    case 7:
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
               \_, \/ ._(
                \_, \/ ._\
                 '._,\/ ._\
                   '._// ./'-._ 
                     '-._-_-_.-'%s
`, bcolors.Colors(), bcolors.Endc)
    }
}

func GraphicsTinny() {
    rand.Seed(time.Now().UnixNano())
    tbanner := rand.Intn(5) + 1

    switch tbanner {
    case 1:
        fmt.Printf(`%s
    ,__,
    (oo)____
    (__)    )\
       ||--|| *%s%s%s
`, bcolors.Colors(), bcolors.Red, bcolors.Italic, bcolors.Endc)

    case 2:
        fmt.Printf(`%s
     __
    /  \
    |  |
    @  @
    |  |
    || |/
    || ||
    |\_/|
    \___/%s%s%s
`, bcolors.Colors(), bcolors.Red, bcolors.Italic, bcolors.Endc)

    case 3:
        fmt.Printf(`%s
               .--,
           ,.-( (o)\
          /   .)/\ ')
        .',./'/   )/
    ()=///=))))==()
      /%s%s%s
`, bcolors.Colors(), bcolors.Red, bcolors.Italic, bcolors.Endc)

    case 4:
        fmt.Printf(`%s
     .'"".'/|\'.""'.
    :  .' / | \ '.  :
    '.'  /  |  \  '.'
     '. /   |   \ .'
       '-.__|__.-'%s%s%s
`, bcolors.Colors(), bcolors.Red, bcolors.Italic, bcolors.Endc)

    case 5:
        fmt.Printf(`%s
         _
       (( )_,     ,
.--.    \ '/     /.\
    )   / \=    /O o\     _
   (   / _/    /' o O| ,_( ))___     ('
    '-|   )_  /o_O_'_(  \'    _ '\    ) 
      '""""'            ='---<___/---'"'%s%s%s
`, bcolors.Colors(), bcolors.Red, bcolors.Italic, bcolors.Endc)
    }
}

func GraphicsTorNet() {
    fmt.Printf(`%s
 _                       _ 
| |_ ___ ___ ___ ___ ___| |_ ___ 
|  _| . |  _|_ -| . |  _| '_|_ -|
|_| |___|_| |___|___|___|_,_|___|%s%s%s
`, bcolors.Colors(), bcolors.Red, bcolors.Italic, bcolors.Endc)
}

func RandomBanners() {
    rand.Seed(time.Now().UnixNano())
    rbanner := rand.Intn(2) + 1

    switch rbanner {
    case 1:
        GraphicsLarge(); fmt.Println()
    case 2:
        GraphicsTinny(); fmt.Println()
    }
}
