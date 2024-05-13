import os
import random
from rgbprint import gradient_print

try:
    from colorama import Fore, Style
except ModuleNotFoundError:
    os.system("pip install colorama")

from urllib.request import urlopen
from Core.helper.color import green, white, blue, red, start, alert , yellow , purple

green = yellow
red = purple
white = purple

Version = "2.2"
yellow = ("\033[1;33;40m")

banner = r"""
                                                     ___
                                                  ,o88888
                                               ,o8888888'
                         ,:o:o:oooo.        ,8O88Pd8888"
                     ,.::.::o:ooooOoOoO. ,oO8O8Pd888'"
                   ,.:.::o:ooOoOoOO8O8OOo.8OOPd8O8O"
                  , ..:.::o:ooOoOOOO8OOOOo.FdO8O8"
                 , ..:.::o:ooOoOO8O888O8O,COCOO"
                , . ..:.::o:ooOoOOOO8OOOOCOCO"
                 . ..:.::o:ooOoOoOO8O8OCCCC"o
                    . ..:.::o:ooooOoCoCCC"o:o
                    . ..:.::o:o:,cooooCo"oo:o:
                 `   . . ..:.:cocoooo"'o:o:::'
                 .`   . ..::ccccoc"'o:o:o:::'
                :.:.    ,c:cccc"':.:.:.:.:.'
              ..:.:"'`::::c:"'..:.:.:.:.:.'
            ...:.'.:.::::"'    . . . . .'
           .. . ....:."' `   .  . . ''
         . . . ...."'
         .. . ."'
        .
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|"""


def connected(host='http://duckduckgo.com'):
    try:
        urlopen(host)
        return True
    except:
        return False


all_col = [Style.BRIGHT + Fore.RED, Style.BRIGHT + Fore.CYAN, Style.BRIGHT + Fore.LIGHTCYAN_EX,
           Style.BRIGHT + Fore.LIGHTBLUE_EX, Style.BRIGHT + Fore.LIGHTCYAN_EX, Style.BRIGHT + Fore.LIGHTMAGENTA_EX,
           Style.BRIGHT + Fore.LIGHTYELLOW_EX]

ran = random.choice(all_col)


def menu():

    gradient_print(banner, start_color='yellow' , end_color='magenta')
    print(white + "\n        ~>[ Select a number from the table below ]<~\n"   )
    print(blue + "[ 1. Instagram................12...................Paypal]")
    print(blue + "[ 2. Facebook.................13..................Discord]")
    print(blue + "[ 3. Gmail....................14..................Spotify]")
    print(blue + "[ 4. Gmail(Si)................15...............Blockchain]")
    print(blue + "[ 5. Twitter..................16................RiotGames]")
    print(blue + "[ 6. Snapchat.................17.................Rockstar]")
    print(blue + "[ 7. SnapchatS................18....................AskFM]")
    print(blue + "[ 8. Steam....................19.............. 000Webhost]")
    print(blue + "[ 9. Dropbox..................21..................Gamehag]")
    print(blue + "[ 10. Linkedin................22....................Mega ]")
    print(blue + "[ 11. Playstation........................................]")
    print(blue + "[ 0. .EXIT and go back to menu...........................]")

def Welcome():
    os.system("clear")

menu()

