import time
import random
from src.core.bcolors import *

class banners(object):
    def __init__(self):
        pass
    def graphics(self):
        menu = random.randrange(1, 7)
        if menu == 1:
            print(color() + r"""
                     _,._
                 __.'   _)
                <_,)'.-"a\
                  /' (    \
      _.-----..,-'   (`"--^
     //              |
    (|   `;      ,   |
      \   ;.----/   ,/
       ) // /   | |\ \
       \ \\`\   | |/ /      Jesus Christ
        \ \\ \  | |\/  The Lamb that was slain
         `" `"  `"`         for our sins.
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
""" + bcolors.ENDC)

        if menu == 2:
            print(color() + r"""
                 _      xxxx      _
                /_;-.__ / _\  _.-;_\
                   `-._`'`_/'`.-'
                       `\   /`
                        |  /
                       /-.(
                       \_._\
                        \ \`;
                         > |/
                        / //      Jesus Christ
                        |//   The Lamb that was slain
                        \(\       for our sins.
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
""" + bcolors.ENDC)

        if menu == 3:
            print(color() + r"""
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
                          Man Of God.
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
""" + bcolors.ENDC)

        if menu == 4:
            print(color() + r"""
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
                               John 3:16
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
""" + bcolors.ENDC)

        if menu == 5:
            print(color() + r"""
           __________   __________ __________
          |          |\|          |          |\
          |  *    *  |||  *  *  * |        * ||
          |  *    *  |||          |     *    ||
          |  *    *  |||  *  *  * |  *       ||
          |__________|||__________|__________||
          |          || `---------------------`
          |  *    *  ||       Jesus Christ
          |          ||  The Lamb that was slain
          |  *    *  ||       for our sins.
          |__________||         John 3:16
           `----------`
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
""" + bcolors.ENDC)

        if menu == 6:
            print(color() + r"""
                                ____
                        __,-~~/~    `---.
                      _/_,---(      ,    )
                  __ /        <    /   )  \___
   - ------===;;;'===----------------------===';;;===----- -
                     \/  ~"~"~"~"~"~\~"~)~"/
                     (_ (   \  (     >    \)
                      \_( _ <         >_>'
                         ~ `-i' ::>|--"
                             I;|.|.|
                            <|i::|i|`.
                           (` ^'"`-' ")
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
""" + bcolors.ENDC)

        if menu == 7:
            print(color() + r"""
                       .--.       .--.
                   _  `    \     /    `  _
                    `\.===. \.^./ .===./`
                           \/`"`\/
                        ,  |     |  ,
                       / `\|`-.-'|/` \
                      /    |  \  |    \
                   .-' ,-'`|   ; |`'-, '-.
                       |   |    \|   | 
                       |   |    ;|   |
                       |   \    //   |
                       |    `._//'   |  John 3:16
                      .'             `.   Agape
                   _,'                 `,_
                   `
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
""" + bcolors.ENDC)

beauty = banners()
if ' __name__' == '__main__':
        sys.exit(beauty())
