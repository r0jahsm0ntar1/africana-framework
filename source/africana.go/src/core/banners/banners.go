package banners

import (
    "time"
    "math/rand"
    "bcolors"
)

func BannerZero() {
    bcolors.Colors(`
                     _,._
                 __.'   _)
                <_,)'.-"a\
                  /' (    \
      _.-----..,-'   ('"--^
     //              |
    (|   ';      ,   |
      \   ;.----/  ,/
       ) // /   | |\ \
       \ \\'\   | |/ /      Jesus Christ
        \ \\ \  | |\/  The Lamb that was slain
         '" '"  '""         for our sins.`)
}

func BannerOne() {
    bcolors.Colors(`
                 _      xxxx      _
                /_;-.__ / _\  _.-;_\
                   '-._ ' _/''.-'
                       '\   /'
                        |  /
                       /-.(
                       \_._\
                        \ \';
                         > |/
                        / //      Jesus Christ
                        |//   The Lamb that was slain
                        \(\       for our sins.`)
}

func BannerTwo() {
    bcolors.Colors(`
                               .--,
                           ,.-( (o)\
                          /   .)/\ ')
                        .',./'/   )/
                    ()=///=))))==()
                      /
                      wake up, Christian
              Lord God Jesus Christ L❤️.VE'S you
                  follow the white Pigeon.
                    knock, knock, knock,
                          Man Of God.`)
}

func BannerThree() {
    bcolors.Colors(`
                                       ___
             _______                  /__/
            |.-----.|            ,---[___]*
            ||     ||           /    printer
            ||_____||    _____ /        ____
            |o_____*|   [o_+_+]--------[=i==]
             |     ________|           drive
             |  __|_
             '-/_==_\       Jesus Christ is
              /_____\\  The Lamb that was slain
                             for our sins.
                               John 3:16`)
}

func BannerFour() {
    bcolors.Colors(`
           __________   __________ __________
          |          |\|          |          |\
          |  *    *  |||  *  *  * |        * ||
          |  *    *  |||          |     *    ||
          |  *    *  |||  *  *  * |  *       ||
          |__________|||__________|__________||
          |          || '---------------------'
          |  *    *  ||       Jesus Christ
          |          ||  The Lamb that was slain
          |  *    *  ||       for our sins.
          |__________||         John 3:16
           '----------'`)
}

func BannerFive() {
    bcolors.Colors(`
                                ____
                        __,-~~/~    '---.
                      _/_,---(      ,    )
                    _/        <    /   )  \___
   - ------===;;;'===----------------------===';;;===----- -
                     \/  ~"~"~"~"~"~\~"~)~"/
                     (_ (   \  (     >    \)
                      \_( _ <         >_>'
                         ~ '-i' ::>|--"
                             I;|.|.|
                            <|i::|i|'.
                           (' ^'"'-' ")`)
}

func BannerSix() {
    bcolors.Colors(`
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
                       |    '._//'   |  John 3:16
                      .'             '.   Agape
                   _,'                 ',_`)
}

func BannerSeven() {
    bcolors.Colors(`
                   _,.---.---.---.--.._ 
               _.-' '--.'---.'---'-. _,'--.._
              /'--._ .'.     '.     ','-.'-._\
             ||   \  '.'---.__'__..-'. ,''-._/
        _  ,'\ '-._\   \    '.    '_.-'-._,''-.
     ,'   '-_ \/ '-.'--.\    _\_.-'\__.-'-.'-._'.
    (_.o> ,--. '._/'--.-',--'  \_.-'       \'-._ \
     '---'    '._ '---._/__,----'           '-. '-\
               /_, ,  _..-'                    '-._\
               \_, \/ ._(      Jesus Christ is
                \_, \/ ._\ The Lamb that was slain
                 '._,\/ ._\     for our sins.
                   '._// ./'-._
                     '-._-_-_.-'`)
}


func BannerEight() {
    bcolors.Colors(`
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
`)
}

func Banner() {
    bannersFunc := []func(){
    BannerZero, BannerOne,    BannerTwo,
    BannerThree, BannerFour, BannerFive,
    BannerSix, BannerSeven,
    }

    rand.Seed(time.Now().UnixNano())
    selectedIndex := rand.Intn(len(bannersFunc))
    selectedFuncs := bannersFunc[selectedIndex]
    selectedFuncs()
}
